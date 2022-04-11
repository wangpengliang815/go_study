// @title Gin Web编程-获取URL参数
// @description
// @author wangpengliang
// @date 2022-04-10 18:03:11
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取Url参数,比如：localhost:8080/wpl/18
func main() {
	r := gin.Default()

	r.GET("/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run()
}
