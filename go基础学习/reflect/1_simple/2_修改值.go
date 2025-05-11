package _simple

import (
	"fmt"
	"reflect"
)

// RefSetValue
//
//	@Description: 通过反射修改值
//	@param obj:	any
func RefSetValue(obj any) {
	valueObj := reflect.ValueOf(obj)
	fmt.Println(valueObj)
	//获取元素,传递地址就需要这操作
	elem := valueObj.Elem()

	fmt.Println(valueObj, elem)

	//
	switch elem.Kind() {
	case reflect.String:
		elem.SetString("htt")
		fmt.Println(valueObj, elem)
	}

}
