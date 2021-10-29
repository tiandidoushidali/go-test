package main

import (
	"context"
	"fmt"
	pro "go-test/module/test-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

const (
	STUDENT_ADDRESS = "127.0.0.1:10002"
)
func main() {
	//通过grpc 库 建立一个连接
	conn ,err := grpc.Dial(STUDENT_ADDRESS,grpc.WithInsecure(),grpc.WithTimeout(60 * time.Second))
	if err != nil{
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pro.NewStudentClient(conn)

	md := metadata.Pairs("color", "api-pop-gateway-test")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	reqstreamData := &pro.StudentInfoReq{
		Id:                   2,
	}
	resp, err := c.StudentInfo(ctx, reqstreamData)
	if err != nil {
		fmt.Printf("获取数据失败：%s", err)
		return
	}

	fmt.Println("接受response：", *resp)
}
