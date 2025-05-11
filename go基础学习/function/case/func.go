package _case

import (
	"errors"
	"log"
)

func Sum(a, b int) (sum int, err error) {
	if a <= 0 && b <= 0 {
		return 0, errors.New("a和b 不能同时为小于零的数")
	}

	return a + b, nil
}

// 将函数定义为类型
type Sumfunc func(a, b int) (sum int, err error)

// 将函数作为输入输出实现中间件
/*func LogMiddleWare1(in func(a, b int) (int, error)) func(a, b int) (int, error) {
	//	返回函数为闭包函数,其中输入值为闭包函数使用的外部函数内部变量??
	return func(a, b int) (int, error) {
		log.Printf("日志中间件，记录操作数： a = %d.b = %d ", a, b)
		return in(a, b)
	}

}*/

// 将函数作为输入输出实现中间件 函数类型版
func LogMiddleWare2(in Sumfunc) Sumfunc {
	//	返回函数为闭包函数,其中输入值为闭包函数使用的外部函数内部变量??
	return func(a, b int) (int, error) {
		log.Printf("日志中间件，记录操作数： a = %d.b = %d ", a, b)
		return in(a, b)
	}
}

// 声明receiver为函数类型的方法.既函数类型的对象方法
func (s Sumfunc) Accomulation(list ...int) (sum int, err error) {
	for _, val := range list {
		//	这是什么意思
		sum, err = s(sum, val)
	}
	return sum, err
}
