package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

/*
本质上是生产者消费者模型
可以有效控制goroutine数量，防止暴涨
需求：
。计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
。随机生成数字进行计算
控制台输出结果如下：
*/

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func main() {
	//	创建通道
	jobChan := make(chan *Job)
	resultChan := make(chan *Result)

	// 创建线程池
	CreatePool(60, jobChan, resultChan)
	//	打印结果
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)
	//	创建随机数
	var id int
	for {
		id++
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job

	}

}

// 创建工作池
// 参数1：开几个协程
func CreatePool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				// 	计算个位数之和
				randNum := job.RandNum
				var sum int
				var temp int
				for randNum != 0 {
					temp = randNum % 10
					sum += temp
					randNum /= 10
				}
				var result = &Result{
					job: job,
					sum: sum,
				}

				resultChan <- result
			}
		}(jobChan, resultChan)

	}
}

func printGoroutineInfo() (buf []byte) {
	// 获取当前goroutine的堆栈信息
	buf = make([]byte, 1024)
	n := runtime.Stack(buf, true)
	return buf[:n]

}
