package grpc_balancer

import (
	"context"
	"fmt"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/resolver"
	"sync"
)

func init() {
	fmt.Println("----进入自定义balancer init-------")
	balancer.Register(NewBalancerBuilder())
}

func NewBalancerBuilder() balancer.Builder {
	return base.NewBalancerBuilder(getBalancerName(), &GrpcPickerBuilder{})
}

func getBalancerName() string {
	return "gbalancer"
}

type GrpcPickerBuilder struct {
	service string
}

// build 返回balancer.Picker 
func (b *GrpcPickerBuilder) Build(readyScs map[resolver.Address]balancer.SubConn) balancer.Picker {
	fmt.Println("-----进入自定义balancer build-------")
	connMap := make(map[string]balancer.SubConn)
	connSlice := make([]balancer.SubConn, 0)
	for addr, sc := range readyScs {
		connMap[addr.Addr] = sc
		connSlice = append(connSlice, sc)
	}
	
	return &GrpcPicker{
		connMap: connMap,
		connSlice: connSlice,
	}
}

type GrpcPicker struct {
	connMap map[string]balancer.SubConn
	connSlice []balancer.SubConn
	
	Server string
	mu sync.Mutex
	next int
}

// 实现负载均衡算法
func (p *GrpcPicker) Pick(ctx context.Context, opts balancer.PickOptions) (balancer.SubConn, func(info balancer.DoneInfo), error) {
	if len(p.connSlice) <= 0 {
		return nil, nil, balancer.ErrNoSubConnAvailable
	}
	
	fmt.Println("---自定义pick----", len(p.connSlice))
	
	// 先选址，这样可以复用负载均衡权重
	return p.connSlice[0], nil, nil
}