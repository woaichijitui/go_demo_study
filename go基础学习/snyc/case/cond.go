package _case

import (
	"fmt"
	"sync"
	"time"
)

// 应用于一发多收的场景，即一组协程需要等待某一个协程完成一些前置准备的情况。
func CondCase() {
	list := make([]int, 0)
	//不能这样 是个空的结构体

	//cond := &sync.Cond{}
	cond := sync.NewCond(&sync.Mutex{})
	go readCond(&list, cond)
	go readCond(&list, cond)
	go readCond(&list, cond)
	time.Sleep(2 * time.Second)
	initList(&list, cond)
}

// 主叫方
// 主叫方可以持有锁，但允许不持有
func initList(list *[]int, cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()

	for i := 0; i < 10; i++ {
		*list = append(*list, i)
	}
	//唤醒所有的有条件等待的协程
	cond.Broadcast()
}

func readCond(list *[]int, cond *sync.Cond) {
	//被叫方必须持有锁
	cond.L.Lock()
	defer cond.L.Unlock()
	for len(*list) == 0 {
		fmt.Println("readCond wait ")
		//等待被叫醒
		cond.Wait()
	}
	fmt.Println(*list)

}

// 队列
type queue struct {
	list []int
	cond *sync.Cond
}

func NewQueue() *queue {
	q := &queue{
		list: []int{},
		cond: sync.NewCond(&sync.Mutex{}),
	}
	return q
}
func (q *queue) put(item int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.list = append(q.list, item)
	//每次写入数据后就唤醒一次协程
	q.cond.Signal()
}

// 自定义可以从list获取n个值
func (q *queue) getMany(n int) []int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.list) < n {
		//当list长度小于所需的长度就会等待
		q.cond.Wait()
	}
	list := q.list[:n]
	q.list = q.list[n:]
	return list
}

// 队列读取
func CondQueueCase() {
	q := NewQueue()
	wg := sync.WaitGroup{}

	//启动十次协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			list := q.getMany(n + 1)
			fmt.Printf("%d:%v\n", n+1, list)
		}(i)
	}
	for i := 0; i < 100; i++ {
		q.put(i)

	}
	wg.Wait()
}
