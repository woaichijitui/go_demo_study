package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type User1 struct {
	Name string `json:"name"  binding:"required,min=1" msg:"必须输入，且必须大于或等于1个字符"` //注意验证器 min之类的不能留空
	Age  int    `json:"age"  binding:"gt=18"  msg:"年龄必须大于或等于18"`
}

func myValidator(err error, obj any) string {
	//断言
	errors := err.(validator.ValidationErrors)
	for _, e := range errors {
		userRef := reflect.TypeOf(obj)
		//	e.Field() 返回字段名字 json标签的名字优先
		if tags, exist := userRef.Elem().FieldByName(e.Field()); exist {
			msg := tags.Tag.Get("msg")
			return msg
		}
	}
	return ""
}

// gin支持自定错误信息，tag：msg
func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		var user User1

		err := c.ShouldBindJSON(&user)

		if err != nil {
			msg := myValidator(err, &user)
			c.JSON(200, gin.H{"错误": msg})
			return
		}

		c.JSON(200, &user)

	})

	router.Run(":8080")
}
