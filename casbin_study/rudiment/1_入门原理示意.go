package rudiment

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

var pathRoot = "rudiment"

func AuthMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetHeader("X-User") // 从请求头中获取用户信息
		obj := c.Request.URL.Path    // 获取请求路径
		act := c.Request.Method      // 获取请求方法

		if ok, err := e.Enforce(sub, obj, act); !ok {
			fmt.Println(err)
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
		}
	}
}

func Rudiment() {
	// 从文件中加载模型和策略
	/*m, err := model.NewModelFromFile(filepath.Join(pathRoot + "/" + "casbin_model.conf"))
	if err != nil {
		fmt.Println(err)
	}
	e, err := casbin.NewEnforcer(m, filepath.Join(pathRoot+"/"+"casbin_policy.csv"))*/

	e, err := casbin.NewEnforcer(filepath.Join(pathRoot+"/"+"casbin_model.conf"), filepath.Join(pathRoot+"/"+"casbin_policy.csv"))
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	//
	router.GET("/api/data", AuthMiddleware(e), func(c *gin.Context) {
		// 检查用户是否有访问资源1的权限
		c.JSON(http.StatusOK, gin.H{"message": "here is data"})
	})
	router.POST("/api/data", AuthMiddleware(e), func(c *gin.Context) {
		// 检查用户是否有访问资源1的权限
		c.JSON(http.StatusOK, gin.H{"message": "data create sussessful"})
	})

	// 启动服务器
	router.Run(":8080")
}
