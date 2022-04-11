// @title Gin-默认模板
// @description 下载gin: go get -u github.com/gin-gonic/gin  导入包: import "github.com/gin-gonic/gin"
// @author wangpengliang
// @date 2022-04-10 17:02:15

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认路由,Default()默认使用Logger和Recovery两个中间件,如果不想使用这两个中间件可以使用gin.New()创建路由
	r := gin.Default()

	// 客户端使用get方式请求默认路径时,执行后面的匿名函数返回字符串
	r.GET("/", func(c *gin.Context) {
		// 这里使用了net/http包内封装好的http状态码
		c.JSON(http.StatusOK, "hello gin!")
	})

	// 启动http服务并监听端口,默认为8080,如果想要修改使用 r.Run(":9090")
	r.Run()
}
