package main

import (
	"context"
	"fmt"
	_case "function/case"
	"os"
	"os/signal"
)

func main() {
	//	调用函数
	fmt.Println(_case.Sum(1, 2))
	// 将函数赋值给变量
	f1 := _case.Sum
	fmt.Println(f1(1, 2))
	//	将函数作为输入输出实现中间件
	f2 := _case.LogMiddleWare2(_case.Sum)
	//	再次附加中间件
	f2 = _case.LogMiddleWare2(f2)
	fmt.Println(f2(1, 2))

	//	函数类型方法的调用
	//	声明一个函数类型的对象
	f3 := _case.Sumfunc(f1)
	fmt.Println(f3.Accomulation(1, 2, 3, 4, 5))

	// 需要中间返回的是SumFunc类型才能转
	fmt.Println(f2.Accomulation(1, 2, 3, 4, 5, 6))

	// 斐波那契数列
	_case.Fib(3)
	_case.Sum1(5)
	_case.ClosureTrap()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
