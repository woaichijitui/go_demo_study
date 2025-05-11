package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

var name *string

func handle(c *gin.Context) {
	bytes, _ := json.Marshal(c.Request.Header)
	fmt.Println(string(bytes))
	c.String(200, fmt.Sprintf("%s:%s name:%s", c.Request.URL.Path, c.ClientIP(), *name))
	return
}

func main() {
	//主机地址和名字
	addr := flag.String("addr", ":8080", "address")
	name = flag.String("name", "default", "name")
	flag.Parse()

	router := gin.Default()
	router.GET("/", handle)
	fmt.Println("服务器运行中...")
	fmt.Printf("服务器运行在：%s \n", *addr)
	router.Run(*addr)
}
