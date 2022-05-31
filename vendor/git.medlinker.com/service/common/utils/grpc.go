package utils

import (
	"context"
	"encoding/json"
	"git.medlinker.com/golang/xtrace"
	"strings"

	"git.medlinker.com/service/common/cmodel"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

var (
	GrpcCtxTraceIdKey = xtrace.TraceIdNameOld // GRPC上下文中链路ID（旧版）
	GrpcCtxUserKey    = "user"                // GRPC上下文中user key
)

// 从 GRPC 上下文中获取IP
func GetIpForGRPCContext(ctx context.Context) string {
	pr, ok := peer.FromContext(ctx)
	if !ok || pr.Addr == nil {
		return ""
	}

	return strings.TrimSpace(pr.Addr.String())
}

// 从 GRPC 上下文中获取 User
func GetUserForGRPCContext(ctx context.Context) *cmodel.User {
	var user *cmodel.User
	u := ValueOfGRPCContext(ctx, GrpcCtxUserKey)
	if err := json.Unmarshal([]byte(u), &user); nil != err {
		return nil
	}

	return user
}

// 从 GRPC 上下文中获取 key 对应的值
func ValueOfGRPCContext(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ValueOfContext(ctx, key)
	}
	for _, data := range md.Get(key) {
		data = strings.TrimSpace(data)
		if data == "" {
			continue
		}
		return data
	}
	return ValueOfContext(ctx, key)
}

// 创建GRPC上下文，用于客户端发起GRPC请求使用。
// 便于统一GRPC请求参数封装、分布式链路跟踪参数封装。
func CreateGRPCContext(ctx context.Context) context.Context {
	newCtx := xtrace.CreateGRPCContext(ctx)
	// 注入用户信息
	if user := ValueOfGRPCContext(ctx, GrpcCtxUserKey); len(user) > 0 {
		// TODO 校验用户结构体
		newCtx = metadata.AppendToOutgoingContext(newCtx, GrpcCtxUserKey, user)
	}
	return newCtx
}

func ValueOfContext(ctx context.Context, key string) string {
	// 尝试从当前上下文取
	data := ctx.Value(key)
	if data == nil {
		return ""
	}
	if _, ok := data.(string); !ok {
		return ""
	}
	return data.(string)
}
