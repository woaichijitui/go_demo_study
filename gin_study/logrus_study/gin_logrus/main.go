package main

import (
	"gin_study/logrus_study/gin_logrus/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 创建一个新的 logrus logger
	logger := logrus.New()

	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)

	// 设置日志格式（可选）
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// 将logrus 放到全局中间件处理
	router.Use(middleware.LoggerMiddleware(logger))
	router.POST("/logrus", func(context *gin.Context) {
		context.JSON(200, gin.H{"data": "收到"})
	})
	router.Run()
}
