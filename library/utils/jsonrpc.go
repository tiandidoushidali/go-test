package utils

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"reflect"
	"sync"
)

type JsonRpcSvr interface {
	SvrName() string
}

type Cal struct {

}

func (s *Cal) SvrName() string {
	return "cal-server"
}

func (s *Cal) Add(a, b int) (int, int, int) {
	return a, b, a + b
}

func (s *Cal) Sub(a, b int) int {
	return a - b
}

func init() {
	RegisterJsonRpc(new(Cal))
}

func RegisterJsonRpc(obj JsonRpcSvr) {
	once := sync.Once{}
	var cachee *cache.Cache
	once.Do(func() {
		cachee = cache.New(0, 0)
	})


	objMethods := make([]string, 0)
	ttyp := reflect.TypeOf(obj)
	for i := 0; i < ttyp.NumMethod(); i ++ {
		md := ttyp.Method(i)
		objMethods = append(objMethods, fmt.Sprintf("%s:%s", md.PkgPath, md.Name))
	}
	cachee.Set(fmt.Sprintf("jsonrpc:cache:%s", obj.SvrName()), objMethods, -1)
}

func JsonRpcCall(obj JsonRpcSvr, method string, params []interface{}) {

}
