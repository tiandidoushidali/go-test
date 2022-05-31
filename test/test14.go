package main

import (
	"encoding/json"
	"fmt"
)

type CallResultWrapper struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Succ bool   `json:"succ"`
	*CallResult
}

//type CallResult struct {
//	IdentificationList []*ElectCodeOrderInfoResp_Entity `json:"list"`
//}
//
//func (c *CallResult) Entity() interface{} {
//	return ""
//}
type CallResult struct {
	Obj                interface{} `json:"obj"` // 窗口缴费
	List interface{} `json:"list"`
	DrugEntityList     interface{} `json:"drugEntityList"`
	ElectCodeOrderList interface{} `json:"electCodeOrderList"` // 电子医保码订单信息
	OrderPreSettleInfo interface{} `json:"orderPreSettleInfo"` // 订单预结算信息
	HealthInsurancePayInfo interface{} `json:"healthInsurancePayInfo"` // 支付信息
}

type ElectCodeOrderInfoResp_Entity struct {
	Amount string `json:"amount"`
	BillDate string `json:"billdate"`
	ExpId string `json:"expid"`
	PatId string `json:"patid"`
	PatName string `json:"patname"`
	Complete string `json:"complete"`
}

func (en *ElectCodeOrderInfoResp_Entity) String() string {
	return "测试的了"
}

func main() {
	j := `{
	"code": "0",
	"list": [{
		"amount": "4.50",
		"billdate": "2022-02-11 14:35:22",
		"expid": "620603c409d47b4625245eaa",
		"patid": "120357950",
		"patname": "杨⽂斌",
		"complete": "0"
	}],
	"msg": "根据电⼦医保码获取订单信息成功",
	"succ": true
}`

	wrapper := new(CallResultWrapper)
	callResult := new(CallResult)
	wrapper.CallResult = callResult
	err := json.Unmarshal([]byte(j), &wrapper)
	if err != nil {
		fmt.Println("err:", err)
	}

	e := callResult.List.([]interface{})

	////e := wrapper.List.([]*ElectCodeOrderInfoResp_Entity)
	//ee := e[0].(map[string]interface{})
	b, _ := json.Marshal(e)

	en := make([]*ElectCodeOrderInfoResp_Entity, 0)
	json.Unmarshal(b, &en)
	fmt.Println(en[0].Amount)

	fmt.Println(fmt.Sprintf("----en--%s--", en))
}

type TestIn interface {
	add() string
}

type aa  = TestIn