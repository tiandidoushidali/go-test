package main

import (
	"context"
	"fmt"
	"git.medlinker.com/golang/xerror"
	"github.com/go-kratos/kratos/pkg/ecode"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	test_student_v1 "go-test/module/test-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"time"
)

type Server struct {
	t int64

	grpcCredentialsAuth *GrpcCredentialsAuth
	grpcAuth            *GrpcAuth
}

type XCode struct {
	ErrCode    int
	ErrMessage string
}

func (e *XCode) Error() string {
	return e.ErrMessage
}

// Code return error code
func (e *XCode) Code() int { return e.ErrCode }

// Message return error message
func (e *XCode) Message() string {
	return e.ErrMessage
}

type GrpcCredentialsAuth struct {
	AppId     string
	AppSecret string
}

type GrpcAuth struct {
	Token string
}

func (auth *GrpcCredentialsAuth) Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "自定义认证token失败")
	}

	var (
		appId     string
		appSecret string
	)

	if values, ok := md["appId"]; ok {
		appId = values[0]
	}
	if values, ok := md["appSecret"]; ok {
		appSecret = values[0]
	}

	if appId == "" || appSecret == "" {
		return status.Errorf(codes.Unauthenticated, "自定义认证token失败")
	}

	yesAuthMap := make(map[string]*GrpcCredentialsAuth)
	yesAuthMap["test_id"] = &GrpcCredentialsAuth{
		AppId:     "test_id",
		AppSecret: "test_secret",
	}
	yesAuthMap["test_id1"] = &GrpcCredentialsAuth{
		AppId:     "test_id1",
		AppSecret: "test_secret1",
	}

	yesAuthInfo, appIdOk := yesAuthMap[appId]
	// 准入的appId appSecret
	if appIdOk && yesAuthInfo.AppSecret == appSecret {
		return nil
	}
	return status.Errorf(codes.Unauthenticated, "自定义认证token失败")
}

func (auth *GrpcAuth) Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "自定义认证token失败")
	}

	var (
		token string
	)

	if values, ok := md[":authority"]; ok {
		token = values[0]
	}

	if token == "" {
		return status.Errorf(codes.Unauthenticated, "自定义认证token失败")
	}

	yesTokens := []string{"test_api", "test_id"}
	for _, tk := range yesTokens {
		if tk == token {
			return nil
		}
	}
	return status.Errorf(codes.Unauthenticated, "自定义认证token失败")
}
func (s *Server) StudentInfo(ctx context.Context, in *test_student_v1.StudentInfoReq) (resp *test_student_v1.StudentInfoResp, err error) {
	err = s.grpcAuth.Check(ctx)
	if err != nil {
		fmt.Println("---grpc auth check fail----", err)
		return
	}
	fmt.Printf(fmt.Sprintf("进来了%+v", in))
	fmt.Println("----ctx-----", ctx)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	}
	fmt.Println("md:", md)
	fmt.Println("md:", md.Get("aaaaaaaaaaaa"))
	if t, ok := md["color"]; ok {
		fmt.Println("color:", t)
	}
	time.Sleep(1 * time.Second)
	ss := &test_student_v1.StudentInfoResp{
		Id:   in.Id,
		Name: "abcdefg",
		Age:  s.t,
	}
	if in.Id == 3 {
		fmt.Println("进入等待", in.Id)
		time.Sleep(2 * time.Second)
		s.t = 1111111
		fmt.Println("进入等待2", in.Id)
	}
	if in.Id == 4 {
		fmt.Println("进入等待", in.Id)
		time.Sleep(5 * time.Second)
		fmt.Println("进入等待2", in.Id)
	}

	//st := status.New(codes.Code(1000), "错误了1111")

	e := xerror.WrapWithCode(errors.New("111111"), 1000000)
	xe := e.(*xerror.Error)
	xe.WithFields(1000, "错误了1111")
	fmt.Println("----===--", ecode.Code(1000).Code())

	return ss, status.New(codes.Code(1000), "错误了1111").Err()
	return ss, status.Errorf(codes.Code(1000), "错误了1111")
	//return &test_student_v1.StudentInfoResp{
	//	Id:                   in.Id,
	//	Name:                 "abcdefg",
	//	Age:                  s.t,
	//}, nil
}


func (s *Server) StudentList(ctx context.Context, in *test_student_v1.StudentListReq) (resp *test_student_v1.StudentListResp, err error) {
	return &test_student_v1.StudentListResp{
		List:  make([]*test_student_v1.StudentEntity, 0),
		Total: s.t,
	}, nil
}


func (s *Server) StudentList2(ctx context.Context, in *test_student_v1.StudentList2Req) (resp *test_student_v1.StudentListResp, err error) {
	for _, vs := range in.Data {
		for k1, v1 := range vs.Entitys {
			fmt.Println("----k1----", k1, "----v1----", v1)
		}
	}
	return &test_student_v1.StudentListResp{
		List:  make([]*test_student_v1.StudentEntity, 0),
		Total: s.t,
	}, nil
}


func main() {
	// 监听本地端口
	lis, err := net.Listen("tcp", ":10002")
	if err != nil {
		fmt.Printf("监听端口失败：%+v", err)
		return
	}

	// 使用拦截器
	var serverIntercept grpc.UnaryServerInterceptor
	serverIntercept = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("in serverIntercept")
		return handler(ctx, req)
	}
	// 创建gRPC服务器
	s := grpc.NewServer(grpc.UnaryInterceptor(serverIntercept))
	// 注册服务
	test_student_v1.RegisterStudentServer(s, &Server{
		t:        time.Now().Unix(),
		grpcAuth: new(GrpcAuth),
	})

	//reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {

		fmt.Printf("开启服务失败:%+v", err)
	}
}
