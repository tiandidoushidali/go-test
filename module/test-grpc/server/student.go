package main

import (
	"context"
	"fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/proto"
	test_student_v1 "go-test/module/test-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {

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

func (s *Server) StudentInfo(ctx context.Context, in *test_student_v1.StudentInfoReq) (resp *test_student_v1.StudentInfoResp, err error) {
	fmt.Printf("进来了")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	}
	fmt.Println("md:", md)
	if t, ok := md["color"]; ok {
		fmt.Println("color:", t)
	}
	return &test_student_v1.StudentInfoResp{
		Id:                   in.Id,
		Name:                 "abcdefg",
		Age:                  1234567,
	}, nil
}

func (s *Server) StudentList(ctx context.Context, in *test_student_v1.StudentListReq) (resp *test_student_v1.StudentListResp, err error) {
	return &test_student_v1.StudentListResp{
		List:                 make([]*test_student_v1.StudentEntity, 0),
		Total:                0,
	}, nil
}

func main() {
	// 坚挺本地端口
	lis, err := net.Listen("tcp", ":10002")
	if err != nil {
		fmt.Printf("监听端口失败：%+v", err)
		return
	}
	// 创建gRPC服务器
	s := grpc.NewServer()
	// 注册服务
	test_student_v1.RegisterStudentServer(s, &Server{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {

		fmt.Printf("开启服务失败:%+v", err)
	}
}
