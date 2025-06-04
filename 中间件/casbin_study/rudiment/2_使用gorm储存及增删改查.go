package rudiment

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm_study-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"path/filepath"
)

var DNS = "root:123456@tcp(47.121.183.173:33069)/zztag?charset=utf8mb4&parseTime=True&loc=Local"

func GormCacheAndCURD() {
	gormDB, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	adapter, err := gormadapter.NewAdapterByDB(gormDB)
	if err != nil {
		panic(err)
	}

	// 从文件中加载模型
	m, err := model.NewModelFromFile(filepath.Join(pathRoot + "/" + "casbin_model.conf"))
	if err != nil {
		fmt.Println(err)
	}
	e, err := casbin.NewEnforcer(m, adapter)

	e.LoadPolicy()
	// 添加策略
	ok, err := e.AddPolicy("admin", "/api/user", "GET")
	ok, err = e.AddPolicy("user", "/api/user", "POST")
	log.Println("add admin /api/user GET: ", ok, err)
	ok, err = e.AddGroupingPolicy("leo", "admin")
	ok, err = e.AddGroupingPolicy("leo", "user")
	ok, err = e.AddGroupingPolicy("htt", "admin")
	ok, err = e.AddGroupingPolicy("zuzu", "user")
	log.Println("add leo to admin group: ", ok, err)
	e.SavePolicy()

	ok, err = e.Enforce("leo", "/api/user", "GET")
	log.Println("leo GET /api/user :", ok, err)
	ok, err = e.Enforce("leo", "/api/user", "DELETE")
	log.Println("leo DELETE /api/user :", ok, err)
}
