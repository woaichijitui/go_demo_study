package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Name string `json:"name"  uri:"name" form:"name"`
	Age  int    `json:"age" uri:"age" form:"age"`
	Msg  string `json:"msg" uri:"msg" form:"msg"`
}

// shouldbind 绑定query url(动态参数) json form xml yaml...
func main() {
	router := gin.Default()

	//	bind query
	//	tag : form
	//	若没有绑定 下面代码不会报错，会响应一个初始值的user
	router.POST("/query", func(c *gin.Context) {
		var user User
		err := c.ShouldBindQuery(&user)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}

		c.JSON(200, &user)
	})

	//	bind uri
	//	既动态参数
	router.POST("/uri/:name/:age/:msg", func(c *gin.Context) {
		var user User
		err := c.ShouldBindUri(&user)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}

		c.JSON(200, &user)
	})

	//	bind json
	router.POST("/json", func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}

		c.JSON(200, &user)
	})

	//	bind json
	router.POST("/form", func(c *gin.Context) {
		var user User
		err := c.ShouldBind(&user)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}

		c.JSON(200, &user)
	})

	router.Run(":8080")
}
