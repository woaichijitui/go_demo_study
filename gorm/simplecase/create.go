package simplecase

import (
	"gorm.io/gorm"
	"log"
	"time"
)

var teacherTemp = Teacher{
	Model:        gorm.Model{},
	Name:         "htt",
	Email:        "1975611740@qq.com",
	Age:          44,
	WorkingYears: 10,
	Birthday:     time.Now().Unix(),
	StuNumber: struct {
		String string
		Valid  bool
	}{String: "10", Valid: true},
	Roles: Roles{
		Title:    "讲师",
		Location: "广东河源",
	},
	Roles1: Roles{
		Title:    "讲师",
		Location: "广东河源"},
}

func CreateRecond() {
	t := teacherTemp

	//	插入一条记录
	//	ID 会反向填充
	res := DB.Create(&t)
	if res.Error != nil {
		log.Println(res)
		return
	}

	Println(t)
	//	正向选择插入
	t1 := teacherTemp
	res = DB.Select("Name", "Age").Create(&t1)
	Println(res.RowsAffected, res.Error, t1)

	//	反向选择
	//	反向选择会主键冲突
	t2 := teacherTemp
	res = DB.Omit("StuNumber", "Roles1").Create(&t2)
	Println(res.RowsAffected, res.Error, t2)

	//	批量操作
	//	若有一条记录失败，则全部失败
	var teachers = []Teacher{{Name: "1", Age: 31}, {Name: "2", Age: 55}, {Name: "3", Age: 45}}

	res = DB.Create(&teachers)
	if res.Error != nil {
		log.Println(res.Error)
	}
	Println(res.RowsAffected, res.Error, teachers)
}
