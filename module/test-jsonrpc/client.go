package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type ParamsC struct {
	Weight, Width, Size int
}

func main() {
	c, err := rpc.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		panic(err)
	}

	for {
		ret := 0
		err := c.Call("Rect.Area", ParamsC{100, 100, 100}, &ret)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("--ret--", ret)
		c.Call("Rect.Area", ParamsC{200, 200, 100}, &ret)
		fmt.Println("--ret--", ret)

		time.Sleep(2 * time.Second)
	}
}
