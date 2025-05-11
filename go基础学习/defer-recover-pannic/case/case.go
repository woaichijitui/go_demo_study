package _case

import (
	"fmt"
	"io"
	"log"
	"os"
)

// defer 执行循序为 后进先出
func DeferCase1() {
	fmt.Println("DeferCase1 函数开始")
	defer func() {
		fmt.Println("执行defer函数1")

	}()
	defer func() {
		f1()

	}()
	defer func() {
		fmt.Println("执行defer 匿名函数2")

	}()
	fmt.Println("DeferCase1 函数结束")
}

func f1() {
	fmt.Println("执行defer具名函数f1")
}

// 参数预计算
func DeferCase2() {

	var i = 32
	//	传参数
	//	defer后，传入的实参变量值是在defer函数之前的值
	defer func(i int) {
		fmt.Printf("defer 匿名函数1， i = %d\n", i)
	}(i + 1)
	//	闭包（没有参数）？
	defer func() {
		fmt.Printf("defer 匿名函数2， i = %d\n", i)
	}()

	i = 64
	defer func(i int) {
		fmt.Printf("defer 匿名函数3， i = %d\n", i)
	}(i + 1)
}

// 返回值
func DeferCase3() {
	i, j := f2()

	fmt.Println("i j g = ", i, *j, g)
}

var g = 100

func f2() (int, *int) {

	//defer 函数执行在return之后
	//若是传入指针，则也会相应改变
	defer func() {

		g = 200
	}()

	fmt.Printf("f2  ，g = %d\n", g)

	return g, &g
}

// recover
func RecoverCase() {

	fmt.Println("RecoverCase 开始")
	defer func() {

		// 	铺获异常，恢复协程
		err := recover()
		//	异常处理
		if err != nil {
			fmt.Println("异常处理 defer recover ", err)
		}
	}()
	panic("主动抛出的一个异常")
	fmt.Println("RecoverCase 结束")
}

// panic
func FileReadCase() {
	fmt.Println("FileReadCase 方法开始")
	file, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	// 释放资源
	defer file.Close()

	var buf = make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		fmt.Println(buf[:n])
	}
	fmt.Println("FileReadCase 方法结束")
}
