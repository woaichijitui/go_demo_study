package migration

import (
	"gorm.io/gorm"
)

// User3 有一张 CreditCard，UserID 是外键
// 若是数字结尾则表名称在创建时不加s
type User3 struct {
	gorm.Model
	CreditCard CreditCard1 `gorm:"foreignKey:UserID"`
}

type CreditCard1 struct {
	gorm.Model
	Number string
	UserID uint
}

// HasOne1
func HasOne1() {

	DB.AutoMigrate(
		&User3{},
		&CreditCard1{},
	)
}
