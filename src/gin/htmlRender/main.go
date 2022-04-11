// @title Gin-HTML渲染
// @description
// @author wangpengliang
// @date 2022-04-10 23:46:48

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用LoadHTMLFiles()或者LoadHTMLGlob()渲染html页面
func main() {
	r := gin.Default()

	// r.LoadHTMLGlob("login.html")	// 加载单个html页面
	r.LoadHTMLGlob("templates/**/*") // 加载templates目录下多层级的模板文件

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			// 这里设置的参数可以在页面模板绑定
			"title":   "login",
			"message": "不学习就没有饭吃",
		})
	})

	// 使用不同目录下名称相同的模板
	r.GET("/books", func(c *gin.Context) {
		c.HTML(http.StatusOK, "books/index.tmpl", gin.H{
			"title": "books",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users",
		})
	})

	// 关于使用自定义模板参考:https://github.com/gin-gonic/examples/tree/master/template
	r.Run()
}
