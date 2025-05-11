package _case

import (
	"fmt"
	"sync"
	"time"
)

// 十亿次两数相乘
// 分别看单进程和两个协程的速度
func WaitGroupCase1() {
	var a, b = 1000, 10000
	now := time.Now()
	for i := 0; i < 10000000000; i++ {
		multi(a, b)
	}
	fmt.Println(time.Since(now))

	//	waitgroup是阻塞主进程的
	now1 := time.Now()

	//创建一个waitgroup对象
	wait := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		//计数加1，若计数不为0，则主线程阻塞
		wait.Add(1)
		go func() {
			//done 是计数减1
			defer wait.Done()
			for i := 0; i < 5000000000; i++ {

				multi(a, b)
			}
		}()
	}
	wait.Wait()

	since := time.Since(now1)
	fmt.Println(since)
}

// 两数相乘
func multi(a, b int) int {
	return a * b
}

// 协程间传递时需要以指针的方式或闭包的方式引用WaitGroup对象。否则将会造成死锁
//
//	以下是案例
func A1() {
	list := make([]int, 4)
	list = append(list, 1, 2, 3, 4)
	fmt.Println("刚开始的list:", list)
	//变量的值,若是引用类型就是地址,值类型就是数值
	fmt.Printf("变量的值地址:%p\n", list)
	//存放变量的地址
	fmt.Printf("存放变量地址:%p\n", &list)

	fmt.Println("____________________________")

	test1(list)
	fmt.Println("____________________________")

	fmt.Printf("改变后,变量的值地址:%p\n", list)
	fmt.Printf("改变后,存放变量地址:%p\n", &list)
	fmt.Println("改变后的list:", list)
}

func test1(list []int) []int {

	fmt.Printf("函数内变量的值地址:%p\n", list)
	fmt.Printf("函数内存放变量地址:%p\n", &list)
	fmt.Println("____________________________")
	//不会改变变量的值(地址)
	list[2] = 10

	//append 操作会改变变量的值(地址),
	list = append(list, 10)
	fmt.Printf("改变后,函数内变量的值地址:%p\n", list)
	fmt.Printf("改变后,函数内存放变量地址:%p\n", &list)
	fmt.Println("函数内append操作后的list:", list)
	return list
}

// 通过协程并行执行一组任务，且任务全部完成后才能进行下一步操作的情况
// 例如：汽车的生成，所有零件可以并行生产，只能等所有零件生成完成后，才能组装
func WaitGroupCase2() {
	now := time.Now()
	wg1 := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}
	ch := make(chan []int, 1000)
	wg2.Add(1)
	//将数据读取出来
	go func() {
		defer wg2.Done()

		i := 0
		for ints := range ch {
			fmt.Println(multi(ints[0], ints[1]))
			i++
		}
		time.Sleep(3 * time.Second)
		fmt.Printf("读取了%d 条数据\n", i)
	}()
	//	将数据填入channel
	for i := 0; i < 2; i++ {
		wg1.Add(1)
		//wg2.Add(1)
		go func() {
			defer wg1.Done()

			//defer wg2.Done()
			for j := 0; j < 500; j++ {
				ch <- []int{i, j}
			}
		}()
	}
	wg1.Wait()
	//为什么不关闭channel 或close 放在wg2.Wait() 会导致死锁
	//当发送者发送完了数据 chanel 在for range 循环里阻塞 一直等待着数据的发送 导致无法执行wg.done 导致主线程阻塞 又导致channel无法关闭 所以死锁
	close(ch)
	wg2.Wait()

	since := time.Since(now)
	fmt.Println("用时：", since)

}
