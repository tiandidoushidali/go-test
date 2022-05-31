package main

import (
	"fmt"
	pro "go-test/module/test-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

const STREAM_SERVER_PORT = ":50052"

type streamServer struct {
}

//服务端 单向流
func (s *streamServer) GetStream(req *pro.StreamReqData, res pro.GreeterStream_GetStreamServer) error {
	i := 0
	for {
		i++
		res.Send(&pro.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(1 * time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

//客户端 单向流
func (this *streamServer) PutStream(cliStr pro.GreeterStream_PutStreamServer) error {

	for {
		if tem, err := cliStr.Recv(); err == nil {
			log.Println(tem)
		} else {
			log.Println("break, err :", err)
			break
		}
	}

	return nil
}

//客户端服务端 双向流
func (this *streamServer) AllStream(allStr pro.GreeterStream_AllStreamServer) error {

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
		wg.Done()
	}()

	go func() {
		for {
			allStr.Send(&pro.StreamResData{Data: "ssss"})
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func main() {
	//监听端口
	lis, err := net.Listen("tcp", STREAM_SERVER_PORT)
	if err != nil {
		return
	}
	//创建一个grpc 服务器
	s := grpc.NewServer()
	//注册事件
	pro.RegisterGreeterStreamServer(s, &streamServer{})
	//处理链接
	s.Serve(lis)
}
