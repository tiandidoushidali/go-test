package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/lithammer/shortuuid"
	uuid "github.com/satori/go.uuid"
	"go-test/module/test-etcd/discovery"
	"net"
	"runtime"
	"time"
)

var (
	port string
	wsPort string
)
func init() {
	flag.StringVar(&port, "port", "6001", "端口:80")
	flag.StringVar(&wsPort, "wsport", "6002", "ws端口：8080")
}
func main() {
	etcdAddr := "http://172.16.70.151:2379"
	dis := discovery.New([]string{etcdAddr})
	if dis == nil {
		panic("dis err")
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	wsln, err := net.Listen("tcp", fmt.Sprintf(":%s", wsPort))

	ln2, err := net.Listen("tcp", fmt.Sprintf(":%s", "6003"))
	wsln2, err := net.Listen("tcp", fmt.Sprintf(":%s", "6004"))

	if err != nil {
		panic(err)
	}
	uuidV4 := uuid.NewV4()
	cometname := shortuuid.DefaultEncoder.Encode(uuidV4)
	fmt.Println("----", ln, port, wsPort)
	go dis.Hold(cometname, ln.Addr().(*net.TCPAddr).Port, wsln.Addr().(*net.TCPAddr).Port)
	uuidV42 := uuid.NewV4()
	cometname2 := shortuuid.DefaultEncoder.Encode(uuidV42)
	go dis.Hold(cometname2, ln2.Addr().(*net.TCPAddr).Port, wsln2.Addr().(*net.TCPAddr).Port)
	go dis.Watch()

	go func() {
		timer := time.NewTicker(5 * time.Second)
		ctx := context.Background()
		for {
			select {
			case <-timer.C:
				dis.Get(ctx, "/grape/comet/")
			}
		}
	}()
	serveComet(ln)
}

func serveComet(ln net.Listener) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("comet crash:", err)
		}
	}()
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Error("accept connection err:", err)
			runtime.Gosched()
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	fmt.Println("new connection ")
}