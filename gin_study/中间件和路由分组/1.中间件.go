package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	fmt.Println("index in")

	c.Next()
	fmt.Println("index out")
}
func _m1(c *gin.Context) {
	fmt.Println("m1 in")

	c.Next()
	c.Abort()
	fmt.Println("m1 out")
}
func _m2(c *gin.Context) {
	fmt.Println("m2 in")

	c.Next()
	fmt.Println("m2 out")
}

func main() {

	router := gin.Default()

	router.GET("/", _m1, index, _m2)

	router.Run(":8080")

}
