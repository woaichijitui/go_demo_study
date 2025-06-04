package main

import "gorm_case/migration"

func main() {
	//simplecase.CreateRecond()
	//simplecase.QueryCase()
	//simplecase.UpdateCase()
	//simplecase.DeleteCase()

	//	迁移
	//migration.BelongTo1()
	migration.BelongTo2()
	migration.HasOne1()
	migration.HasMany()
	migration.ManyToMany1()
	migration.ManyToMany2()
}
