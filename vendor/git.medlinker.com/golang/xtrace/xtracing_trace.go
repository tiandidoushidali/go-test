package xtrace

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gconv"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

// 链路跟踪信息管理对象
type Trace struct {
	SpanId      string         // 层级ID
	TraceId     string         // 链路ID
	TraceMap    gmap.StrAnyMap // 应用自定义Key-Value信息，类似于OpenTracing中的annotation
	ctxOrHeader interface{}    // 数据来源
}

// Header信息接口，用于断言判断
type apiGetHeader interface {
	GetHeader(key string) string
}

// 根据Context或者Header获取链路跟踪信息。
// 参数 ctxOrHeader 可以为 http.Header/*http.Header/context.Context 类型。
func GetTrace(ctxOrHeader interface{}) *Trace {
	// 如果给定的参数为nil，返回一个带默认值的Trace对象
	if ctxOrHeader == nil {
		return new(Trace)
	}

	// 常见接口断言判断，例如：*gin.Context
	if v, ok := ctxOrHeader.(apiGetHeader); ok {
		trace := &Trace{
			SpanId:      v.GetHeader(SpanIdName),
			TraceId:     v.GetHeader(TraceIdName),
			ctxOrHeader: ctxOrHeader,
		}
		if s := v.GetHeader(TraceMapName); s != "" {
			if err := json.Unmarshal([]byte(s), &trace.TraceMap); err != nil {
				panic(err)
			}
		}
		return trace
	}

	switch v := ctxOrHeader.(type) {
	case context.Context:
		trace := &Trace{
			SpanId:      getValueFromStdContext(v, SpanIdName),
			TraceId:     getValueFromStdContext(v, TraceIdName),
			ctxOrHeader: ctxOrHeader,
		}
		if trace.SpanId == "" {
			trace.SpanId = getValueFromGRPCContext(v, SpanIdName)
		}
		if trace.TraceId == "" {
			trace.TraceId = getValueFromGRPCContext(v, TraceIdName)
		}
		// TODO 兼容旧版的TraceId, 未来删除
		if trace.TraceId == "" {
			trace.SpanId = getValueFromGRPCContext(v, TraceIdNameOld)
		}
		// TraceMap
		s := getValueFromStdContext(v, TraceMapName)
		if s == "" {
			s = getValueFromGRPCContext(v, TraceMapName)
		}
		if s != "" {
			err := json.Unmarshal([]byte(s), &trace.TraceMap)
			if err != nil {
				panic(err)
			}
		}
		return trace

	case http.Header:
		trace := &Trace{
			SpanId:      v.Get(SpanIdName),
			TraceId:     v.Get(TraceIdName),
			ctxOrHeader: ctxOrHeader,
		}
		if s := v.Get(TraceMapName); s != "" {
			if err := json.Unmarshal([]byte(s), &trace.TraceMap); err != nil {
				panic(err)
			}
		}
		return trace

	case *http.Header:
		trace := &Trace{
			SpanId:      v.Get(SpanIdName),
			TraceId:     v.Get(TraceIdName),
			ctxOrHeader: ctxOrHeader,
		}
		if s := v.Get(TraceMapName); s != "" {
			if err := json.Unmarshal([]byte(s), &trace.TraceMap); err != nil {
				panic(err)
			}
		}
		return trace

	default:
		panic(fmt.Sprintf("invalid parameter: %v", ctxOrHeader))
	}
	return nil
}

// 创建HTTP Header，用于客户端发起HTTP请求使用，便于统一HTTP请求的分布式链路跟踪参数封装。
// 注意返回的数据类型为map，以方便后续扩展。
func (t *Trace) HTTPHeader() map[string]string {
	header := make(map[string]string)
	// 将原本Header中的数据获取到新创建的header map中
	switch value := t.ctxOrHeader.(type) {
	case http.Header:
		for k, v := range value {
			if len(v) > 0 {
				header[k] = v[0]
			}
		}
	case *http.Header:
		for k, v := range *value {
			if len(v) > 0 {
				header[k] = v[0]
			}
		}
	}
	// 只有当链路跟踪信息不为空时才传递
	if t.SpanId != "" {
		header[SpanIdName] = t.SpanId
	}
	if t.TraceId != "" {
		header[TraceIdName] = t.TraceId
	}
	if t.TraceMap.Size() > 0 {
		result, err := t.TraceMap.MarshalJSON()
		if err != nil {
			panic(err)
		}
		header[TraceMapName] = gconv.UnsafeBytesToStr(result)
	}
	return header
}

