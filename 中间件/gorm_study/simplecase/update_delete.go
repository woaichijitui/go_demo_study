package simplecase

import "gorm.io/gorm"

var t Teacher = Teacher{Model: gorm.Model{ID: 10}}

func UpdateCase() {
	t = Teacher{}

	res := DB.First(&t)
	Println(res.RowsAffected, res.Error, t)
	t.Name = "htt1"
	t.Age = 33
	res = DB.Save(&t)
	Println(res.RowsAffected, res.Error, t)
}

func DeleteCase() {
	//	单条更新 删除对象需要指定主键，否则会触发 批量删除
	res := DB.Delete(&t)
	Println(res.RowsAffected, res.Error, t)

	//	根据条件删除
	res = DB.Where("name = ?", "htt").Delete(&Teacher{})
	Println(res.RowsAffected, res.Error)
}
