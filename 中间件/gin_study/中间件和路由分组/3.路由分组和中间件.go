package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func mindleWare(c *gin.Context) {
	fmt.Println("mindle ware in")

	c.Next()
	fmt.Println("mindle ware out")

}

func _login(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "登录成功"})

}

func _user(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "登录成功"})

}

func Middleware(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware in")
		//需要将defer放到return之前
		defer fmt.Println("middleware out")

		token := c.Query("token")
		if token == "1234" {
			//若是next之后return 就不再执行c.abort之后
			c.Next()
			return
		}
		c.Abort()
		c.JSON(200, gin.H{"msg": msg})
		// defer ？

	}
}

func routerGroupInit(router *gin.RouterGroup) {

	userManger := router.Group("/manger").Use(Middleware("用户登录失败"))

	{
		userManger.GET("/login2", _login)
	}
}

func main() {
	router := gin.Default()

	//路由分组
	users := router.Group("/users").Use(Middleware("路由分组中间件 token出现错误"))
	{
		users.POST("/login", _login)
	}
	api := router.Group("/api")
	//路由分开写
	routerGroupInit(api)
	router.Run(":8080")
}
