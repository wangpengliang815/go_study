package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db = createDbConn()

// 返回数据库连接
func createDbConn() *gorm.DB {
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
		log.Fatal("create dbconn failed!:" + err.Error())
	}

	// 启用自动迁移生成表
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Person{})
	return db
}

func main() {
}

// // 查询:返回所有数据
// func Get_All_Handler(c *gin.Context) {
// 	var users []User
// 	db.Find(&users)
// 	c.JSON(http.StatusOK, users)
// }

// // 查询:获取第一条记录(主键升序)  SELECT * FROM users ORDER BY id LIMIT 1;
// func Get_First_Handler(c *gin.Context) {
// 	var user User
// 	db.First(&user)
// 	c.JSON(http.StatusOK, user)
// }

// // 查询:获取一条记录不指定排序字段 SELECT * FROM users LIMIT 1;
// func Get_Take_Handler(c *gin.Context) {
// 	var user User
// 	db.Take(&user)
// 	c.JSON(http.StatusOK, user)
// }

// // 查询:获取最后一条记录(主键降序) SELECT * FROM users ORDER BY id DESC LIMIT 1;
// func Get_Last_Handler(c *gin.Context) {
// 	var user User
// 	db.Last(&user)
// 	c.JSON(http.StatusOK, user)
// }
