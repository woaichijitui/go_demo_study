package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 请求query
//
//	在url后面接上?username =
func _query(c *gin.Context) {
	fmt.Println(c.Query("username"))
	fmt.Println(c.QueryArray("username"))
	fmt.Println(c.GetQuery("username"))
	fmt.Println(c.DefaultQuery("username", "htt"))

}

// 请求动态参数
func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("user_age"))
}

// 表单form
func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	//fmt.Println(c.PostFormMap("name")) 不会用
	fmt.Println(c.ContentType())
}

// 原始参数
func _raw(c *gin.Context) {
	body, _ := c.GetRawData()
	fmt.Println(string(body))
}

// 封装一个解析json到结构体函数
func unmarshalJson(c *gin.Context, a any) error {
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

// json
func _bindJson(c *gin.Context) {
	type User struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
	}
	var user User
	err := unmarshalJson(c, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func main() {
	router := gin.Default()

	router.GET("/query", _query)
	// 请求动态参数
	router.GET("/param/:user_id/", _param)
	router.GET("/param/:user_id/:user_age", _param)
	// 表单form
	router.POST("/form", _form)
	// 原始参数
	router.POST("/raw", _raw)
	// json
	router.POST("/json", _bindJson)

	//启动监听
	router.Run("127.0.0.1:8080")
}
