package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/serialx/hashring"
	"go-test/module/test-etcd/discovery"
)

func main() {
	cli, err := client.New(client.Config{
		//Endpoints:               []string{"http://172.31.7.205:2379"}, // 腾讯云
		Endpoints:               []string{"http://172.25.9.243:2379"},
	})
	if err != nil {
		panic(err)
	}
	kapi := client.NewKeysAPI(cli)

	ctx := context.TODO()
	resp, _ := kapi.Get(ctx, "/grape/comet", nil)
	fmt.Println(fmt.Sprintf("----resp----%+v", resp.Node))
	nodes := make([]string, 0)
	tmp := make(map[string]discovery.Addr)
	if resp != nil {
		for _, ev := range resp.Node.Nodes {
			addr := new(discovery.Addr)
			if err := json.Unmarshal([]byte(ev.Value), addr); err != nil {
				continue
			}
			nodes = append(nodes, ev.Key)
			tmp[ev.Key] = *addr

			fmt.Println("----addr----", addr)
		}
	}

	hr := hashring.New(nodes)
	node, ok := hr.GetNode("78361646")
	if ok {
		fmt.Println("---node----", node)
	}
	fmt.Println("---tmp---", tmp)
}