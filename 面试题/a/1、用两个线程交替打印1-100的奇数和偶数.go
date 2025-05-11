package a

import (
	"fmt"
	"sync"
)

func A(n int) {

	j1 := make(chan bool)
	o1 := make(chan bool)

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(2)

	go func(n int) {
		defer waitGroup.Done()
		for i := 1; i <= n; i += 2 {
			<-o1
			fmt.Printf("奇数线程：%v\n", i)
			if i < n {
				j1 <- true
			}
		}
	}(n)

	go func(n int) {
		defer waitGroup.Done()
		for i := 2; i <= n; i += 2 {
			<-j1
			fmt.Printf("偶数线程：%v\n", i)
			if i < n {
				o1 <- true
			}
		}
	}(n)

	o1 <- true

	waitGroup.Wait()

}

func B(n int) {

	j1 := make(chan bool, 1)
	o1 := make(chan bool, 1)

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(2)

	go func(n int) {
		defer waitGroup.Done()
		for i := 1; i <= n; i += 2 {
			<-o1
			fmt.Printf("奇数线程：%v\n", i)
			j1 <- true
		}
	}(n)

	go func(n int) {
		defer waitGroup.Done()
		for i := 2; i <= n; i += 2 {
			<-j1
			fmt.Printf("偶数线程：%v\n", i)
			o1 <- true
		}
	}(n)

	o1 <- true

	waitGroup.Wait()

}
func C(n int) {

	c := make(chan bool)

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(2)

	go func(n int) {
		defer waitGroup.Done()
		for i := 0; i <= n; i++ {
			<-c

			if i%2 == 1 {
				fmt.Printf("奇数线程：%v\n", i)
			}
		}
	}(n)

	go func(n int) {
		defer waitGroup.Done()
		for i := 0; i <= n; i++ {
			c <- true

			if i%2 == 0 {
				fmt.Printf("奇数线程：%v\n", i)
			}
		}
	}(n)

	waitGroup.Wait()

}
