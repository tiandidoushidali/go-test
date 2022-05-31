package utils

import (
	"fmt"
	"reflect"
)

func TypeNameOf(v interface{}) string {
	dt := reflect.TypeOf(v)
	dv := reflect.ValueOf(v)
	// 找到未初始化的指针
	for dt.Kind() == reflect.Ptr {
		if dv.IsNil() {
			return ""
		}

		dv = dv.Elem()
		dt = dv.Type()
	}

	return dt.Name()
}

func FullTypeNameOf(v interface{}) string {
	rt := reflect.TypeOf(v)
	return fmt.Sprintf("%s/%s", rt.PkgPath(), TypeNameOf(v))
}
