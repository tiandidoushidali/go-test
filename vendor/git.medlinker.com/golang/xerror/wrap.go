package xerror

import (
	"git.medlinker.com/golang/xerror/xcode"
	"github.com/pkg/errors"
)

func Wrap(err error, message string) XError {
	// 如果本身是 XError 则包装 error
	// 偷个懒，直接用 pkg/errors 来包装
	if xe, ok := err.(XError); ok {
		err = xe.Unwrap()
	}

	return &Error{
		code:  xcode.NewUseMessage(message),
		error: errors.Wrap(err, message),
	}
}

func WrapWithCode(err error, code int) XError {
	return WrapWithXCode(err, xcode.New(code))
}

func WrapWithXCode(err error, code xcode.XCode) XError {
	if nil == code {
		code = xcode.InternalServerError
	}

	if xe, ok := err.(XError); ok {
		err = xe.Unwrap()
	}

	return &Error{
		code:  code,
		error: errors.Wrap(err, code.String()),
	}
}
