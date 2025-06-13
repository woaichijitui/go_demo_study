package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})

	file, err := os.OpenFile("./logfile/logrus.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	//同时输出到控制台和文件
	writers := []io.Writer{
		file,
		os.Stdout,
	}

	fileAndStdout := io.MultiWriter(writers...)

	logrus.SetOutput(fileAndStdout)

	logrus.Errorln("你好")
	logrus.Warnln("你好")
	logrus.Infoln("你好")
	logrus.Debugln("你好")
	logrus.Traceln("你好")

}
