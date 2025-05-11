package migration

import (
	"gorm.io/gorm"
)

// 若是数字结尾则表名称在创建时不加s
type User4 struct {
	gorm.Model
	CreditCards2 []CreditCard2 `gorm:"foreignKey:UserRefer"`
}

type CreditCard2 struct {
	gorm.Model
	Number    string
	UserRefer uint
}

// HasMany
func HasMany() {

	DB.AutoMigrate(
		&User4{},
		&CreditCard2{},
	)
}
