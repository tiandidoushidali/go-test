package main

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	pro "go-test/module/test-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/metadata"
	"reflect"
	"runtime"
	"time"

	"github.com/gogo/protobuf/jsonpb"
)

type rpcJson []byte

func (j rpcJson) String() string {
	return string(j)
}

type Reply struct {
	res rpcJson
}

func (j *Reply) String() string {
	return string(j.res)
}

func (j *Reply) Byte() []byte {
	return j.res
}

// JSON is impl of encoding.Codec
type JSON struct {
	jsonpb.Marshaler
	jsonpb.Unmarshaler
}

// Name is name of JSON
func (j JSON) Name() string {
	return "json"
}

// Marshal is json marshal
func (j JSON) Marshal(v interface{}) (out []byte, err error) {
	//v.(rpcJson)
	log.Info("proxy json marshal out %v", out)
	return v.(rpcJson), nil
}

// Unmarshal is json unmarshal
func (j JSON) Unmarshal(data []byte, v interface{}) (err error) {
	log.Info("proxy json unmarshal v typeof %v", reflect.TypeOf(v))
	v.(*Reply).res = data
	log.Info("proxy json unmarshal v2 typeof %v", reflect.TypeOf(v))
	return nil
}

func main() {
	//通过grpc 库 建立一个连接
	//conn ,err := grpc.Dial(STUDENT_ADDRESS,grpc.WithInsecure(),grpc.WithTimeout(60 * time.Second))
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	//auth := &GrpcAuth{
	//	AppId:     "test_id",
	//	AppSecret: "test_secret",
	//}
	//conn, err := grpc.DialContext(ctx, STUDENT_ADDRESS, grpc.WithInsecure(), grpc.WithPerRPCCredentials(auth))
	encoding.RegisterCodec(JSON{})
	conn, err := grpc.DialContext(ctx, "127.0.0.1:10002", grpc.WithInsecure(), grpc.WithAuthority("test_id1"), grpc.WithDefaultCallOptions(grpc.CallContentSubtype(JSON{}.Name())))
	if err != nil {
		fmt.Println("-------err-----", err)
		return
	}
	defer func() {
		cancel()
		conn.Close()
	}()
	//通过刚刚的连接 生成一个client对象。
	c := pro.NewStudentClient(conn)

	md := metadata.Pairs("color", "api-pop-gateway-test")
	ctx = metadata.NewOutgoingContext(ctx, md)
	reqstreamData := &pro.StudentInfoReq{
		Id: 2,
	}

	resp, err := c.StudentInfo(ctx, reqstreamData)
	if err != nil {
		fmt.Printf("获取数据失败：%s", err)
		return
	}
	fmt.Println(fmt.Sprintf("接受response：%+v %p", resp, &resp))

	return
	go func(reqstreamData *pro.StudentInfoReq) {
		resp, err := c.StudentInfo(ctx, reqstreamData)
		if err != nil {
			fmt.Printf("获取数据失败：%s", err)
			return
		}
		fmt.Println(fmt.Sprintf("接受response：%+v %p", resp, &resp))
	}(reqstreamData)

	go func(reqstreamData *pro.StudentInfoReq) {
		runtime.Gosched()
		r := *reqstreamData
		r.Id = 3
		resp, err := c.StudentInfo(ctx, &r)
		if err != nil {
			fmt.Printf("获取数据失败：%s", err)
			return
		}
		fmt.Println(fmt.Sprintf("接受response：%+v %p", resp, &resp))
	}(reqstreamData)

	go func(reqstreamData *pro.StudentInfoReq) {
		runtime.Gosched()
		time.Sleep(6 * time.Second)
		r := *reqstreamData
		r.Id = 4
		resp, err := c.StudentInfo(ctx, &r)
		if err != nil {
			fmt.Printf("获取数据失败：%s", err)
			return
		}
		fmt.Println(fmt.Sprintf("接受response：%+v %p", resp, &resp))
	}(reqstreamData)

	time.Sleep(10 * time.Second)
}
