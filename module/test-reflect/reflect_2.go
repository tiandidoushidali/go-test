package main

import (
	"fmt"
	"reflect"
)

func main() {
	cal := &ReflectCal{}

	fmt.Println("----", reflect.TypeOf(cal), reflect.ValueOf(cal))
	fmt.Println("----", reflect.TypeOf(cal).Name())

	for i := 0; i < reflect.TypeOf(cal).NumMethod(); i ++ {
		fmt.Println("----method---", reflect.TypeOf(cal).Method(i))
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

	reflect.New(reflect.TypeOf(cal))
}
