package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// simple case
func main() {

	//	定义一个简单的路由
	router := gin.Default()
	// 请求方法
	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello tt")
	})

	//	开启监听
	//	等同于0.0.0.0地址
	/*err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Println(err)
	}*/

	//	或者启用这种监听
	http.ListenAndServe("127.0.0.1:8080", router)
}