// 创建GRPC上下文，用于客户端发起GRPC请求使用，便于统一GRPC请求的分布式链路跟踪参数封装。
// 返回的 context.Context 也支持标准库的Value方法获取上下文变量值，即返回值也可以用于模块、方法间传递上下文。
func (t *Trace) GRPCContext() context.Context {
	ctx := t.StdContext()
	kvs := make([]string, 0)
	// 如果原本存在请求的metadata数据，那么不能覆盖，采用追加数据；否则新建metadata数据
	md, _ := metadata.FromOutgoingContext(ctx)
	if md != nil {
		if md.Get(SpanIdName) == nil && t.SpanId != "" {
			kvs = append(kvs, SpanIdName, t.SpanId)
		}
		if md.Get(TraceIdName) == nil && t.TraceId != "" {
			kvs = append(kvs, TraceIdName, t.TraceId)
		}
		// TODO 兼容旧版的TraceId, 未来删除
		if md.Get(TraceIdNameOld) == nil && t.TraceId != "" {
			kvs = append(kvs, TraceIdNameOld, t.TraceId)
		}
		if t.TraceMap.Size() > 0 {
			// 注意TraceMapName在StdContext已经序列化
			kvs = append(kvs, TraceMapName, gconv.String(ctx.Value(TraceMapName)))
		}
	} else {
		if t.SpanId != "" {
			kvs = append(kvs, SpanIdName, t.SpanId)
		}
		if t.TraceId != "" {
			kvs = append(kvs, TraceIdName, t.TraceId)
			// TODO 兼容旧版的TraceId, 未来删除
			kvs = append(kvs, TraceIdNameOld, t.TraceId)
		}
		if t.TraceMap.Size() > 0 {
			// 注意TraceMapName在StdContext已经序列化
			kvs = append(kvs, TraceMapName, gconv.String(ctx.Value(TraceMapName)))
		}
	}
	if len(kvs) > 0 {
		ctx = metadata.AppendToOutgoingContext(ctx, kvs...)
	}
	return ctx
}

// 创建标准上下文接口对象，用于模块、方法间传递链路信息。
func (t *Trace) StdContext() context.Context {
	var ctx context.Context
	if v, ok := t.ctxOrHeader.(context.Context); ok {
		ctx = v
	} else {
		ctx = context.Background()
	}
	if ctx.Value(SpanIdName) == nil && t.SpanId != "" {
		ctx = context.WithValue(ctx, SpanIdName, t.SpanId)
	}
	if ctx.Value(TraceIdName) == nil && t.TraceId != "" {
		ctx = context.WithValue(ctx, TraceIdName, t.TraceId)
	}
	if ctx.Value(TraceMapName) == nil && t.TraceMap.Size() > 0 {
		result, err := t.TraceMap.MarshalJSON()
		if err != nil {
			panic(err)
		}
		ctx = context.WithValue(ctx, TraceMapName, gconv.UnsafeBytesToStr(result))
	}
	return ctx
}

// 这里的键值经过约定必定为string类型，否则返回空字符串
func getValueFromStdContext(ctx context.Context, key string) string {
	return gconv.String(ctx.Value(key))
}

// 从GRPC上下文中获取key对应的键值字符串
func getValueFromGRPCContext(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md, ok = metadata.FromOutgoingContext(ctx)
		if !ok {
			return ""
		}
	}
	for _, data := range md.Get(key) {
		data = strings.TrimSpace(data)
		if data == "" {
			continue
		}
		return data
	}
	return ""
}
