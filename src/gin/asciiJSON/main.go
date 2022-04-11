// @title Gin-asciiJSON
// @description 使用AsciiJSON生成具有转义的非ASCII字符的ASCII-onlyJSON
// @author wangpengliang
// @date 2022-04-11 09:20:23

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run()
}
