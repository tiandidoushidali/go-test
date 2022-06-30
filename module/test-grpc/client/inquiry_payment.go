package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	med_inquiry "go-test/module/test-grpc/proto-med/inquiry"
)

func main() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial("172.25.14.145:8000", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := med_inquiry.NewPaymentClient(conn)
	ctx := context.Background()
	resp, err := c.UpdateYanHuaHealthInsurancePayLog(ctx, &med_inquiry.UpdateYanHuaHealthInsurancePayLogReq{
		Id:                   1,
		PayCode:              "1111",
		LoginUserId:          "1111",
		ReadyPayNum:          "111100010B220609004179",
		PersonCountPay:       "0.00",
	})
	if err != nil {
		fmt.Println("----", err)
	}
	fmt.Println("-----resp-----", resp)
}
