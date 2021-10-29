package test_mongodb

import (
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey" // . 可以不用写包名 _ 只导入init
)

// -- go test -v
var testMongodb *TestMongodb
func init() {
	fmt.Println("-----")
	testMongodb = &TestMongodb{}
}

func TestHandle(t *testing.T) {
	Convey("测试handle", t , func() {
		So(testMongodb.Handle(), ShouldBeNil)
	})
}

