package main

import (
	"fmt"
	. "github.com/panjf2000/gnet/v2"
	"time"
)
func main() {
	srv := &DemoHandler{}
	Run(srv, "tcp4://127.0.0.1:9000")
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
	return []byte("OnOpen"), None
}

func (handler *DemoHandler) OnClose(c Conn, err error) (action Action) {

	fmt.Println("OnClose")
	return Close
}

func (handler *DemoHandler) OnTraffic(c Conn) (action Action) {
	var b = make([]byte, 24)
	n , err := c.Read(b)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("OnTraffic\n", string(b), n)

	c.Write(b)
	return None
}

func (handler *DemoHandler) OnTick() (delay time.Duration, action Action) {
	fmt.Println("OnTick")
	return time.Second * 10, None
}