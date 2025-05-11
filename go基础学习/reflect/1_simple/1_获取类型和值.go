package _simple

import (
	"fmt"
	"reflect"
)

// 判断一个变量是否为结构体，切片，map
func Reftype(obj any) {

	typeObj := reflect.TypeOf(obj)
	fmt.Println(typeObj, typeObj.Kind())

	// 去判断具体的类型
	switch typeObj.Kind() {
	case reflect.Slice:
		fmt.Println("切片")
	case reflect.Int:
		fmt.Println("数字")
	case reflect.Map:
		fmt.Println("map")
	case reflect.Struct:
		fmt.Println("结构体")
	case reflect.String:
		fmt.Println("字符串")
	}
}

// RefValue
//
//	@Description: 通过反射获取值
//	@param obj
func RefValue(obj any) {
	value := reflect.ValueOf(obj)
	fmt.Println(value, value.Type())
	switch value.Kind() {
	case reflect.Int:
		fmt.Println(value.Int())
	case reflect.Struct:
		fmt.Println(value.Interface())
	case reflect.String:
		fmt.Println(value.String())
	case reflect.Map:
		fmt.Println(value.MapRange())

	}
}
