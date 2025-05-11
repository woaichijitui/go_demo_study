package main

import (
	"flag"
	"fmt"
)

func main() {
	//		定义标志参数,得到一个相应类型的指针
	//	name：命令行使用的参数名字
	//	value：对应的值
	//	usage:	说明

	//	若命令行没有键入参数 则使用默认参数
	s := flag.String("string", "htt", "a string")
	i := flag.Int("int", 12, "a int")
	b := flag.Bool("bool", false, "a int")

	// 定义一个带有默认值的标志参数
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//	解析命令行参数
	flag.Parse()

	// 打印解析后的命令行参数
	fmt.Println(*s)
	fmt.Println(*i)
	fmt.Println(*b)
	fmt.Println(svar)
	//	不加参数名字的值
	fmt.Println(flag.Args())

}
