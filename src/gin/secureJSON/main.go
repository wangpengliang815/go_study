// @title Gin-SecureJSON
// @description 使用SecureJSON防止json劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体
// @author wangpengliang
// @date 2022-04-11 00:16:36
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 默认前缀为while(1):也可以使用自定义的SecureJSON前缀
	r.SecureJsonPrefix("wangpengliang \n")
	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		c.SecureJSON(http.StatusOK, names)
	})
	r.Run()
}
