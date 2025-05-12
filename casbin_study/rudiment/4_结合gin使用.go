package rudiment

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var srv *CasbinService = CasbinInit()

func Test4() {
	srv := CasbinInit()
	roles, _ := srv.GetRolesPolicy()
	fmt.Println(roles)

	router := gin.Default()

	//测试鉴权api
	authApiGroup := router.Group("/auth")
	authApiGroup.GET("/users", AuthGetAllUser)
	authApiGroup.GET("/roles", AuthGetAllRole)
	authApiGroup.GET("/policys", AuthGetAllPolicy)

	authApiGroup.POST("/create", AuthCreatePolicy)
	authApiGroup.GET("/AddUserInRole", AuthAddUserInRoles)
	authApiGroup.GET("/DeleteUserInRole", AuthDeleteUserInRoles)

	auth := router.Group("/api")
	auth.Use(Casbin_RBAC(srv))

	auth.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "POST /api/user success"})
	})

	auth.DELETE("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "DELETE /api/user success"})
	})

	// 处理 router.Run 可能返回的错误
	if err := router.Run(":8081"); err != nil {
		// 打印错误信息，方便排查问题
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}

}

func Casbin_RBAC(srv *CasbinService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := srv.Enforcer.LoadPolicy()
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "权限校验失败"})
			return
		}
		// 从请求头中获取用户信息
		roleName, _ := ctx.GetQuery("roleName")
		ok, err := srv.Enforcer.Enforce(roleName, ctx.Request.URL.Path, ctx.Request.Method)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "权限校验失败"})
			ctx.Abort()
			return
		} else if !ok {
			ctx.AbortWithStatusJSON(403, gin.H{"error": "权限不足"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// 获取所有用户以及关联的角色 curl -X GET "http://localhost:8081/auth/users"
func AuthGetAllUser(ctx *gin.Context) {
	roles, err := srv.GetUsers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "获取所有用户失败"})
		return
	}
	ctx.JSON(200, gin.H{"roles": roles})

}

// 获取所有角色 curl -X GET "http://localhost:8081/auth/roles"
func AuthGetAllRole(ctx *gin.Context) {
	roles, err := srv.GetRoles()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "获取所有角色失败"})
		return
	}
	ctx.JSON(200, gin.H{"roles": roles})

}

// 获取所有角色组的策略 curl -X GET "http://localhost:8081/auth/policys"
func AuthGetAllPolicy(ctx *gin.Context) {
	roles, err := srv.GetRolesPolicy()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "获取策略失败"})
		return
	}
	ctx.JSON(200, gin.H{"roles": roles})

}

// 创建角色组策略
// curl -X POST -H "Content-Type: application/json" -d "{\"roleName\": \"user\", \"path\": \"/auth/users\", \"method\": \"GET\"}" "http://localhost:8081/auth/create"
func AuthCreatePolicy(ctx *gin.Context) {

	var p rolePolicy
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(400, gin.H{"error": "参数错误"})
		return
	}
	err := srv.CreateRolePolicy(p.RoleName, p.Path, p.Method)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "创建策略失败"})
		return
	}
	ctx.JSON(200, gin.H{"msg": "创建成功"})

}

// 添加用户到组, curl -X GET "http://localhost:8081/auth/AddUserInRole?username=htt1&rolename=user"
func AuthAddUserInRoles(ctx *gin.Context) {
	username := ctx.Query("username")
	rolename := ctx.Query("rolename")
	err := srv.AddUserInRoles(username, rolename)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "添加用户失败"})
		return

	}
	ctx.JSON(200, gin.H{"msg": "添加成功"})
}

// 从组中删除用户, /casbin/user-role?username=leo&rolename=admin
func AuthDeleteUserInRoles(ctx *gin.Context) {
	username := ctx.Query("username")
	rolename := ctx.Query("rolename")
	err := srv.DeleteUserInRoles(username, rolename)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "删除用户失败"})
		return
	}
	ctx.JSON(200, gin.H{"msg": "删除成功"})
}
