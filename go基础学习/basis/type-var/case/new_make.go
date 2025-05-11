package _case

import "fmt"

func New_case() {
	//var slice1 []int
	slice1 := new([]int) //new 传递的是地址

	slice1 = &[]int{1, 2, 3, 4}

	slice2 := make([]int, 4, 10)
	slice2[0] = 11
	slice2[1] = 11
	slice2[2] = 11
	slice2[3] = 11

	slice2 = append(*slice1, 5, 6, 7, 8)

	*slice1 = slice2[:2]
	fmt.Println(slice1, slice2)
	fmt.Println(len(*slice1), cap(*slice1))
}

func Map_case() {
	//var mp map[string]string
	mp := make(map[string]int, 10)
	mp["A"] = 1
	mp["B"] = 2
	mp["C"] = 3
	fmt.Println(mp)

	for k, v := range mp {
		fmt.Println(k, "=", v)
	}
}
