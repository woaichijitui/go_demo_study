package huidiao

import "fmt"

type Callback func(i, j int) int

// 提供接口
func Huidiao2(a, b int, callback Callback) int {
	return callback(a, b)
}

//回调函数的具体实现

func calculationXOR(i, j int) int {
	return i * j
}

func calculationAnd(i, j int) int {
	return i + j
}

func Test2() {

	fmt.Println(Huidiao2(11, 12, calculationAnd))

	fmt.Println(Huidiao2(11, 12, calculationXOR))
}
