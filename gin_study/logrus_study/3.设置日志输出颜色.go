package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	// 开启颜色设置
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	//
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})

	logrus.Errorln("你好")
	logrus.Warnln("你好")
	logrus.Infoln("你好")
	logrus.Debugln("你好")
	logrus.Traceln("你好")

	// "/033[30m 黑色 /033[0m " \033[ 和 /033[0m 开头和结尾 ；30m 代表颜色
	/*// 前景色
	fmt.Println("\033[30m 黑色 \033[0m")
	fmt.Println("\033[31m 红色 \033[0m")
	fmt.Println("\033[32m 绿色 \033[0m")
	fmt.Println("\033[33m 黄色 \033[0m")
	fmt.Println("\033[34m 蓝色 \033[0m")
	fmt.Println("\033[35m 紫色 \033[0m")
	fmt.Println("\033[36m 青色 \033[0m")
	fmt.Println("\033[37m 灰色 \033[0m")
	// 背景色
	fmt.Println("\033[40m 黑色 \033[0m")
	fmt.Println("\033[41m 红色 \033[0m")
	fmt.Println("\033[42m 绿色 \033[0m")
	fmt.Println("\033[43m 黄色 \033[0m")
	fmt.Println("\033[44m 蓝色 \033[0m")
	fmt.Println("\033[45m 紫色 \033[0m")
	fmt.Println("\033[46m 青色 \033[0m")
	fmt.Println("\033[47m 灰色 \033[0m")*/
}
