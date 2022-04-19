package common

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

// CreateDbConn 返回数据库连接
func CreateDbConn() *gorm.DB {
	var (
		server   = "localhost"
		port     = 1433
		user     = "sa"
		database = "go_db"
		password = "wpl19950815"
	)
	// 数据库连接字符串
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", server, port, database, user, password)
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("create connString failed!:" + err.Error())
	}

	// 启用自动迁移生成表
	_ = db.AutoMigrate(&User{})
	//_ = db.AutoMigrate(&Person{})
	_ = db.AutoMigrate(&Address{})
	return db
}
