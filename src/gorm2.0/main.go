// @title gorm version 2.0
// @description
// @author wangpengliang
// @date 2022-04-12 15:52:04
package main

import (
	_ "goProject/gorm2.0/docs"

	handler "goProject/gorm2.0/handler"

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

	r.GET("/users", handler.Get_All_Handler)

	r.GET("/users/first", handler.Get_First_Handler)

	r.GET("/users/take", handler.Get_Take_Handler)

	r.GET("/users/last", handler.Get_Last_Handler)

	r.POST("/users", handler.Insert_Default_Handler)

	r.POST("/users/select", handler.Insert_SelectField_Handler)

	r.POST("/users/excludeField", handler.Insert_ExcludeField_Handler)

	r.POST("/users/batch", handler.Insert_Batch_Handler)

	r.POST("/users/byMap", handler.Insert_ByMap_Handler)

	r.Run()
}
