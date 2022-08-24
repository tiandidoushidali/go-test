package main

import (
	"fmt"
	. "github.com/panjf2000/gnet/v2"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	//client, err := NewClient(&DemoHandler{})
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	conn, _ := net.Dial("tcp4", "127.0.0.1:9000")
	go func() {
		for {
			var b []byte = make([]byte, 1024)
			n, err := conn.Read(b)
			if err == io.EOF {
				fmt.Println("connect is close")
				return
			}
			fmt.Println("r:", string(b), n)
		}
	}()
	for i := 0; i < 1000; i ++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("----i----", i, strconv.FormatInt(int64(i), 2), "===")
		conn.Write(([]byte)(strings.Repeat("hello", 5)+ strconv.FormatInt(int64(i), 10)+":"))
	}
}

type DemoHandler struct {

}


func (handler *DemoHandler) OnBoot(eng Engine) (action Action) {

	fmt.Println("OnBoot")
	return None
}

func (handler *DemoHandler) OnShutdown(eng Engine) {
	fmt.Println("OnShutdown")
}

func (handler *DemoHandler) OnOpen(c Conn) (out []byte, action Action) {

	fmt.Println("OnOpen")
	return []byte{}, None
}

func (handler *DemoHandler) OnClose(c Conn, err error) (action Action) {

	fmt.Println("OnClose")
	return Close
}

func (handler *DemoHandler) OnTraffic(c Conn) (action Action) {
	var b = make([]byte, 1024)
	n , err := c.Read(b)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("OnTraffic", string(b), n)

	c.Write(b)
	return None
}

func (handler *DemoHandler) OnTick() (delay time.Duration, action Action) {
	fmt.Println("OnTick")
	return time.Second * 10, None
}