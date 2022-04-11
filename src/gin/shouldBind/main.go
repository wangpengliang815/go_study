// @title Gin Web编程-参数绑定
// @description
// @author wangpengliang
// @date 2022-04-10 18:03:11
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 自定义请求参数,注意这里字段首字母必须大写否则反射是找不到这个属性的,无法完成参数绑定
type User struct {
	UserName string `json:"user" form:"username"` // 这里的标签指定如果使用json方式参数名称叫什么,使用form表单方式参数名称叫什么
	PassWord string `json:"pwd"  form:"password"`
}

/*
 参数绑定：ShouldBind可以基于请求自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象

*/

func main() {
	r := gin.Default()
	// 加载htmt页面
	r.LoadHTMLFiles("./login.html")

	// 通过ShouldBind()获取queryString参数自动绑定,localhost:8080/?UserName=1&PassWord=2 需要注意:必须与结构字段参数同名或者设置结构的tag
	r.GET("/", func(c *gin.Context) {
		var user User
		// 通过反射进行参数绑定,因为函数传递的是值拷贝所以这里需要使用指针
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username": user.UserName,
				"password": user.PassWord,
			})
		}
	})

	// 处理get请求返回默认的login登录页面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 通过ShouldBind()获取Form表单参数自动绑定
	r.POST("/login", func(c *gin.Context) {
		var user User
		// 通过反射进行参数绑定,因为函数传递的是值拷贝所以这里需要使用指针
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username": user.UserName,
				"password": user.PassWord,
			})
		}
	})

	// 通过ShouldBind()获取Json方式提交参数自动绑定
	r.POST("/postJson", func(c *gin.Context) {
		var user User
		// 通过反射进行参数绑定,因为函数传递的是值拷贝所以这里需要使用指针
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username": user.UserName,
				"password": user.PassWord,
			})
		}
	})

	r.Run()
}
