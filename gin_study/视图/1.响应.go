package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context) {
	c.String(200, "响应string")
}

// 响应json
func _json(c *gin.Context) {
	//	结构体转json
	//	`json:"-"` 表示忽略该字段 不json序列化
	type User struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		PassWord string `json:"-"`
	}
	user := User{
		UserName: "htt",
		Age:      23,
		PassWord: "123456",
	}

	c.JSON(200, user)
	//	map转
	mp := map[string]string{"UserName": "htt", "Age": "23", "PassWord": "123456"}
	c.JSON(200, mp)

	//	gin.H 转
	c.JSON(200, gin.H{"UserName": "htt", "Age": "23", "PassWord": "123456"})
}

// 响应html
func _html(c *gin.Context) {
	//	第二个形参是模板下的文件名
	//	第三个形参可以传递任意东西
	c.HTML(200, "index.html", gin.H{"username": 11})
}

// 重定向 redirect 301和302的区别
func _redirec(c *gin.Context) {
	c.Redirect(302, "https://www.baidu.com")
}

// query请求
// 请求动态参数
// 表单form
// 原始参数
// 封装一个解析json到结构体函数
func main() {
	router := gin.Default()

	//加载全局模板模板
	//router.LoadHTMLGlob("./templates/*")
	//	加载一个文件夹的模板
	router.LoadHTMLFiles("./templates/index.html")

	//	加载静态文件
	//	relativePath 即url
	router.StaticFS("/static/test.txt", http.Dir("./static/test.txt"))
	router.StaticFS("/static/config.yaml", http.Dir("./static/config.yaml"))

	//响应字符串
	router.GET("/string", _string)
	//响应json
	router.GET("/json", _json)
	// 响应html
	router.GET("/html", _html)
	router.GET("/redirect", _redirec)
	//启动监听
	router.Run("127.0.0.1:8080")
}
