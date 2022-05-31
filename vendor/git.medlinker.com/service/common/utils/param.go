/**
 * Created by GoLand
 * Author: zhangjian
 * Date: 2019-08-15 13:50
 */

package utils

import (
	"net/url"
	"reflect"
	"strings"
)

// 从结构体生成对应的参数
// 统一只处理结构体和map
func CreateParams(req interface{}) string {
	sv := reflect.ValueOf(req)
	st := sv.Type()
	// 判断是否有效
	if st.Kind() == reflect.Ptr {
		if sv.IsNil() {
			return ""
		}
		sv = sv.Elem()
		st = sv.Type()
	}
	// 判断是否符合条件
	if st.Kind() != reflect.Struct && st.Kind() != reflect.Map {
		return ""
	}
	ps := url.Values{}
	// 结构体
	if st.Kind() == reflect.Struct {
		fn := sv.NumField()
		for i := 0; i < fn; i++ {
			// json tag
			sf := st.Field(i)
			name := sf.Tag.Get("json")
			if name != "" {
				name = strings.SplitN(name, ",", 2)[0]
				if name == "-" {
					continue
				}
			}
			if name == "" {
				name = ToSnake(sf.Name)
			}
			sfv := sv.Field(i)
			// NOTE 暂时只处理字符串(后续有需要可以自定义)
			if sfv.Kind() == reflect.Ptr {
				if sfv.IsNil() {
					continue
				}
				sfv = sfv.Elem()
			}
			if sfv.Kind() != reflect.String {
				continue
			}
			ps.Add(name, sfv.String())
		}
		return ps.Encode()
	}
	// map
	it := sv.MapRange()
	for it.Next() {
		sf := it.Key()
		if sf.Kind() != reflect.String {
			continue
		}
		sfv := it.Value()
		// NOTE 暂时只处理字符串(后续有需要可以自定义)
		if sfv.Kind() == reflect.Ptr {
			if sfv.IsNil() {
				continue
			}
			sfv = sfv.Elem()
		}
		if sfv.Kind() != reflect.String {
			continue
		}
		ps.Add(ToSnake(sf.String()), sfv.String())
	}
	return ps.Encode()
}
