package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main() {

	//输出日志到logfile文件
	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// 不输出debug 内容的模式
	gin.SetMode(gin.ReleaseMode)

	//自定义路由输出的日志格式
	//	若该函数里面没有内容，则执行默认的内容
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("[HTT]  %s %s %s %d \n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	// 	自定义日志输出格式
	//	要新建一个没有配置过的路由
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"[htt] %s   |%d|   	%d  	%s		%s \n",
			params.TimeStamp.Format("2006-01-02 15:04:05"),
			params.StatusCode,
			params.Latency,
			params.Method,
			params.Path,
		)
	}))

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "htt"})
	})

	router.Run()

}
