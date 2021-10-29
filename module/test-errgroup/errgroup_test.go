package test_errgroup

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey" // . 可以不用写包名 _ 只导入init
)



func TestHandle(t *testing.T) {
	testErrGroup := TestErrGroup{}
	Convey("测试handle", t , func() {
		So(testErrGroup.Handle(), ShouldNotBeNil)
	})
}
