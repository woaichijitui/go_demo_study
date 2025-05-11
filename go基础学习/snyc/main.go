package main

import (
	"os"
	"os/signal"
	_case "snyc/case"
)

func main() {
	//_case.WaitGroupCase1()
	//_case.A1()
	_case.WaitGroupCase2()
	//_case.CondCase()
	//_case.CondQueueCase()
	//_case.MutexCase()

	//_case.PoolCase1()
	//_case.OnceCase()

	/*m := make(map[int]string, 1)
	for key, val := range m {
		fmt.Printf("%d,%v", key, val)
	}*/
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
