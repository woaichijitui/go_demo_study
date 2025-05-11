package _case

import (
	"fmt"
	"sync"
)

// 创建一个带有once、map的结构体
type onceMap struct {
	sync.Once
	Data map[string]int
}

func (o *onceMap) LoadData() {
	//  once.do 方法里的操作在并发场景下，只能执行一次
	o.Once.Do(func() {
		o.Data = make(map[string]int, 0)
		list := []string{"A", "B", "C", "D"}

		for _, item := range list {
			_, ok := o.Data[item]
			if !ok {
				o.Data[item] = 0
			}
			o.Data[item] += 1
		}

	})
	/*list := []string{"A", "B", "C", "D"}

	for _, item := range list {
		_, ok := o.Data[item]
		if !ok {
			o.Data[item] = 0
		}
		o.Data[item] += 1
	}*/
}

// 用多个协程测试
func OnceCase() {
	om := &onceMap{
		Data: make(map[string]int),
	}

	wg := sync.WaitGroup{}

	//	用十个协程触发onceMap.once.do 方法
	//	正常情况有几个循环，list上每个字母都等于几
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			om.LoadData()
		}()
	}

	for key, val := range om.Data {
		fmt.Printf("%v : %d ", key, val)
	}

	wg.Wait()
}
