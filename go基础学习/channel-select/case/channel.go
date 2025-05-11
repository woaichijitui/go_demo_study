package _case

import (
	"fmt"
	"time"
)

// C1 协程间的通信
func C1() {
	//定义一个可读可写的通道
	ch := make(chan int, 0)
	//	只读或者只写的通道 可以接收一个可读可写的通道
	//	写入时,channel缓存满了,则阻塞
	//	读取时,channel无数据,则阻塞
	go F1(ch)
	go F2(ch)

}

// F1 接受一个只写通道
func F1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

// F2 接受一个只读通道
func F2(ch <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}

// ConcurrentSync 并发场景下,同步机制
//
//	指的是相同通道下(数据同步
func ConcurrentSync() {
	//	带缓冲的通道
	ch := make(chan int, 10)
	//	多个协程同时写入
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	//	只有单个协程读取数据
	go func() {
		for i := range ch {
			fmt.Printf("读取数据: %d\n", i)
		}
	}()

}

// 通知协程推出与多路复用
func NoticeAndMutiplexing() {
	ch := make(chan int)
	strCh := make(chan string)
	done := make(chan struct{}, 0)
	go readIntF1(ch)
	go readStrF2(strCh)
	go mutiplexing(ch, strCh, done)

	//	等待其他协程
	time.Sleep(5 * time.Second)
	close(done)
}

// 写入int
func readIntF1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}

}

// 写入string
func readStrF2(ch chan<- string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("写入数据str : %d\n", i)
	}
}

// 多路复用
func mutiplexing(ch <-chan int, strCh <-chan string, done <-chan struct{}) {

	i := 0
	//	select
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case str := <-strCh:
			fmt.Println(str)
		case <-done:
			fmt.Println("收到通知,协程退出")
			return
			/*default:
			fmt.Println("执行default ")*/
		}

		i++
		fmt.Printf("select 执行了 %d 次\n", i)
	}
}
