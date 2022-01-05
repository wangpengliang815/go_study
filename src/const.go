package main

import (
	"fmt"
)

// 全局常量定义，一旦定义不能修改,如果没有使用不会编译报错
const PI = 3.141592653589793

// 常量
func ConstTest() {
	// 显式类型定义,常量名称推荐全大写
	const LENGTH int = 10
	// 隐式类型定义,其实也是使用类型推导
	const WIDTH = 20
	area := LENGTH * WIDTH
	fmt.Println(area)

	// 多重定义
	const a, b, c = 1, false, "hello world"
	fmt.Println(a, b, c)
}
