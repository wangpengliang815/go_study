// @title Gin-pureJSON
// @description JSON使用unicode替换特殊HTML字符,例如<变为\u003c。如果要按字面对字符进行编码，可以使用PureJSON。Go 1.6及更低版本无法使用
// @author wangpengliang
// @date 2022-04-11 00:09:50

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 返回unicode实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 返回字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	r.Run()
}
