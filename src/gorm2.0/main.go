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

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
结构体定义,标签解释：
	1.`gorm:"primaryKey"`: 表示设置某列为主键
	2.`gorm:"default:'test'"`: 表示为某字段设置默认值
*/
type User struct {
	Id      int    `gorm:"primaryKey"`
	Name    string `gorm:"default:'test'"`
	Age     int
	Address string
}

// 重命名表名
func (User) TableName() string {
	return "User"
}

func main() {
	r := gin.Default()
	db := createDbConn()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/users", func(c *gin.Context) {
		var user []User
		db.Find(&user)
		c.JSON(http.StatusOK, user)
	})

	// @title Gin swagger
	// @version 1.0
	// @description Gin swagger 示例项目

	// @contact.name youngxhu
	// @contact.url https://youngxhui.top
	// @contact.email youngxhui@g mail.com

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @host localhost:8080
	r.POST("/users", func(c *gin.Context) {
		user := User{Name: "wangpengliang", Age: 18, Address: "北京"}
		result, _ := insert(user, db)
		c.JSON(http.StatusOK, result)
	})

	// @Summary 提交新的文章内容
	// @Id 1
	// @Tags 文章
	// @version 1.0
	// @Accept application/x-json-stream
	// @Param article body model.Article true "文章"
	// @Success 200 object model.Result 成功后返回值
	// @Failure 409 object model.Result 添加失败
	// @Router /article [post]
	r.POST("/users/select", func(c *gin.Context) {
		user := User{Name: "wangpengliang", Age: 18, Address: "北京"}
		result, _ := selectInsert(user, db)
		c.JSON(http.StatusOK, result)
	})

	// @Summary 提交新的文章内容
	// @Id 1
	// @Tags 文章
	// @version 1.0
	// @Accept application/x-json-stream
	// @Param article body model.Article true "文章"
	// @Success 200 object model.Result 成功后返回值
	// @Failure 409 object model.Result 添加失败
	// @Router /article [post]
	r.POST("/users/excludeField", func(c *gin.Context) {
		user := User{Name: "wangpengliang", Age: 18, Address: "北京"}
		result, _ := excludeFieldInsert(user, db)
		c.JSON(http.StatusOK, result)
	})

	// @Summary 提交新的文章内容
	// @Id 1
	// @Tags 文章
	// @version 1.0
	// @Accept application/x-json-stream
	// @Param article body model.Article true "文章"
	// @Success 200 object model.Result 成功后返回值
	// @Failure 409 object model.Result 添加失败
	// @Router /article [post]
	r.POST("/users/batch", func(c *gin.Context) {
		users := []User{{Name: "wangpengliang", Age: 18, Address: "北京"}, {Name: "lizimeng", Age: 18, Address: "上海"}}
		result := batchInsert(users, db)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, result.Error)
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	r.Run()
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
	db.AutoMigrate(&User{})

	// // 创建2条测试数据
	// db.Create(&Person{Name: "wangpengliang", Address: "beijing", Age: 18})
	// db.Create(&Person{Name: "lizimeng", Address: "shanghai", Age: 18})
	return db
}

// 默认创建
func insert(user User, db *gorm.DB) (model User, dbResult *gorm.DB) {
	result := db.Create(&user)
	return user, result
}

// 使用选定字段创建
func selectInsert(user User, db *gorm.DB) (model User, dbResult *gorm.DB) {
	// 这里只插入了2个字段虽然传入的user是3个字段,INSERT INTO `User` (`name`,`age`) VALUES ("xx", 18)
	result := db.Select("Name", "Age").Create(&user)
	return user, result
}

// 排除指定字段创建
func excludeFieldInsert(user User, db *gorm.DB) (model User, dbResult *gorm.DB) {
	// 这里排除Age字段,INSERT INTO `User` (`name`,`address`) VALUES ("xx", "xx")
	result := db.Omit("Age").Create(&user)
	return user, result
}

// 批量创建
func batchInsert(users []User, db *gorm.DB) (dbResult *gorm.DB) {
	result := db.Create(&users)
	return result
}
