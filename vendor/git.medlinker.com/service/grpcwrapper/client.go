package grpcwrapper

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"git.medlinker.com/service/grpcwrapper/middleware"
)

// gRPC客户端连接配置
type ClientOpt struct {
	Interceptors []grpc.UnaryClientInterceptor
	CallBlock    bool
}

var (
	// 默认连接超时时间
	defaultDialTimeout = time.Second * 15
	// 默认客户端连接配置
	defaultClientOpt = &ClientOpt{
		Interceptors: make([]grpc.UnaryClientInterceptor, 0),
	}
)

// 默认客户端选项
func DefaultClient() *ClientOpt {
	return defaultClientOpt
}

// 创建gRPC客户端
func (co *ClientOpt) DialContext(ctx context.Context, address string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	co.Interceptors = append(co.Interceptors, middleware.ClientErrorHook)
	co.Interceptors = append(co.Interceptors, middleware.UnaryClientRecoverInterceptor())

	dops := make([]grpc.DialOption, 0)
	dops = append(dops, grpc.WithUnaryInterceptor(ChainUnaryClient(co.Interceptors)))
	dops = append(dops, grpc.WithInsecure())
	if co.CallBlock {
		dops = append(dops, grpc.WithBlock())
	}
	dops = append(dops, opts...)

	newCtx, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	conn, err = grpc.DialContext(newCtx, address, dops...)
	defer cancel()

	return
}

// 创建gRPC客户端对象(使用默认配置)
func NewClient(address string) (conn *grpc.ClientConn, err error) {
	return DefaultClient().DialContext(context.Background(), address, grpc.WithBlock())
}
