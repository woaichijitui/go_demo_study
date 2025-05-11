package main

import (
	_case "basis/generic-T/case"
	"context"
	"os"
	"os/signal"
)

func main() {
	_case.SimpleCase()
	_case.CusNumCase()
	_case.ComparableCase()

	_case.TTypeCase()
	_case.TTypeCase()
	// 为了让主进程不那么快结束
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
