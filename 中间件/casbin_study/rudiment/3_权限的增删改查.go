package rudiment

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm_study-adapter/v3"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"path/filepath"
)

type CasbinService struct {
	Enforcer *casbin.Enforcer
	Adapte   *gormadapter.Adapter
}

type User struct {
	Username string
	Roles    []string
}

func TestCasbin() {

	s := CasbinInit()

	// 角色组和用户增删改查
	users, err := s.GetUsers()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(users)

	roles, err := s.GetRoles()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(roles)

	err = s.AddUserInRoles("ddd", "admin")
	if err != nil {
		log.Println(err)
	}
	err = s.AddUserInRoles("eee", "eeeRole")
	if err != nil {
		log.Println(err)
	}

	err = s.DeleteUserInRoles("ddd", "admin")
	if err != nil {
		log.Println(err)
	}

	// 角色组权限的增删改查
	policy, err := s.GetRolesPolicy()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(policy)

	err = s.CreateRolePolicy("eeeRole", "/api/delete", "DELETE")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("create role policy success")
	}
	err = s.CreateRolePolicy("user", "/api/user/login", "DELETE")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("create role policy success")
	}

	oldp := rolePolicy{
		RoleName: "user",
		Path:     "/api/user/login",
		Method:   "DELETE",
	}
	newp := rolePolicy{
		RoleName: "user",
		Path:     "/api/user/login",
		Method:   "POST",
	}
	err = s.UpdateRolePolicy(oldp, newp)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("update role policy success")
	}

	dp := rolePolicy{
		RoleName: "htt",
		Path:     "/api/user",
		Method:   "POST",
	}
	err = s.DeleteRolePolicy(dp)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("delete role policy success")

	}

	policy, err = s.GetRolesPolicy()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(policy)

	// 角色组权限的校验
	ok, err := s.CanAccess("user", "/api/user/login", "POST")
	if err != nil {
		log.Println(err)
	} else {
		if ok {
			fmt.Println("zuzu can access /api/user/login POST")
		} else {
			fmt.Println("zuzu can not access /api/user/login POST")
		}
	}

}

func CasbinInit() (casbinService *CasbinService) {
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

	return &CasbinService{
		Enforcer: e,
		Adapte:   adapter,
	}
}

// 角色组和用户增删改查

// 获取所有用户以及关联的角色
func (c *CasbinService) GetUsers() (users []User, err error) {
	policy, err := c.Enforcer.GetGroupingPolicy()
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]*User, 0)

	for _, _p := range policy {
		username, userGroup := _p[0], _p[1]
		if u, ok := userMap[username]; ok {
			u.Roles = append(u.Roles, userGroup)
		} else {
			userMap[username] = &User{
				Username: username,
				Roles:    []string{userGroup},
			}
		}
	}
	for _, user := range userMap {
		users = append(users, *user)
	}

	return
}

// 获取所有角色组
func (c *CasbinService) GetRoles() (roles []string, err error) {

	return c.Enforcer.GetAllRoles()
}

// 角色组添加用户，没有组默认添加
func (c *CasbinService) AddUserInRoles(username, role string) (err error) {
	_, err = c.Enforcer.AddRoleForUser(username, role)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy() // 保存策略 持久化到数据库

}

// 角色组删除用户
func (c *CasbinService) DeleteUserInRoles(username, role string) (err error) {
	_, err = c.Enforcer.DeleteRoleForUser(username, role)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy() // 保存策略 持久化到数据库

}

// 角色组权限的增删改查
type rolePolicy struct {
	RoleName string `gorm_study:"column:v0" json:"roleName"`
	Path     string `gorm_study:"column:v1" json:"path" `
	Method   string `gorm_study:"column:v2" json:"method"`
}

// 获取所有角色组的权限
func (c *CasbinService) GetRolesPolicy() (roles []rolePolicy, err error) {
	err = c.Adapte.GetDb().Model(&gormadapter.CasbinRule{}).Where("ptype = 'p'").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return
}

// 创建角色组权限, 已有的会忽略
func (c *CasbinService) CreateRolePolicy(role string, path string, method string) (err error) {

	err = c.Enforcer.LoadPolicy()
	if err != nil {
		return err
	}

	_, err = c.Enforcer.AddPolicy(role, path, method)
	if err != nil {
		return err
	}

	return c.Enforcer.SavePolicy() // 保存策略 持久化到数据库

}

// 修改角色组权限
func (c *CasbinService) UpdateRolePolicy(old, new rolePolicy) (err error) {
	_, err = c.Enforcer.UpdatePolicy([]string{old.RoleName, old.Path, old.Method}, []string{new.RoleName, new.Path, new.Method})
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy() // 保存策略 持久化到数据库

}

// 删除角色组权限
func (c *CasbinService) DeleteRolePolicy(rolePolicy rolePolicy) (err error) {
	_, err = c.Enforcer.RemovePolicy(rolePolicy.RoleName, rolePolicy.Path, rolePolicy.Method)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy() // 保存策略 持久化到数据库
}

// 验证用户权限
func (c *CasbinService) CanAccess(roleName, path, method string) (ok bool, err error) {
	return c.Enforcer.Enforce(roleName, path, method)
}
