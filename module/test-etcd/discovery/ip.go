package discovery


import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/pkg/log"
	"go-test/module/test-etcd/config"
	"net"
	"sync"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/serialx/hashring"
)

type inetType uint8

const (
	LOOPBK inetType = 0 + iota
	PUBLIC
	PRIVATE
)

var (
	// rfc1918
	rfc1918 = []*net.IPNet{
		{IP: net.IPv4(0x0A, 0x00, 0x00, 0x00), Mask: net.CIDRMask(0x08, 32)}, // 10.0.0.0/8
		{IP: net.IPv4(0xAC, 0x10, 0x00, 0x00), Mask: net.CIDRMask(0x0C, 32)}, // 172.16.0.0/12
		{IP: net.IPv4(0xC0, 0xA8, 0x00, 0x00), Mask: net.CIDRMask(0x10, 32)}, // 192.168.0.0/16
	}

	PATH = "/grape/comet"
)


type Discover struct {
	m    sync.RWMutex       //读写锁
	c    client.Client      //etcd 客户端
	h    *hashring.HashRing //一个hash环，用来分布式保存comet节点，node为comet的名字
	t    map[string]*Addr   //key是comet的名字，value为他的三个地址
	f    client.Config
	quit bool
}
type Addr struct {
	Loopback Server `json:"loopback,omitempty"`
	Public   Server `json:"public,omitempty"`
	Private  Server `json:"private,omitempty"`
}

type Server struct {
	Tcp string `json:"tcp"`
	Ws  string `json:"ws"`
}

func New(addrs []string) *Discover {
	conf := client.Config{
		Endpoints: addrs,
		Transport: client.DefaultTransport,
	}
	d := new(Discover)
	d.f = conf
	d.connect()
	return d
}

// 连接etcd 初始化
func (d *Discover) connect() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("connect fail:", err)
		}
	}()
	var err error
	d.m.Lock()
	defer d.m.Unlock()
	d.c, err = client.New(d.f)
	if err != nil {
		panic(err)
	}
	kapi := client.NewKeysAPI(d.c)
	ctx := context.TODO()
	resp, _ := kapi.Get(ctx, PATH, nil)
	nodes := make([]string, 0)
	tmp := make(map[string]*Addr)
	if resp != nil {
		for _, ev := range resp.Node.Nodes {
			addr := new(Addr)
			if err := json.Unmarshal([]byte(ev.Value), addr); err != nil {
				continue
			}
			nodes = append(nodes, ev.Key)
			tmp[ev.Key] = addr
		}
	}
	d.h = hashring.New(nodes)
	d.t = tmp
}

func DetectInetType(ip net.IP) inetType {
	if ip.IsLoopback() {
		return LOOPBK
	}
	for _, cidr := range rfc1918 {
		if cidr.Contains(ip) {
			return PRIVATE
		}
	}
	return PUBLIC
}

func detectCometAddr(port int, wsPort int, nodeName string) *Addr {
	selfAddr := new(Addr)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				addr := net.TCPAddr{IP: ipnet.IP, Port: port}
				wsAddr := net.TCPAddr{IP: ipnet.IP, Port: wsPort}

				switch DetectInetType(ipnet.IP) {
				case LOOPBK:
					selfAddr.Loopback.Tcp = addr.String()
					selfAddr.Loopback.Ws = wsAddr.String()
				case PRIVATE:
					selfAddr.Private.Tcp = addr.String()
					selfAddr.Private.Ws = wsAddr.String()
				}
			}
		}
		selfAddr.Public.Tcp = config.NodeHubMap[nodeName].Tcp
		selfAddr.Public.Ws = config.NodeHubMap[nodeName].Ws
	}
	return selfAddr
}

// 注册自身到etcd 并保持更新 仅comet需要调用
func (d *Discover) Hold(name string, port int, wsPort int) {
L:
	time.Sleep(3 * time.Second)

	// 注册
	kapi := client.NewKeysAPI(d.c)
	ctx := context.TODO()
	jsonByte, _ := json.Marshal(detectCometAddr(port, wsPort, name))
	fmt.Println("path and name ", PATH+"/"+name)
	_, err := kapi.Set(ctx, fmt.Sprintf(PATH+"/"+name), string(jsonByte), &client.SetOptions{TTL: 15 * time.Second})
	log.Info("register node:", string(jsonByte))
	fmt.Println("register node:", string(jsonByte))
	if err != nil {
		log.Error("register node fail:", err)
		goto L
	}
	// 更新
	for {
		select {
		case <-time.After(time.Second * 12):
			if d.quit {
				return
			}
			_, err := kapi.Set(ctx, PATH+"/"+name, "", &client.SetOptions{TTL: 15 * time.Second, Refresh: true})
			if err != nil {
				log.Error("reset expire fail:", err)
				goto L
			}
		}
	}
}


// 监听更改
func (d *Discover) Watch() {
	goto L2
L1:
	time.Sleep(3 * time.Second)
	log.Info("reconnecting etcd...")
	d.connect()
L2:
	kapi := client.NewKeysAPI(d.c)
	watcher := kapi.Watcher(PATH, &client.WatcherOptions{
		Recursive: true,
	})
	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			log.Error("watch etcd fail:", err)
			goto L1
		}
		log.Info("update node:", res.Action, res.Node.Key, res.Node.Value)
		d.m.Lock()
		if res.Action == "expire" || res.Action == "delete" {
			d.h = d.h.RemoveNode(res.Node.Key)
			delete(d.t, res.Node.Key)
		}
		if res.Action == "set" || res.Action == "update" {
			addr := new(Addr)
			if err := json.Unmarshal([]byte(res.Node.Value), addr); err != nil {
				log.Info("delete bad value node", res.Node.Key, res.Node.Value)
				d.h = d.h.RemoveNode(res.Node.Key)
				delete(d.t, res.Node.Key)
				d.m.Unlock()
				continue
			}
			if res.Action == "set" {
				d.h = d.h.AddNode(res.Node.Key)
			}
			d.t[res.Node.Key] = addr
		}
		d.m.Unlock()
		log.Info("all current nodes:", d.t)
	}
}

func (d *Discover) Get(ctx context.Context, key string) {
	kapi := client.NewKeysAPI(d.c)
	resp, err := kapi.Get(ctx, key, &client.GetOptions{
		Recursive: false,
		Sort:      false,
		Quorum:    false,
	})
	if err != nil {
		log.Errorc(ctx, "get from etcd error ", err)
		return
	}
	fmt.Println(fmt.Sprintf("data from etcd %+v", resp))
	nodes := make([]string, 0)
	tmp := make(map[string]*Addr)
	if resp != nil {
		for _, v := range resp.Node.Nodes {
			log.Infoc(ctx, "node data from etcd %+v %v, %s", v, v.Nodes, v.Value)
			addr := new(Addr)
			json.Unmarshal([]byte(v.Value), addr)
			nodes = append(nodes, v.Key)

			tmp[v.Key] = addr
		}
	}

	hr := hashring.New(nodes)
	node, ok := hr.GetNode("10000")
	fmt.Println("-----node-----", ok , node, *tmp[node])
}