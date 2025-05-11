package migration

import (
	"gorm.io/gorm"
)

// 不定义foreignKey 外键 和 references 引用
type User1 struct {
	//若不带gorm.Model 会报错
	gorm.Model
	Name string
	Age  int
}
type Comment1 struct {
	gorm.Model
	Title   string
	Content string
	User    User1
	// UserID 和 UserId 都可以
	UserID uint
}

func BelongTo1() {

	DB.AutoMigrate(
		&User1{},
		&Comment1{},
	)
}

// 自定义foreignKey 外键 和 references 引用
type User2 struct {
	//若不带gorm.Model 会报错
	gorm.Model
	Name string
	Age  int `gorm:"index:idx_age,unique"`
}
type Comment2 struct {
	gorm.Model
	Title   string
	Content string
	User    User2 `gorm:"foreignKey:UserAge;references:Age"`
	// UserID 和 UserId 都可以
	UserAge int
}

func BelongTo2() {

	DB.AutoMigrate(
		&User2{},
		&Comment2{},
	)
}
