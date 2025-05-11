package _case

import "fmt"

func SimpleCase() {
	var a, b = 1, 2
	var c, d float64 = 3.3, 4.4
	//	编译器自动推导泛型类型
	fmt.Println("泛型数字比较: ", getMaxNum(a, b))
	//
	fmt.Println("泛型数字比较: ", getMaxNum[float64](c, d))

}

// 泛型的使用
func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// 自定义泛型类型接口
type cusnumT interface {

	//~表示该类型的衍生类型
	//	|表示并集
	//	换行表示交集
	int | uint8 | ~int64 | ~float64 | ~int32 | float32
	float64 | ~uint | ~int64 | int | ~int32 | float32
}

func CusNumCase() {
	var a, b = 1, 2
	var c, d float64 = 3.3, 4.4
	var e, f float32 = 5.5, 6.6
	var g, h myInt64 = 7, 8
	var i, j myInt32 = myInt32(g), myInt32(h)

	fmt.Println("自定义泛型数字比较: ", getMaxCusNum(a, b))
	fmt.Println("自定义泛型数字比较: ", getMaxCusNum(c, d))
	fmt.Println("自定义泛型数字比较: ", getMaxCusNum(e, f))
	fmt.Println("自定义泛型数字比较: ", getMaxCusNum(g, h))
	fmt.Println("自定义泛型数字比较: ", getMaxCusNum(i, j))

}

// 是int64的衍生类型，具有int64的新类型，但不是int64
type myInt64 int64

// 相当于int32的别名
type myInt32 = int32

func getMaxCusNum[T cusnumT](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// go自带的泛型类型（内置类型）comparable
func ComparableCase() {
	var a, b = "123", "123"

	fmt.Println("comparable :", comparableNum(a, b))

}

func comparableNum[T comparable](a, b T) bool {
	return a == b
}

// go自带的泛型类型（内置类型）any

//
