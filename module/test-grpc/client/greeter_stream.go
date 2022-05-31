package main

import (
	"context"
	_ "github.com/gogo/protobuf/gogoproto"
	pro "go-test/module/test-grpc/proto"
	"go-test/third/github.com/gogo/protobuf/types"
	"log"
	"time"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/grpclb"
)

const (
	GREETER_STREAM_ADDRESS = "localhost:50052"
)

func main() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(GREETER_STREAM_ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pro.NewGreeterStreamClient(conn)
	//调用服务端推送流
	reqstreamData := &pro.StreamReqData{Data: "aaa"}
	res, _ := c.GetStream(context.Background(), reqstreamData)
	for {
		aa, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(aa)
	}
	//客户端 推送 流
	putRes, _ := c.PutStream(context.Background())
	i := 1
	for {
		i++
		putRes.Send(&pro.StreamReqData{Data: "ss"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	//服务端 客户端 双向流
	allStr, _ := c.AllStream(context.Background())
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
	}()

	go func() {
		for {
			allStr.Send(&pro.StreamReqData{Data: "ssss", IllName: &types.StringValue{Value: "aaa"}})
			time.Sleep(time.Second)
		}
	}()

	select {}

}
