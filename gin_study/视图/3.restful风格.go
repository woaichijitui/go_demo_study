package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// 请求文章列表
func _getList(c *gin.Context) {
	articleList := map[string]string{
		"1": "go语言基础 ",
		"2": "go语言圣经 ",
		"3": "go语言中间件 ",
	}
	c.JSON(200, Response{
		0, articleList, "成功",
	})
}

// 请求文章详情
func _getListDetil(c *gin.Context) {
	id := c.Param("id")
	//这里可以跟着一个查询操作
	article := map[string]string{
		string(id): "go语言基础 ",
	}
	c.JSON(200, Response{0, article, "成功"})
}

// 封装一个解析json到结构体函数
func _unmarshalJson(c *gin.Context, a any) error {
	body, err := c.GetRawData()
	if err != nil {
		return err
	}
	contentType := c.ContentType()
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, &a)
		if err != nil {
			return err
		}

	}
	return nil
}

// 添加一篇文章
func _create(c *gin.Context) {
	//	数据库操作不展示
	var articles map[string]string
	//	接收数据 ,并解析
	err := _unmarshalJson(c, &articles)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, Response{0, articles, "添加成功"})
}

// 修改一篇文章
func _upDate(c *gin.Context) {
	id := c.Param("id")
	id = "2"
	//这里可以跟着一个查询操作
	article := map[string]string{
		id: "go语言中间件  ",
	}
	c.JSON(200, Response{0, article, "修改成功"})
}

// 删除一篇文章
func _delete(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	c.JSON(200, Response{0, map[string]string{}, "修改成功"})
}

// 四大请求方式
func main() {
	router := gin.Default()

	router.GET("/articles", _getList)
	router.GET("/articles/:id", _getListDetil)
	router.POST("/articles", _create)
	router.PUT("/articles/:id", _upDate)
	router.DELETE("/articles/:id", _delete)

	router.Run(":8080")
}
