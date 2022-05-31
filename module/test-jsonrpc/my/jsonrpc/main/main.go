package main

import (
	"encoding/json"
	"fmt"
	jsrpc "go-test/module/test-jsonrpc/my/jsonrpc"
)

func main() {

	jsrpcService := jsrpc.New()
	args := map[string]interface{}{
		"id": 123,
		"method": "add",
		"params": map[string]interface{}{
			"appKey": "123456",
			"ts": 12345678, // 时间戳
			"sign": "12344566",
			"payload": map[string]interface{}{
				"a":1,
				"b":2,
				"c":3,
			},
		},
		"jsonrpc":"2.0",
	}

	argss, _ := json.Marshal(args)
	var req jsrpc.JsonRpcReq
	json.Unmarshal([]byte(argss), &req)

	fmt.Println("----req-----", req, req.Params.Sign, req.Params.Payload)

	resp, err := jsrpcService.Call("budget", "Add", `{"a":1,"b":2, "c":1}`)
	fmt.Println("-----", resp, err)
	b, _ := json.Marshal(resp)
	fmt.Println("------b----", string(b))
	
}
