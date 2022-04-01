// @title 《Go语言编程》- 最简单的函数串行执行
// @description
// @author wangpengliang
// @date 2022-03-31 17:24:44
package main

import (
	"fmt"
)

func printHello() {
	fmt.Println("hello")
}

func main() {
	printHello()
	fmt.Println("end...")
}
