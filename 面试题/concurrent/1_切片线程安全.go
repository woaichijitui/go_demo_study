package main

import (
	"fmt"
	"sync"
	"time"
)

func SliceTest1() {
	// 创建一个空切片
	var mySlice []int
	wait := sync.WaitGroup{}
	now := time.Now().UnixMilli()
	for i := 0; i < 10000; i++ {
		wait.Add(1)
		go func(i int) {
			mySlice = append(mySlice, i)

			defer wait.Done()
		}(i)
	}
	wait.Wait()
	for v, i := range mySlice {
		fmt.Printf("mySlice[%d] = %d\n", v, i)
	}
	over := time.Now().UnixMilli()
	t := over - now
	fmt.Printf("耗时：%d msec\n ", t)

}

func SliceTest2() {
	var mySlice []int
	var mu sync.Mutex
	wait := sync.WaitGroup{}
	now := time.Now().UnixMilli()
	for i := 0; i < 10000; i++ {
		wait.Add(1)
		go func(i int) {
			mu.Lock()
			mySlice = append(mySlice, i)
			mu.Unlock() // 修正方法名拼写错误
			wait.Done()
		}(i)
	}
	wait.Wait()
	for v, i := range mySlice {
		fmt.Printf("mySlice[%d] = %d\n", v, i)
	}
	over := time.Now().UnixMilli()
	t := over - now
	fmt.Printf("耗时：%d msec \n ", t)
}
