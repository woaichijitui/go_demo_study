package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

/*
	自定义hook
	实现hook 接口：
*/

type myHook struct {
	filePath string
}

// 定义在何种等级时 调用hook
func (h myHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
	}
}

func (h myHook) Fire(entry *logrus.Entry) error {

	file, _ := os.OpenFile(h.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	str, _ := entry.String()

	file.Write([]byte(str))
	return nil

}

func main() {
	// 将hook 对象添加到logrus上
	logrus.AddHook(&myHook{"logfile/logrus.log"})

	logrus.Errorln("hook test")
	logrus.Infoln("hook test")

}
