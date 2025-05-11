package main

import (
	_case "channel_select/case"
	"os"
	"os/signal"
)

func main() {

	//_case.C1()
	//_case.ConcurrentSync()
	_case.NoticeAndMutiplexing()

	//	主协程要等待其他协程结束才能结束
	//	利用channel的阻塞机制
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
