package simplecase

func QueryCase() {
	t := Teacher{}
	t1 := Teacher{}
	t2 := Teacher{}
	/*t1 := teacherTemp
	t2 := teacherTemp
	t3 := teacherTemp*/

	//	 获取第一条记录（主键升序）
	res := DB.First(&t)
	if res.Error == nil {
		Println(res.RowsAffected, res.Error, t)
	}
	// 获取最后一条记录（主键降序）
	res = DB.Last(&t1)
	if res.Error == nil {
		Println(res.RowsAffected, res.Error, t1)
	}
	// 获取一条记录，没有指定排序字段
	res = DB.Take(&t2)
	if res.Error == nil {
		Println(res.RowsAffected, res.Error, t2)
	}

	//	将结果填充到集合,不支持特殊类型，无法完成类型转换？ 什么是特殊类型
	result := map[string]interface{}{}
	res = DB.Model(&Teacher{}).Select("Name", "Age").First(&result)
	Println(res.RowsAffected, res.Error, result)

	//Last 和 First 可以用模型
	//take 若没有特殊类型 模型和表名都可以
	res = DB.Table("teachers").Select("Name", "Age", "Birthday").Take(&result)
	Println(res.RowsAffected, res.Error, result)

	//	条件查询
	//		查询多行
	var teacher []Teacher
	res = DB.Where("name=?", "htt").Or("name=?", "null").Order("id desc").Limit(10).Find(&teacher)
	Println(res.RowsAffected, res.Error, teacher)
}
