package main

import (
	"net"
	"net/rpc"
)

type Params struct {
	Weight, Width int
}

type Rect struct {

}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Weight
	return nil
}

func main() {
	rect := new(Rect)
	rpc.Register(rect)

	lis, err := net.Listen("tcp", "0.0.0.0:8899")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}

}
