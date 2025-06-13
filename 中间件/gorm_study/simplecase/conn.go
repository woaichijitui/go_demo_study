package simplecase

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
//
//	[username[:password] 账户和密码
//	[protocol 协议 tcp / udp
//	(address)	连接地址 ip:端口号
//	dbname 数据库名
//	[?param1=value1&...&paramN=valueN] 可选的查询参数部分，可以用来指定额外的连接选项和配置信息
var dsn string = "root:1234@tcp(127.0.0.1:3306)/gormcase?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func init() {
	var err error
	//	自定义驱动
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		//	定义日志输出
		//	设置日志输出等级
		//	若想改变输出的位置？
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	setPool(DB)
}

func setPool(DB *gorm.DB) {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	//	设置空闲连接池的最大连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	//最大连接数
	sqlDB.SetMaxOpenConns(5)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}
