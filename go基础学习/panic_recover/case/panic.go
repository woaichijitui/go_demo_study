package _case

import (
	"fmt"
)

func PanicCase() {
	foo()
	fmt.Println("reture normally from PanicCase func")
	deferCase()
}

func foo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in foo :", err)
		}
	}()

	fmt.Println("calling bar")
	bar(0)
	fmt.Println("return normol from bar ")
}

//

func bar(i int) {

	if i >= 3 {
		panic("err occurred")
	}

	defer fmt.Println("defer in bar : ", i)

	fmt.Println("println in bar :", i)
	bar(i + 1)
}

// 若不铺获 defer不会触发
func deferCase() {
	panic("pannic")

	defer fmt.Println("defer println ")

}
