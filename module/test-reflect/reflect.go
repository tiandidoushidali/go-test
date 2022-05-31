package main

import (
	"fmt"
	"reflect"
)

type JsonRpcRef interface {
	SvrName() string
}

type ReflectCal struct {
	Name string
}

func (s *ReflectCal) Add(a, b int) (int, int, int) {
	return a, b, a + b
}

func (s *ReflectCal) Sub(a, b int) int {
	return a - b
}

func (s *ReflectCal) Print(a, b, c int, str string) (int, string){
	return 1, "2"
}

func (s *ReflectCal) SvrName() string {
	return "cal"
}
func main() {
	cal := &ReflectCal{
		Name: "zs",
	}
	var cal1 *ReflectCal = &ReflectCal{
		Name: "ls",
	}

	*cal1 = *cal
	cal1.Name = "ls2"
	fmt.Println("----", &cal1, &cal, cal1, cal)

	fmt.Println("----", reflect.TypeOf(cal), reflect.ValueOf(cal))
	fmt.Println("----", reflect.TypeOf(cal).Name())

	for i := 0; i < reflect.TypeOf(cal).NumMethod(); i ++ {
		fmt.Println("----method---", reflect.TypeOf(cal).Method(i).Name, reflect.TypeOf(cal).Method(i), reflect.TypeOf(cal).Method(i).Type.NumIn())
	}

	//fmt.Println("----filed num---", reflect.TypeOf(cal).NumField())
	//for i := 0; i < reflect.TypeOf(cal).NumField(); i ++ {
	//	fmt.Println("----filed---", reflect.TypeOf(cal).Field(i))
	//}

	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf(123))
	args = append(args, reflect.ValueOf(234))

	values := reflect.ValueOf(cal).MethodByName("Add").Call(args)
	for k, v := range values {
		fmt.Println("----k---", k, "----value----", v)
	}

	obj := reflect.ValueOf(cal)
	fmt.Println("--", obj.Type().Implements(reflect.TypeOf((*JsonRpcRef)(nil)).Elem()))

	m , b := reflect.TypeOf(cal).MethodByName("SvrName")
	fmt.Println("----method-----", m, b)
	fmt.Println("----method22-----", m.Type, "-", m.Name, "-", m.PkgPath, "-", m.Func, "-", m.Index)

	m , b = reflect.TypeOf(cal).MethodByName("Sub")
	fmt.Println("----method-----", m, b)
	fmt.Println("----method22-----", m.Type, "-", m.Name, "-", m.PkgPath, "-", m.Func, "-", m.Index)

	m , b = reflect.TypeOf(cal).MethodByName("Print")
	fmt.Println("----method-----", m, b)
	fmt.Println("----method22-----", m.Type, "-", m.Name, "-", m.PkgPath, "-", m.Func, "-", m.Index)

	m , b = reflect.TypeOf(cal).MethodByName("aa")
	fmt.Println("----method-----", m, b)


}
