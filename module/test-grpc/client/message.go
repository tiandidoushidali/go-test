package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-test/grape/protocol/message"
	pro_med "go-test/module/test-grpc/proto-med"
	"go-test/third/protobuf/google/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"strings"
	"time"
)

func main() {

	//通过grpc 库 建立一个连接
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "172.20.17.225:8000", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pro_med.NewMessageClient(conn)
	resp, err := c.MessageGetHistory(context.Background(), &pro_med.MessageGetHistoryReq{
		GroupId:              13755019,
		Limit:                50,
		Start:                0,
		Sort:                 1,
	})
	if err != nil {
		fmt.Println("----", err)
	}
	fmt.Println("----resp.List----", len(resp.List))
	for _, v := range resp.List {
		data := v
		decoded, _ := base64.StdEncoding.DecodeString(data)
		//fmt.Println("---decoded-----", decoded)
		msg := &message.Message{}
		proto.Unmarshal(decoded, msg)
		if strings.Contains(msg.GetJson().GetContent(), "2204249163150638") {

			b, _ := json.Marshal(msg)
			fmt.Println(string(b))
		}
	}
	fmt.Println("-----resp-----", resp.Index)
}
