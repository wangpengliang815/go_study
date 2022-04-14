// @title gin-restful
// @description
// @author wangpengliang
// @date 2022-04-12 15:52:04
package main

import (
	"fmt"
	_ "goProject/gin/restful/docs"
	"log"
	"net/http"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
	Id      int `gorm:"primaryKey"`
	Name    string
	Age     int
	Address string
}

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

	return db
}

// @Summary 创建数据
// @Description
// @Param user body User true "用户"
// @Tags Users
// @Accept json
// @Router /users [post]
func InsertHandler(c *gin.Context) {
	user := User{}
	c.ShouldBind(&user) // 需要通过指针创建
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
	}
	// user.ID 插入数据的主键
	// result.Error        // 返回error
	// result.RowsAffected // 返回插入记录的条数
	c.JSON(http.StatusOK, user)
}

// @Summary 返回所有数据
// @Description
// @Tags Users
// @Accept json
// @Router /users [get]
func GetListHandler(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// @title gorm2.0 sample
// @version 1.0
// @description
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080/
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/users", GetListHandler)

	r.POST("/users", InsertHandler)

	r.Run()
}
