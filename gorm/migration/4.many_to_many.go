package migration

import (
	"gorm.io/gorm"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
type User5 struct {
	gorm.Model
	Languages []Language1 `gorm:"many2many:user_languages;"`
}

type Language1 struct {
	gorm.Model
	Name string
}

// ManyToMany1
func ManyToMany1() {

	DB.AutoMigrate(
		&User5{},
		&Language1{},
	)
}

// User 拥有并属于多种 language，`user_languages` 是连接表
type User6 struct {
	gorm.Model
	Languages []*Language2 `gorm:"many2many:user_languages2;"`
}

type Language2 struct {
	gorm.Model
	Name  string
	Users []*User6 `gorm:"many2many:user_languages2;"`
}

// ManyToMany2
func ManyToMany2() {

	DB.AutoMigrate(
		&User6{},
		&Language2{},
	)
}
