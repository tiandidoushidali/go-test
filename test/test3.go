package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-test/grape/protocol/message"
	"path"
	"time"
)

func main() {

	//var val int64 = 10000
	//t := time.Now().Nanosecond()
	//for i := int64(0); i < 100; i ++ {
	//	go func(i int64) {
	//		val = val +i
	//		//atomic.AddInt64(&val, -i)
	//	}(i)
	//}
	//
	//for i := int64(0); i < 100; i ++ {
	//	go func(i int64) {
	//		val = val +i
	//		//atomic.AddInt64(&val, -i)
	//	}(i)
	//}
	//
	//fmt.Println("========", time.Now().Nanosecond() - t)
	//
	//time.Sleep(2 * time.Second)
	//fmt.Println("-------", val)
	//return

	fmt.Println("path base :", path.Base("/aa/bb/cc/"))
	data := "CICAte7Mwe7iEKIBEzEyMDg1NzY2MzkxNzQwMDA2NDASSgixmKUlGg7ogZTlj4s3ODIwMzk1MyIxaHR0cDovL3B1Yi1tZWQtYXZhdGFyLmltZ3MubWVkbGlua2VyLm5ldC9tYWxlLnBuZyghIuABCNOMuAEaDOWSqOivouWvueivnSIxaHR0cDovL3B1Yi1tZWQtYXZhdGFyLmltZ3MubWVkbGlua2VyLm5ldC9tYWxlLnBuZyKGAWh0dHA6Ly90aGlyZHd4LnFsb2dvLmNuL21tb3Blbi92aV8zMi9EWUFJT2dxODNlbzBFaWFlQkc1WFF0YUI1YU82OG4xbjJ6WEdpY2VvaWFpY2FVeVlOZFhtZGVta0xvaWEyOXd1b0hWNnh5WmliRVI0TnRpY21BSGdNZ0t6UFM1dHcvMTMyKAIwsZilJTgGQMSU07CovStSRgpE6K6i5Y2V5bCG5LuO6K6i5Y2V5Yib5bu65pe26Ze05byA5aeL6K6h5pe277yMMzDliIbpkp/lkI7oh6rliqjlhbPpl60="
	decoded, _ := base64.StdEncoding.DecodeString(data)
	msg := &message.Message{}
	proto.Unmarshal(decoded, msg)
	fmt.Println("msg:", msg.Id, msg.To, msg.From, msg.GetBusinessCard())

	timeScope := "2012-12-12"
	times, _ := time.ParseInLocation("2006-01-02", timeScope, time.Local)
	fmt.Println(times.String())

	type aa struct{
		Name string
	}

	aas := make([]*aa, 0)
	aas = append(aas, &aa{Name: "zs"})
	aas = append(aas, &aa{Name: "ls"})

	mp := make(map[string]string)
	mp["zs"] = "zs"
	aasj, _ := json.Marshal(aas)

	fmt.Println("aaas json:", string(aasj))

}


