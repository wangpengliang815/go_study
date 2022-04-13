package handler

import (
	"fmt"
	"log"
	"net/http"

	"goProject/gorm2.0/entity"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
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
	db.AutoMigrate(&entity.User{})

	return db
}

// @Summary 创建数据
// @Description
// @Param user body User true "用户"
// @Tags Users
// @Accept json
// @Router /users [post]
func Insert_Default_Handler(c *gin.Context) {
	user := entity.User{}
	c.ShouldBind(&user) // 需要通过指针创建
	db.Create(&user)
	c.JSON(http.StatusOK, user)
}

// 创建:使用选定字段
func Insert_SelectField_Handler(c *gin.Context) {
	user := entity.User{}
	newUser := entity.User{}
	c.ShouldBind(&user)
	// 这里只会插入2个字段,虽然传入的user是3个字段,INSERT INTO `User` (`name`,`age`) VALUES ("xx", 18)
	db.Select("Name", "Age").Create(&user)
	// 查询将新的模型返回
	result := db.First(&newUser, user.Id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
	}
	c.JSON(http.StatusOK, newUser)
}

// 创建:排除指定字段
func Insert_ExcludeField_Handler(c *gin.Context) {
	user := entity.User{}
	newUser := entity.User{}
	c.ShouldBind(&user)
	// 这里排除Age字段,INSERT INTO `User` (`name`,`address`) VALUES ("xx", "xx")
	db.Omit("Age").Create(&user)
	// 查询将新的模型返回
	result := db.First(&newUser, user.Id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
	}
	c.JSON(http.StatusOK, newUser)
}

// 创建:批量
func Insert_Batch_Handler(c *gin.Context) {
	users := []entity.User{}
	c.ShouldBind(&users)
	result := db.Create(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// 创建:根据Map创建,注意:根据map创建时,association不会被调用,且主键也不会自动填充
func Insert_ByMap_Handler(c *gin.Context) {
	db.Model(&entity.User{}).Create(map[string]interface{}{
		"Name": "wangpengliang", "Age": 18, "Address": "beijing",
	})

	// TODO:这里有点问题
	// db.Model(&entity.User{}).Create([]map[string]interface{}{
	// 	{"Name": "wangpengliang", "Age": 18, "Address": "beijing"},
	// 	{"Name": "wangpengliang", "Age": 18, "Address": "beijing"},
	// })
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// 查询:返回所有数据
func Get_All_Handler(c *gin.Context) {
	var users []entity.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// 查询:获取第一条记录(主键升序)  SELECT * FROM users ORDER BY id LIMIT 1;
func Get_First_Handler(c *gin.Context) {
	var user entity.User
	db.First(&user)
	c.JSON(http.StatusOK, user)
}

// 查询:获取一条记录不指定排序字段 SELECT * FROM users LIMIT 1;
func Get_Take_Handler(c *gin.Context) {
	var user entity.User
	db.Take(&user)
	c.JSON(http.StatusOK, user)
}

// 查询:获取最后一条记录(主键降序) SELECT * FROM users ORDER BY id DESC LIMIT 1;
func Get_Last_Handler(c *gin.Context) {
	var user entity.User
	db.Last(&user)
	c.JSON(http.StatusOK, user)
}
