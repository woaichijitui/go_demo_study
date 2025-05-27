package main

import (
	"fmt"
	runtime "runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
}
