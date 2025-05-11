package _case

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	ON  = 1
	OFF = 0
)

func PoolCase1() {
	target := "192.144.144.12"
	wg := sync.WaitGroup{}
	//	创建一个连接池
	pool := getPool(target)

	// 	通过10个协程 不停的利用连接池创建对象、释放对象
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//	每个协程循环五次
			//	获取对象并释放对象
			for j := 0; j < 5; j++ {
				//get方法 当pool里有可用的连接时 就创建
				c := pool.Get()
				fmt.Println(c.getId())
				pool.Put(c)
			}

		}()
	}
	wg.Wait()
}

// 创建一个放入pool中的连接对象
type conn struct {
	id     int64
	target string
	status int
}

// 获取conn对象
func getConn(t string) *conn {
	return &conn{
		id:     rand.Int63(),
		target: t,
		status: ON,
	}
}

func (c *conn) getId() int64 {
	return c.id
}

// pool连接池对象
type connPool struct {
	sync.Pool
}

// 获取连接池对象
func getPool(target string) *connPool {
	return &connPool{
		Pool: sync.Pool{
			//用于声明conn对象的创建方式
			New: func() any {
				fmt.Println(" Creating new conn")
				return getConn(target)
			},
		},
	}
}

// 这里进一步处理下 Get操作
func (cp *connPool) Get() *conn {
	c := cp.Pool.Get().(*conn)

	if c.status == OFF {
		c = cp.Pool.New().(*conn)

	}
	return c
}

// 这里进一步处理下 Get操作
func (cp *connPool) Put(c *conn) {
	if c.status == OFF {
		return
	}
	cp.Pool.Put(c)
}
