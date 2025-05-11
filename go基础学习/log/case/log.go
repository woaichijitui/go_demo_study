package _case

import (
	"log"
	"os"
)

// 日志的初始化
func init() {

	//设置前缀
	log.SetPrefix("htt ")

	//设置输出的东西
	/*log.Ldate：日志中包含日期（形如 2009/01/23）。
	log.Ltime：日志中包含时间（形如 01:23:23）。
	log.Lmicroseconds：日志中包含微秒级时间戳（形如 01:23:23.123123）。注意，这个标志会覆盖 log.Ltime，使其显示微秒级别的时间。
	log.Llongfile：日志中包含完整的文件路径和行号（形如 /a/b/c/d.go:23）。
	log.LUTC：使用 UTC 时间而不是本地时间。
	log.Lmsgprefix：在日志消息前添加前缀。
	log.LstdFlags：这是默认的标志组合，相当于 log.Ldate | log.Ltime。
	然而，log.LstdFlags 已经包含了 log.Ldate 和 log.Ltime，所以它在这里是冗余的。为了避免冗余，可以去掉 log.LstdFlags。最终设置可以这样：*/
	//	按位或运算符（|）？
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.LUTC | log.Lmsgprefix | log.LstdFlags)

	//打开的文件地址： 项目根地址（mod所在地址）+ dir/。。。/filename
	file, err := os.OpenFile("logfile/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//	设置输出
	log.SetOutput(file)
}

func LogCase() {

	//	有三种类型
	log.Printf("日志打印")

	//	在panic之前和退出之前都会把日志输出了
	log.Fatal("程序退出")
	log.Panic("err")

}
