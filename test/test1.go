package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/golang/protobuf/proto"
	test_student_v1 "go-test/module/test-grpc/proto"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"time"
)

func main() {
	transport := &http2.Transport{
		TLSClientConfig:        &tls.Config{
			InsecureSkipVerify:          true,
		},
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			fmt.Println("进入链接:", network, addr)
			return net.Dial(network, addr)
		},
		AllowHTTP:	true,
		ReadIdleTimeout: 10 * time.Second,
		PingTimeout: 10 * time.Second,
		MaxHeaderListSize: 1000000000,
	}

	c := http.Client{
		Transport:     transport,
		Timeout:       60 * time.Second,
	}
	//header := new(http.Header)
	//header.Set("Content-Type", "application/grpc")
	//header.Set("user-agent", "grpc-go/1.29.1")
	header1 := map[string][]string{}
	header1["content-type"] = []string{"application/grpc"}
	header1["user-agent"] = []string{"grpc-go/1.29.1"}
	header1["te"] = []string{"trailers"}

	req1 := &test_student_v1.StudentInfoReq{
		Id:                   111,
	}

	bb, _ := proto.Marshal(req1)
	b := bytes.NewBuffer(bb)
	//b = bytes.NewBuffer([]byte(`{"id":2}`))
	//fmt.Println("----", header.Get("user-agent"))

	req, err := http.NewRequest("POST", "http://localhost:10002/test.student.v1.Student/StudentInfo", b)
	if err != nil {
		fmt.Println("request error :", err)
	}
	req.Header = header1
	req.Proto = "HTTP/2"
	req.Host = "localhost:10002"
	//ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	//req.WithContext(ctx)

	r, err := c.Do(req) // .Get("http://127.0.0.1:10002/test.student.v1.Student/StudentInfo")
	if err != nil {
		fmt.Printf("errrrrror:", err)
	}
	fmt.Println("rrrrrrrrrrrespon:", r)

	fmt.Println("aaaaaaaaaaaaa")
	//time.Sleep(10 * time.Second)
	fmt.Println("bbbbbbbbbbbbb")
}
