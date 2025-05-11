package main

import (
	"flag"
	"fmt"
	"time"
)

var j = flag.Int("j", 0, "")

func main() {
	flag.Parse()
	var i int = 0
	for {
		fmt.Println("go demo ", i, *j)
		i++
		time.Sleep(time.Second)
	}

}
