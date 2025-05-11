package main

import _case "type_var/case"

func main() {

	//var i1 = 111
	//var f float32 = 23.4
	//var b1 = true
	//
	//var arr = [4]int32{1, 2, 2, 3}
	//
	//fmt.Println(i1, f, b1, arr)
	//
	////var i2 *int
	//var i2 int
	//f1(&i2) //&取地址值
	//
	//fmt.Println(i2)
	//
	//var inter interface{}
	//
	//inter = i1
	//fmt.Println(inter)
	//
	//_case.F1()
	//_case.New_case()
	_case.Map_case()
}

func f1(i *int) {
	*i = *i + 100
}
