// @title Gin-中间件
// @description
// @author wangpengliang
// @date 2022-04-11 09:36:25
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 最普通的中间件,HandlerFunc()其实就是个函数类型,接收参数是*gin.Context
func m1(c *gin.Context) {
	fmt.Println("m1 start...")
	c.Next()
	fmt.Println("m1 end...")
}

// c.Next()
func m2(c *gin.Context) {
	fmt.Println("m2 start...")
	c.Next() // 调用后续处理函数
	fmt.Println("m2 end...")
}

// c.About()
func m3(c *gin.Context) {
	fmt.Println("m3 start...")
	c.Abort() // 阻止请求后续的处理函数
	fmt.Println("m3 end...")
}

// 需要全局注册的中间件
func getTime(c *gin.Context) {
	start := time.Now()
	c.Next() // 调用后续处理函数
	cost := time.Since(start)
	fmt.Printf("%v \n", cost)
}

// 一般中间件会使用闭包函数返回
func getTime2() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		fmt.Printf("%v \n", cost)
	}
}

func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局注册中间件
	r.Use(gin.Logger()) // Logger中间件将日志写入gin.DefaultWriter，哪怕将GIN_MODE设置为release

	r.Use(gin.Recovery()) // Recovery中间件会recover任何panic。如果有pani 的话，会写入500

	r.Use(getTime)

	// 不调用任何自定义中间件
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	// 调用自定义中间件m1
	r.GET("/m1", m1, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	// 调用自定义中间件m1+m2
	r.GET("/m2", m1, m2, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	// 调用自定义中间件m3
	r.GET("/m3", m1, m3, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.Run()
}
