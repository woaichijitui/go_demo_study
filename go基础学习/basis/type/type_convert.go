package _type

import (
	"fmt"
	"strconv"
	time2 "time"
	"unsafe"
)

func ConvertCase() {
	//	同一类型数据之间的转换 数字和数字、字符串和字符、字节
	//	不同类型数据的转换 数字和字符串
	//接口类型转其他类型

	//数字类型转换
	var int1 = 32
	int2 := int64(int1)
	fmt.Println(int2)

	//	字符串和数字
	//	包括各种进制的情况
	var str1 = "12"
	int1, _ = strconv.Atoi(str1)
	str1 = strconv.Itoa(int1)
	fmt.Println(int1, str1)

	fmt.Println(strconv.FormatInt(int2, 16))   //转成16进制的字符串
	fmt.Println(strconv.ParseInt(str1, 8, 64)) //转成16进制,大小是多大的字符

	//	字符串和[]byte
	var str2 = "htt123"
	byte1 := []byte(str2)
	fmt.Println(byte1)

	fmt.Println(string(byte1))

	//	rune
	r1 := []rune(str2)
	fmt.Println(r1)
	fmt.Println(string(r1))
	fmt.Println(r1[1])
	fmt.Println(string(r1[1]))

	// 	interface 与 其他类型
	var inter interface{} = 100
	var inter1 interface{} = User{
		Name:  "htt",
		Age:   26,
		Class: "1",
	}
	//	断言
	int3, err1 := inter.(int) //断言
	user, err := inter1.(User)
	fmt.Println(int3, err1)
	fmt.Println(user, err)

	//	字符产和时间
	time1 := time2.Now()
	fmt.Println(time1)
	timeFormat := time1.Format("2006.01.02 15:04:05Z07:00")

	fmt.Println(timeFormat)
	time1, err2 := time2.Parse("2006.01.02 15:04:05Z07:00", timeFormat)
	fmt.Println(time1, err2)

	// unitptr
	// unsafe.Pointer 是一个通用的指针类型，不可以参与计算
	u := new(User)
	uPtr := unsafe.Pointer(u)
	namePtr := unsafe.Pointer(uintptr(uPtr) + unsafe.Offsetof(u.Name))
	*(*string)(namePtr) = "htt"
	fmt.Println(u)

}
