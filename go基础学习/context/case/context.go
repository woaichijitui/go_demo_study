package _case

import (
	"context"
	"fmt"
	"time"
)

func ContextCase() {
	ctx := context.Background()
	//创建一个携带键值对信息的上下文，用于传递额外的数据。
	ctx = context.WithValue(ctx, "desc", "ContextCase")
	ctx, cancle := context.WithCancel(ctx)
	//ctx, cancle = context.WithDeadline(ctx, time.Now())

	//创建一个带有超时时间的上下文，当超过指定的时间后，衍生的上下文会被取消。
	ctx, cancle = context.WithTimeout(ctx, 2*time.Second)

	data := [][]int{[]int{1, 2}, []int{3, 4}}
	ch := make(chan []int)
	go calculate(ctx, ch)

	for i := 0; i < len(data); i++ {
		ch <- data[i]

	}
	//手动通知关闭
	defer cancle()
	time.Sleep(10 * time.Second)
}

func calculate(ctx context.Context, ch <-chan []int) {

	for {
		i := 0
		//多路复用
		select {
		case item := <-ch:

			i++
			str := "calculate" + string(i)
			ctx = context.WithValue(ctx, "desc1", str)
			fmt.Println(item)
			ch1 := make(chan []int)
			ch2 := make(chan []int)
			go sumContext(ctx, ch1)
			go multiContext(ctx, ch2)
			ch1 <- item
			ch2 <- item
		//	ctx超时或者主动被关闭 则会在这个通道里发送东西
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("calculate 协程退出，context desc：%s ,错误消息：%s\n", desc, ctx.Err())
			return
		}
	}
}

func sumContext(ctx context.Context, ch <-chan []int) {
	for {
		//多路复用
		select {
		case item := <-ch:
			a, b := item[0], item[1]

			res := sum(a, b)

			fmt.Printf("%d + %d = %d \n", a, b, res)
		//	ctx超时或者主动被关闭 则会在这个通道里发送东西
		case <-ctx.Done():
			desc := ctx.Value("desc1").(string)
			fmt.Printf("sumContext 协程退出，context desc1：%s ,错误消息：%s\n", desc, ctx.Err())
			return
		}
	}
}

func multiContext(ctx context.Context, ch <-chan []int) {
	for {
		//多路复用
		select {
		case item := <-ch:
			a, b := item[0], item[1]

			res := multi(a, b)

			fmt.Printf("%d * %d = %d \n", a, b, res)
		//	ctx超时或者主动被关闭 则会在这个通道里发送东西
		case <-ctx.Done():
			desc := ctx.Value("desc1").(string)
			fmt.Printf("multiContext 协程退出，context desc1：%s ,错误消息：%s\n", desc, ctx.Err())
			return
		}
	}
}
func sum(a, b int) int {
	return a + b

}
func multi(a, b int) int {
	return a * b

}
