// @title Gin-restful接口
// @description
// @author wangpengliang
// @date 2022-04-10 18:40:45
package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认路由
	r := gin.Default()

	books := make([]string, 3)
	books[0] = "c#编程"
	books[1] = "java编程"
	books[2] = "go编程"

	// get http://localhost:8080/books
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	// post http://localhost:8080/books
	r.POST("/books", func(c *gin.Context) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		books = append(books, "大话西游"+strconv.Itoa(r.Intn(100)))
		c.JSON(http.StatusOK, books)
	})

	// put
	r.PUT("/books", func(c *gin.Context) {
		index := c.Query("index")
		data, err := strconv.Atoi(index)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "index is undefined",
			})
		}
		books[data] = "平凡的世界"
		c.JSON(http.StatusOK, books)
	})

	// delete
	r.DELETE("/books", func(c *gin.Context) {
		index := c.Query("index")
		data, err := strconv.Atoi(index)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "index is undefined",
			})
		}
		books = append(books[:data], books[data:]...)
		c.JSON(http.StatusOK, books)
	})

	r.Run()
}
