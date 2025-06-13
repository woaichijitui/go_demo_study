package simplecase

func init() {
	//	迁移 建表
	DB.AutoMigrate(Teacher{})
	//	其他操作
	//	DB.Migrator().ColumnTypes()
}
