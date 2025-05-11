package f1

import (
	"fmt"
	_ "go_structure/func2"
)

func init() {
	fmt.Println("调用了 f1 包的 init ")
}
func F1() {
	fmt.Println("调用了 f1 的 F1 函数  ")
}
