// @title gorm version 2.0
// @description
// @author wangpengliang
// @date 2022-04-12 15:52:04
package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	_ "goProject/gorm2.0/docs"
	"goProject/gorm2.0/entity"
	_ "goProject/gorm2.0/entity"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title gorm2.0 sample
// @version 1.0
// @description
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080/
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/users", getListHandler)

	r.POST("/users", insertHandler)

	r.POST("/users/select", selectFieldInsertHandler)

	r.POST("/users/excludeField", excludeFieldInsertHandler)

	r.POST("/users/batch", batchInsertHandler)

	r.Run()
}

// @Summary 查询所有数据
// @Description
// @Tags Users
// @Accept json
// @Router /users [get]
func getListHandler(c *gin.Context) {
	db := createDbConn()
	var user []entity.User
	db.Find(&user)
	c.JSON(http.StatusOK, user)
}

// @Summary 创建数据
// @Description
// @Param user body User true "用户"
// @Tags Users
// @Accept json
// @Router /users [post]
func insertHandler(c *gin.Context) {
	db := createDbConn()
	user := entity.User{}
	c.ShouldBind(&user)
	db.Create(&user)
	c.JSON(http.StatusOK, user)
}

// @Summary 使用选定字段创建数据
// @Description
// @Param user body User true "用户"
// @Tags Users
// @Accept json
// @Router /users/select [post]
func selectFieldInsertHandler(c *gin.Context) {
	db := createDbConn()
	user := entity.User{}
	c.ShouldBind(&user)
	// 这里只插入了2个字段虽然传入的user是3个字段,INSERT INTO `User` (`name`,`age`) VALUES ("xx", 18)
	db.Select("Name", "Age").Create(&user)
	c.JSON(http.StatusOK, user)
}

// @Summary 排除指定字段创建
// @Description
// @Param user body User true "用户"
// @Tags Users
// @Accept json
// @Router /users/excludeField [post]
func excludeFieldInsertHandler(c *gin.Context) {
	db := createDbConn()
	user := entity.User{}
	c.ShouldBind(&user)
	// 这里排除Age字段,INSERT INTO `User` (`name`,`address`) VALUES ("xx", "xx")
	db.Omit("Age").Create(&user)
	c.JSON(http.StatusOK, user)
}

// @Summary 批量创建
// @Description
// @Param user body []User true "用户"
// @Tags Users
// @Accept json
// @Router /users/batch [post]
func batchInsertHandler(c *gin.Context) {
	db := createDbConn()
	users := []entity.User{}
	c.ShouldBind(&users)
	result := db.Create(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

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
	db.AutoMigrate(&entity.User{})

	return db
}
