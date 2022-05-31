package main

import (
	"context"
	"encoding/json"
	"fmt"
	pro "go-test/module/test-grpc/proto"
	"go-test/third/protobuf/google/gogo/protobuf/types"
	"google.golang.org/grpc"
	"net"
)

const PORT = ":50051"

type server struct {
}

func (this *server) GetGreeterInfo(ctx context.Context, req *pro.GetGreeterInfoReq) (resp *pro.GetGreeterInfoResp, err error) {
	//mp := map[string]interface{}{
	//	"a": 1,
	//	"b": "b",
	//}
	mp2 := map[string]interface{}{
		"a": 2,
		"c": "c",
	}
	//b, _ := json.Marshal(mp)
	b2, _ := json.Marshal(mp2)
	et := pro.GetGreeterInfoEntity{
		Name: "张三",
	}
	var b3 []byte
	b3, _ = et.Marshal()
	fmt.Println("======", string(b3))
	resp = &pro.GetGreeterInfoResp{Data: []*types.Any{{
		TypeUrl: "type.googleapis.com/test.student.v1.GetGreeterInfoEntity",
		Value:   []byte(b3),
	}, {
		TypeUrl: "type.googleapis.com/google.protobuf.StringValue",
		Value:   []byte(b2),
	}}}

	return
}

func main() {
	//监听端口
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		return
	}
	//创建一个grpc 服务器
	s := grpc.NewServer()
	//注册事件
	pro.RegisterGreeterServer(s, &server{})
	//处理链接
	s.Serve(lis)
}
