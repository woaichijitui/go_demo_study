package simplecase

import (
	"database/sql"
	"gorm.io/gorm"
)

// 模型定义
// 遵循蛇形复数 表名会改成小写末尾加上s
type Teacher struct {
	gorm.Model
	Name         string `gorm_study:"size:256"`
	Email        string `gorm_study:"size:256"`
	Age          uint8  `gorm_study:"check:age>30"`
	WorkingYears uint8
	Birthday     int64 `gorm_study:"serializer:unixTime;type:time"` //为什么要序列化传输，又定义为time
	StuNumber    sql.NullString
	Roles        Roles `gorm_study:"embedded;embeddedPrefix:job_"` //不定义内嵌
	Roles1       Roles `gorm_study:"type:bytes;serializer:gob"`
}
type Roles struct {
	Title    string
	Location string
}
