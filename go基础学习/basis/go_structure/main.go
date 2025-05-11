package main

import (
	"fmt"
	f1 "go_structure/func1"
)

func init() {
	fmt.Println("调用了 init 函数")
}
func main() {
	fmt.Println("调用了 main 函数")
	f1.F1()
}
