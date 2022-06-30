package main

import (
	"context"
	"encoding/json"
	"fmt"
	"git.medlinker.com/service/common/ecode"
	"git.medlinker.com/service/grpcwrapper"
	"go-test/library/code"
	med_wallet_go "go-test/module/test-grpc/proto-med/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"reflect"
)
type User struct {
	// 用户ID：med_primary.user.id
	Id uint32
	// 用户类型
	Type uint32
	// 用户姓名
	FullName string
}

type XCode struct {
	ErrCode    int
	ErrMessage string
}

func NewXCode(code int, msg string) *XCode {
	return &XCode{
		ErrCode:    code,
		ErrMessage: msg,
	}
}

func (e *XCode) SetMess(mess string) *XCode {
	e.ErrMessage = mess
	return e
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

// Details return details.
func (e *XCode) Details() []interface{} { return nil }

func (e *XCode) ToGrpcStatus() error {
	return status.Errorf(codes.Code(e.Code()), e.ErrMessage)
}

func main() {

	////通过grpc 库 建立一个连接
	//conn, err := grpc.Dial("172.25.5.67:8000", grpc.WithInsecure())

	//conn, err := grpc.DialContext(context.Background(), "127.0.0.1:7001", grpc.WithInsecure(), grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	if s, ok := err.(*ecode.Status); ok {
	//		fmt.Println("------client recv----", s.Code(), s.Message())
	//	}
	//	fmt.Println("---client----", reflect.TypeOf(err))
	//	return err
	//}))
	conn, err := grpcwrapper.NewClient("127.0.0.1:7001")

	ctx := context.Background()
	u := &User{
		Id:       17736010,
		Type:     1,
		FullName: "彭航娅",
	}
	ubyte, _ := json.Marshal(u)
	fmt.Println("---", string(ubyte))
	md := metadata.Pairs("user",  string(ubyte))
	ctx = metadata.NewOutgoingContext(ctx, md)
	if err != nil {
		return
	}
	defer conn.Close()
	fmt.Println("-----")
	c := med_wallet_go.NewBankClient(conn)
	// 绑卡
	bindResp, err := c.BindBankCard(ctx, &med_wallet_go.BindBankCardReq{
		CardNo:               "6212263602036640662",
		IdCard:               "530181199008302627",
		RealName:             "彭航娅",
	})
	if err != nil {
		if eee, ok := err.(*ecode.Status); ok {
			fmt.Println("--eeee----", eee.Code(), eee.Message(), eee.Details())
		}
		fmt.Println("----err---", reflect.TypeOf(err))
		s := status.Convert(err)
		fmt.Println("---s---", s.Code(), s.Message(), s.Details())

		//sl := s.Details()
		//if st, ok := sl[0].(*types.Status); ok {
		//
		//	fmt.Println("----st----", st)
		//
		//}
		//fmt.Println("c.BindBankCard err ", sl[0])
	}
	fmt.Println("bindResp:", bindResp)
	return
	//// 解绑
	//unbindResp, err := c.UnbindBankCard(ctx, &med_wallet_go.UnbindBankCardReq{
	//	Id:                   4,
	//})
	//if err != nil {
	//	s := status.Convert(err)
	//	fmt.Println("c.UnbindBankCard err ", err, s.Message(), s.Code())
	//}
	//fmt.Println("unbindResp:", unbindResp)

	resp, err := c.GetUserCardList(ctx, &med_wallet_go.GetUserCardListReq{
		Page:                 1,
		Limit:                10,
	})
	if err != nil {
		e := err.(*code.XCode)
		fmt.Println("--e--", e.ErrCode, e.ErrMessage)
	}

	fmt.Println("----resp.List----", resp)
	//for _, v := range resp.List {
	//	data := v
	//	decoded, _ := base64.StdEncoding.DecodeString(data)
	//	//fmt.Println("---decoded-----", decoded)
	//	msg := &message.Message{}
	//	proto.Unmarshal(decoded, msg)
	//	if strings.Contains(msg.GetJson().GetContent(), "2204249163150638") {
	//
	//		b, _ := json.Marshal(msg)
	//		fmt.Println(string(b))
	//	}
	//}
	//fmt.Println("-----resp-----", resp.Index)
}