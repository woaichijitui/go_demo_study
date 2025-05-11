package _simple

import (
	"fmt"
	"reflect"
)

// GetStructTag
//
//	@Description:读取json标签对应的值，如果没有就用属性的名称
func GetStructTag() {
	type User struct {
		Name string `json:"name"`
		Age  int
	}

	var user User = User{Age: 2, Name: "htt"}

	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		if tag == "" {
			//当没有json标签时使用字段名
			tag = t.Field(i).Name
		}

		fmt.Println(t, v, t.Field(i).Name, tag)

	}

}
