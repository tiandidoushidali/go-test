/**
 * Created by GoLand
 * Author: zhangjian
 * Date: 2019-06-13 14:36
 */

/**
 * 1 为什么common库中有errors包
 *   a: pkg/errors中的Wrap/Wrapf方法会调用堆栈，当业务层层往上包装时有一定的性能开销
 *   b: 目前项目中已经在使用Wrap/Wrapf，保持函数名并调用pkg/errors的WithMessage方法
 *
 * 2 注意点
 *   a: 只要在明确需要堆栈的才调用Wrap/Wrapf方法（如：panic）
 */

package errors

import (
	err "errors"
	"fmt"

	pkgerr "github.com/pkg/errors"
)

// New returns an error with the supplied message.
func New(message string) error {
	return err.New(message)
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// NOTE WithMessage
// Wrap annotates err with a new message.
// If err is nil, WithMessage returns nil.
func Wrap(err error, message string) error {
	return pkgerr.WithMessage(err, message)
}

// NOTE WithMessagef
// Wrapf annotates err with a new message.
// If err is nil, WithMessage returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	return pkgerr.WithMessagef(err, format, args...)
}
