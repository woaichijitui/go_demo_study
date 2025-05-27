package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

var n int32 = 0

func incr() {
	time.Sleep(10 * time.Millisecond)
	//n++
	atomic.AddInt32(&n, 1)
}

func main() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	go func() {
		for {
			<-t.C
			fmt.Println(runtime.NumGoroutine())
		}
	}()

	g := NewGorotine_ch(100)

	for i := 0; i < 10001; i++ {
		g.go_chan(incr)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(atomic.LoadInt32(&n))
	time.Sleep(2 * time.Second)
	fmt.Println(atomic.LoadInt32(&n))
}

type Gorotine_ch struct {
	Limi int
	ch   chan struct{}
}

func NewGorotine_ch(n int) *Gorotine_ch {
	return &Gorotine_ch{
		Limi: n,
		ch:   make(chan struct{}, n),
	}
}

func (g *Gorotine_ch) go_chan(run func()) {
	g.ch <- struct{}{}
	go func() {
		run()
		<-g.ch
	}()
}
