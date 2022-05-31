package main

import (
	"context"
	"fmt"
	pro "go-test/module/test-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
	"runtime"
	"time"

	_ "go-test/module/test-grpc/balancer"
	_ "go-test/module/test-grpc/resolver"
)

const (
	STUDENT_ADDRESS = "127.0.0.1:10002"
	//STUDENT_ADDRESS = "dis://default/student"
)

type DisBuilder struct {
}

type customResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
}

func (r customResolver) start() {
}
func (*customResolver) ResolveNow(o resolver.ResolveNowOption) {}
func (*customResolver) Close()                                 {}

func (b *DisBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	r := &customResolver{
		target: target,
		cc:     cc,
	}

	return r, nil
}

func (b *DisBuilder) Scheme() string {
	return "dis"
}
func init() {
	resolver.Register(&DisBuilder{})
}

type GrpcAuth struct {
	AppId     string
	AppSecret string
}

func (auth *GrpcAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"appKey": auth.AppId, "appSecret": auth.AppSecret}, nil
}

func (auth *GrpcAuth) RequireTransportSecurity() bool {
	return true
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
	conn, err := grpc.DialContext(ctx, "grpcc://default/service", grpc.WithInsecure(), grpc.WithAuthority("test_id"), grpc.WithBalancerName("gbalancer"))
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
		fmt.Println("----errr1----", err)
		errStatus := status.Convert(err)
		fmt.Println("----details----", errStatus.Proto().Details)
		fmt.Printf("获取数据失败：%v, %s", errStatus.Code(), errStatus.Message())
		return
	}
	fmt.Println(fmt.Sprintf("接受response resp：%+v", resp))

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
