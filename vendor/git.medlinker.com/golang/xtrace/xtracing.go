// 分布式链路跟踪封装组件
package xtrace

import (
	"context"
	"fmt"
	"net/http"
)

const (
	SpanIdName     = "Span-Id"   // 分布式链路跟踪SpanId名称
	TraceIdName    = "Trace-Id"  // 分布式链路跟踪TraceId名称
	TraceMapName   = "Trace-Map" // 分布式链路跟踪TraceMap名称
	TraceIdNameOld = "trace_id"  // 分布式链路跟踪TraceId名称(旧版，短期保持兼容)
)

// 创建HTTP Header，用于客户端发起HTTP请求使用，便于统一HTTP请求的分布式链路跟踪参数封装。
// 参数 ctxOrHeader 可以为 http.Header/*http.Header/context.Context 类型。
// 注意返回的数据类型为map，以方便后续扩展。
func CreateHTTPHeader(ctxOrHeader interface{}) map[string]string {
	return GetTrace(ctxOrHeader).HTTPHeader()
}

// 创建GRPC上下文，用于客户端发起GRPC请求使用，便于统一GRPC请求的分布式链路跟踪参数封装。
// 返回的 context.Context 也支持标准库的Value方法获取上下文变量值，即返回值也可以用于模块、方法间传递上下文。
// 参数 ctxOrHeader 可以为 http.Header/*http.Header/context.Context 类型。
func CreateGRPCContext(ctxOrHeader interface{}) context.Context {
	return GetTrace(ctxOrHeader).GRPCContext()
}

// 创建标准上下文接口对象，用于模块、方法间传递链路信息。
// 参数 ctxOrHeader 可以为 http.Header/*http.Header/context.Context 类型。
func CreateStdContext(ctxOrHeader interface{}) context.Context {
	return GetTrace(ctxOrHeader).StdContext()
}

// 从Header/Context中获取指定键名的数据，返回的是一个字符串。
// 参数 ctxOrHeader 可以为 http.Header/*http.Header/context.Context 类型。
func GetValueFrom(ctxOrHeader interface{}, key string) string {
	switch v := ctxOrHeader.(type) {
	case context.Context:
		value := getValueFromStdContext(v, key)
		if value == "" {
			value = getValueFromGRPCContext(v, key)
		}
		return value
	case http.Header:
		return v.Get(key)
	case *http.Header:
		return v.Get(key)
	default:
		panic(fmt.Sprintf("unsupport parameter type: %t", ctxOrHeader))
	}
	return ""
}
