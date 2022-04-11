// @title Gin Web编程-获取queryString参数
// @description
// @author wangpengliang
// @date 2022-04-10 17:09:35
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取queryString参数,比如:/users?id=1&name=wpl&address=beijing
func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		id := c.Query("id")                             // 第一种方式：使用Query()函数获取queryString参数
		name := c.DefaultQuery("name", "wangpengliang") // 第二种方式：使用DefaultQuery()函数获取queryString参数,如果没有获取到会使用默认值

		address, ok := c.GetQuery("address") // 第三种方式：使用GetQuery()函数,返回两个值string,bool
		if !ok {
			// 没有获取到address参数
			address = "beijing"
		}
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"name":    name,
			"address": address,
		})
	})
	r.Run()
}
