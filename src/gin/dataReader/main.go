// @title Gin-从reader读取数据
// @description
// @author wangpengliang
// @date 2022-04-11 00:39:58

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/someDataFromReader", func(c *gin.Context) {
		// 这里地址是一张图片
		response, err := http.Get("https://tse3-mm.cn.bing.net/th/id/OIP-C.kVOiwrDVW0tbUcY9pzqJ4gHaG4?pid=ImgDet&rs=1")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		// 下载图片返回到浏览器
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	r.Run()
}
