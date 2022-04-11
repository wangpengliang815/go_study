// @title Gin Web编程-获取form表单传递的参数
// @description
// @author wangpengliang
// @date 2022-04-10 17:27:49
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取form表单参数
func main() {
	r := gin.Default()
	// 加载htmt页面
	r.LoadHTMLFiles("./login.html")

	// 处理get请求返回默认的login登录页面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 处理表单提交的post请求
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username") // 方式1：使用PostForm()获取表单参数
		password := c.PostForm("password")
		address := c.DefaultPostForm("address", "beijing") // 方式2：使用DefaultPostForm()获取表单参数,如果没有获取到会使用默认值

		age, ok := c.GetPostForm("age") // 方式3：使用GetPostForm()函数,返回两个值string,bool
		if !ok {
			// 没有获取到age参数
			age = fmt.Sprint(18)
		}

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"address":  address,
			"age":      age,
		})
	})

	r.Run()
}
