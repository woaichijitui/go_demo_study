package main

import (
	_case "basis/var-func/case"
	"fmt"
)

func main() {
	//局部变量
	a := 10
	b := 20
	//实参
	fmt.Println(_case.Add(a, b))
	fmt.Println(b)
	//	引用类型传递
	fmt.Println(_case.Add1(a, &b))
	fmt.Println(b)

	//	全局变量的使用
	fmt.Println(_case.G)

	//	结构体中的方法
	user := _case.NewUser()
	fmt.Println(user.GetName(), user.GetAge())
}
