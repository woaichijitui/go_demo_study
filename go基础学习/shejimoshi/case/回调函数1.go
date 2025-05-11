package huidiao

import (
	"fmt"
	"sort"
)

// 使用回调函数定义排序逻辑
func SortWithCustomLogic(values []int, compare func(a, b int) bool) {
	sort.Slice(values, func(i, j int) bool {
		return compare(values[i], values[j])
	})
}

func Test1() {
	ints := []int{4, 3, 11, 6}

	SortWithCustomLogic(ints, func(a int, b int) bool {
		return a < b
	})

	fmt.Println(ints)
}
