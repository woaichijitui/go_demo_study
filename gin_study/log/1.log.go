package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main() {

	//输出日志到logfile文件
	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	//自定义路由输出的日志格式
	//	若该函数里面没有内容，则执行默认的内容
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("[HTT]  %s %s %s %d \n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "htt"})
	})
	//	自定义log配置
	//	1
	gin.Default()
	//	2

	router.Run()
}
