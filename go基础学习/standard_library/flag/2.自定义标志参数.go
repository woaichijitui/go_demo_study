package main

import (
	"flag"
	"fmt"
	"strings"
)

// 定义一个自定义类型
type ArrayFlag []string

// 实现flag.value的string方法
func (f *ArrayFlag) String() string {
	return strings.Join(*f, ", ")
}

// 实现flag.value的Set方法
func (f *ArrayFlag) Set(value string) error {

	*f = append(*f, value)
	return nil
}

func main() {
	var arrayFlag ArrayFlag

	//list
	flag.Var(&arrayFlag, "item", "List of items")

	flag.Parse()

	fmt.Println(arrayFlag)
	fmt.Println(len(arrayFlag))
}
