package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
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

var c = make(chan int)

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("时间到")
			<-c
		default:
			fmt.Println("default")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	js := `{
 "actualPayable": 98.12,
 "amountPayable": 100,
 "code": "1234",
 "compatRecord": "7800000001",
 "orderNumber": "165465300992",
 "patCardId": "411502198612088781",
 "patMobile": "15868175923",
 "patName": "⻩俊辉",
 "patSex": "⼥",
 "strActualPayable": "98.00",
 "strAmountPayable": "100.00"
 }`
	type WindowOrderConfirmResp struct {
		ActualPayable decimal.Decimal `json:"actualPayable"` // 实际⽀付⾦额
		AmountPayable decimal.Decimal `json:"amountPayable"` // 总共⽀付⾦额
		Code string `json:"code"` // 4位缴费码
		CompatRecord string `json:"compatRecord"` // 就诊卡号或医保卡号
		OrderNumber string `json:"orderNumber"` // 订单号
		PatCardId string `json:"patCardId"` // 身份证号
		StrActualPayable string `json:"strActualPayable"` // 转换实际⽀付⾦额
		StrAmountPayable string `json:"strAmountPayable"` // 转换总共⽀付⾦额
		PatSex string `json:"patSex"` // 性别
		PatName string `json:"patName"` // patName
		PatMobile string `json:"patMobile"` // ⼿机号
	}
	var wo WindowOrderConfirmResp

	er1 := json.Unmarshal([]byte(js), &wo)
	if er1 != nil {
		fmt.Println("---err----", er1)
	}
	up := make(map[string]interface{})
	json.Unmarshal([]byte(js), &up)
	fmt.Println("----", up)
	fmt.Println(1 << 20)
	return
	fundDe, _ := decimal.NewFromString("1.036")
	v := fundDe.Mul(decimal.NewFromInt(100)).BigInt().Int64()
	fmt.Println("---v-----", v)
	return
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3 * time.Second)
	defer cancel()

	go watch(ctx)
	go watch(ctx)

	c <- 1
	fmt.Println("---")



	return
	var err error
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
	err = json.Unmarshal([]byte(j), &wrapper)
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