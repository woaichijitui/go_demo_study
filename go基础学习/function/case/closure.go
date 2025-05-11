package _case

import (
	"fmt"
	"log"
)

// 斐波那契数列 X0+X1=X2...求Xn的值
func tool() func() int {
	x0, x1, x2 := 0, 1, 0

	//闭包
	//	闭包会使得外部函数的局部变量从栈里面到堆里
	//	闭包是:
	//	此返回的为函数，函数在后面加括号就是调用函数
	return func() int {
		x2 = x0 + x1
		x0 = x1
		x1 = x2
		return x2
	}
}
func Fib(n int) int {
	if n <= 2 {
		log.Fatal("n 不能小于")
	}
	// 函数若是没有入参，那么这个函数返回值是固定的 但是调用的次数会改变它
	t := tool()
	for i := 0; i < n-2; i++ {
		// tool 没调用一次
		fmt.Println(t())

	}
	return t()
}

// 模仿斐波那契数列 写一个1+2+...+n 值
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func Sum1(n int) {
	a := adder()
	for i := 1; i < n+1; i++ {

		fmt.Printf("1+2+...+%d = %d \n", i, a(i))
	}
}
func ClosureTrap() {
	//	错误方式
	/*for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)//会全部显示为10
		}()
	}*/
	//正确方式
	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println(j)
		}(i)
	}

}
