package code

import (
	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	BusinessException = ecode.New(400)
	// 公用err 放这里，详细err 放具体目录里面
	Success  = &XCode{200, "成功"}
	ParamErr = &XCode{400, "参数 有误"}
	LoginErr = &XCode{401, "登录失败"}
	SessErr  = &XCode{401, "session 有误"}
	// passport相关错误
	UserBlockErr   = &XCode{10401, "该账号已被禁用！"}
	DataNotFindErr = &XCode{10001, "数据不存在！"}
	UserNotFindErr = &XCode{10002, "用户不存在！"}
	VerifyCodeErr  = &XCode{10021, "验证码错误！"}
)

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

// 判断是否是这个错误, 这个函数有问题，但是多方有依赖，先暂时放这里，如果要用，请使用 IsSameErr
//Deprecated
//func (e *XCode) IsThisErr(err error) bool {
//	codes, ok := errors.Cause(e).(ecode.Codes)
//	if !ok {
//		return false
//	}
//	if codes.Code() == e.ErrCode && codes.Message() == e.ErrMessage {
//		return true
//	}
//	return false
//}

// 判断是否是这个错误
func (e *XCode) IsSameErr(err error) bool {
	codes, ok := errors.Cause(err).(ecode.Codes)
	if !ok {
		return false
	}
	if codes.Code() == e.ErrCode && codes.Message() == e.ErrMessage {
		return true
	}
	return false
}

// 判断是否是错误
func IsStatusErr(err error, statusErr *ecode.Status) bool {
	codes, ok := errors.Cause(err).(ecode.Codes)
	if !ok {
		return false
	}
	if codes.Code() == statusErr.Code() {
		return true
	}
	return false
}

func (e *XCode) ToStatus() error {
	return status.Errorf(codes.Code(e.Code()), e.ErrMessage)
}