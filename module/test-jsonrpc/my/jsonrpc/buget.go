package jsrpc

import (
	"context"
)

type BudgetJsonRpc struct {

}

func (rpc *BudgetJsonRpc) SvrName() string {
	return "budget"
}
func NewBudgetJsonRpc() *BudgetJsonRpc {

	rpc := &BudgetJsonRpc{}
	return rpc
}

type BudgetJsonRpcAddReq struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type BudgetJsonRpcAddResp struct {
	Sum float64
}

func (rpc *BudgetJsonRpc) Add(ctx context.Context, req *BudgetJsonRpcAddReq) (resp *BudgetJsonRpcAddResp, err error) {
	resp = &BudgetJsonRpcAddResp{Sum: req.A+req.B}
	return
}

//func (rpc *BudgetJsonRpc) Add2(a, b float64) (float64, float64, float64, error) {
//
//	return a, b, a+b, errors.New("参数错误")
//}