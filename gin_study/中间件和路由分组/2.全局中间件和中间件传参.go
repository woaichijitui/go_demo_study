package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func _m3(c *gin.Context) {
	fmt.Println("m3 in")
	//	中间件传递参数
	c.Set("name", "tt")
	c.Next()
	fmt.Println("m3 out")

}

func _index(c *gin.Context) {
	fmt.Println("index in")

	c.Next()
	c.JSON(200, gin.H{"data": c.Value("name")})
	fmt.Println("index out")
}
func main() {
	router := gin.Default()

	api := router.Use(_m3)
	{
		api.GET("/use", _index)
	}

	router.Run(":8080")
}
