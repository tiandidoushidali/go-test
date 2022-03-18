package grp_builder

import (
	"fmt"
	"google.golang.org/grpc/resolver"
)

func init()  {
	resolver.Register(&GrpcBuilder{})
}

type GrpcBuilder struct {
	
}

func (b *GrpcBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	r := &GrpcResolver{
		target: target,
		cc:     cc,
	}
	fmt.Println("---进入自定义builder---", target)
	r.ResolveNow(resolver.ResolveNowOption{})
	return r, nil
}

func (*GrpcBuilder) Scheme() string {
	return "grpcc"
}


type GrpcResolver struct {
	target resolver.Target
	cc resolver.ClientConn
}

func (r *GrpcResolver) ResolveNow(opts resolver.ResolveNowOption) {
	fmt.Println("----进入resolver-----", r.target)
	r.cc.NewAddress([]resolver.Address{{Addr: "127.0.0.1:10002"}, {Addr: "127.0.0.1:10003"}})
}

func (r *GrpcResolver) Close() {
	
}