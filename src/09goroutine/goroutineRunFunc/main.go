// @title 《Go语言编程》- 使用goroutine执行函数
// @description
// @author wangpengliang
// @date 2022-03-31 17:25:12

package main

import "fmt"

func printHello() {
	fmt.Println("hello")
}

// 启动goroutine执行函数
func run() {
	go printHello()
	fmt.Println("end...")
}

func main() {
	run() // 这里只会输出end...
}
