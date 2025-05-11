package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	//设置样式
	//	json
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//	text
	//logrus.SetFormatter(&logrus.TextFormatter{})

	//设置字段 并返回*entry
	log := logrus.WithField("name", "htt").WithField("ip", "127.0.0.1")

	//	从上往下  输出其上面和自身等级日志
	//	FatalLevel
	//	ErrorLevel
	//	WarnLevel
	//	InfoLevel
	//	DebugLevel
	//	TraceLevel
	log.Errorln("你好")
	log.Warnln("你好")
	log.Infoln("你好")
	log.Debugln("你好")
	log.Traceln("你好")
}
