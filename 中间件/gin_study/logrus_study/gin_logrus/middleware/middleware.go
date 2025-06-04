package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		//	让日志在所有东西都执行完毕后 再执行
		c.Next()

		since := time.Now()
		//	执行时间
		latencyTime := since.Sub(start)

		hookInit("logrus_study/gin_logrus/gin_logrus.log", logger)
		//LogOutFile("logrus_study/gin_logrus/gin_logrus.log", logger)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		logger.Infof("[GIN] %s |%d| %d | %s | %s | %s",
			start.Format("2006-01-02 15:04:06"),
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)

	}
}

func LogOutFile(path string, logger *logrus.Logger) {
	logger.SetFormatter(&logrus.TextFormatter{})

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	logger.SetOutput(file)
}

/*
logrus hook

*/

func hookInit(path string, logger *logrus.Logger) {
	logger.AddHook(&myHook{path})
}

type myHook struct {
	filePath string
}

// 定义在何种等级时 调用hook
func (h myHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h myHook) Fire(entry *logrus.Entry) error {

	file, _ := os.OpenFile(h.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	str, _ := entry.String()

	file.Write([]byte(str))
	return nil

}
