package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	//	从上往下  输出其上面和自身等级日志
	//	FatalLevel
	//	ErrorLevel
	//	WarnLevel
	//	InfoLevel
	//	DebugLevel
	//	TraceLevel
	logrus.Errorln("你好")
	logrus.Warnln("你好")
	logrus.Infoln("你好")
	logrus.Debugln("你好")
	logrus.Traceln("你好")
}
