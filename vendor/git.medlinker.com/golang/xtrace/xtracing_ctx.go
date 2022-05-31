package xtrace

import "context"

// 创建用于链路信息传递的context.Context标准库接口对象。
// 参数 ctxOrHeader 可以为
// http.Header/*http.Header
// gin.Context/*gin.Context (实现了 apiGetHeader 接口)
// context.Context
// 类型。
func Ctx(ctxOrHeader interface{}) context.Context {
	return CreateGRPCContext(ctxOrHeader)
}
