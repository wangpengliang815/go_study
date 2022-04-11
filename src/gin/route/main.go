// @title Gin Web编程-路由
// @description
// @author wangpengliang
// @date 2022-04-11 00:48:38
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 普通路由
	r.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodGet,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodPost,
		})
	})

	// any()可以匹配login所有的请求,不论get/post/put/delete/...
	r.Any("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodPost,
		})
	})

	// 没有配置处理函数的路由添加处理程序，默认情况下返回404代码，下面的代码为没有匹配到路由的请求都返回404页面
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	// 将拥有共同URL前缀的路由划分为一个路由组。习惯性一对{}包裹同组的路由，只是为了看着清晰用不用{}包裹功能上没什么区别
	books := r.Group("/books")
	{
		books.GET("/index", func(c *gin.Context) {})
		books.GET("/login", func(c *gin.Context) {})
		books.POST("/login", func(c *gin.Context) {})

	}
	users := r.Group("/users")
	{
		users.GET("/index", func(c *gin.Context) {})
		users.GET("/cart", func(c *gin.Context) {})
		users.POST("/checkout", func(c *gin.Context) {})
	}

	// 路由组支持嵌套
	person := r.Group("/shop")
	{
		person.GET("/index", func(c *gin.Context) {})
		person.GET("/cart", func(c *gin.Context) {})
		person.POST("/checkout", func(c *gin.Context) {})
		xx := person.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {})
	}

	// 路由本质参考：https://github.com/julienschmidt/httprouter

	// 重定向到外部页面
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	// 请求转发到指定handler
	r.POST("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	r.Run()
}
