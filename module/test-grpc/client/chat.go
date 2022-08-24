package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"go-test/grape/protocol/message"
	pro_med "go-test/module/test-grpc/proto-med"
	"go-test/third/protobuf/google/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"time"
)

func main() {

	//通过grpc 库 建立一个连接
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "172.31.8.153:8000", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pro_med.NewChatClient(conn)
	resp, err := c.UserChatList(ctx, &pro_med.UserChatListReq{
		UserId:               79183974,
	})
	if err != nil {
		fmt.Println("----", err)
	}
	fmt.Println("----resp.List----", len(resp.UserChatList))
	for _, v := range resp.UserChatList {
		data := v
		decoded, _ := base64.StdEncoding.DecodeString(data.Message)
		//fmt.Println("---decoded-----", decoded)
		msg := &message.Message{}
		proto.Unmarshal(decoded, msg)


		fmt.Println("-----", data, msg)
	}
}
