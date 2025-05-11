package _case

import (
	"fmt"
	"sync"
)

func MutexCase() {
	//singleGoutine()
	//multipleGoutine()
	//multipleSafeGoutine()
	multipleSafeGoutineRWmutex()
}

// 单协程操作
func singleGoutine() {
	list := []string{"A", "B", "C", "D"}

	mp := make(map[string]int, 0)
	for i := 0; i < 20; i++ {
		for _, str := range list {
			_, ok := mp[str]
			if !ok {
				//添加或者修改键值对
				mp[str] = 0
			}
			mp[str] += 1
		}
	}
	fmt.Println(mp)
}

// 多协程操作，非线程安全
func multipleGoutine() {

	wg := sync.WaitGroup{}

	list := []string{"A", "B", "C", "D"}

	mp := make(map[string]int, 0)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, str := range list {
				_, ok := mp[str]
				if !ok {
					//添加或者修改键值对
					mp[str] = 0
				}
				mp[str] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp)
}

// 互斥锁线程安全
func multipleSafeGoutine() {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	list := []string{"A", "B", "C", "D"}

	mp := make(map[string]int, 0)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			//可以写进来
			mutex.Lock()

			defer wg.Done()
			defer mutex.Unlock()
			for _, str := range list {
				_, ok := mp[str]
				if !ok {
					//添加或者修改键值对
					mp[str] = 0
				}
				mp[str] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp)
}

type cache struct {
	mp map[string]string
	sync.RWMutex
}

func newCache() cache {
	return cache{
		mp: make(map[string]string),
		//这里相当
		RWMutex: sync.RWMutex{},
	}
}

// 输入key 取出value
// 读取的时候上读锁
func (c *cache) get(str string) string {
	c.RLock()
	defer c.RUnlock()
	val, ok := c.mp[str]
	if ok {
		return val
	}
	//
	val = c.mp[str]
	return ""
}
func (c *cache) set(key string, val string) {
	c.Lock()
	c.Unlock()
	c.mp[key] = val
}

// 读写锁线程安全
func multipleSafeGoutineRWmutex() {
	c := newCache()
	wg := sync.WaitGroup{}
	// 一个写的协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.mp["name"] = "htt"
	}()
	//用十个协程来读
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			val := c.mp["name"]
			fmt.Printf("协程 %d 读取：%v\n ", n, val)
		}(i)
	}

	wg.Wait()
}
