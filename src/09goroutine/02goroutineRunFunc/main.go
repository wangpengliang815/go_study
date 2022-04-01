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

/*
 Go程序启动时就会为main函数创建一个默认的goroutine
 上面代码中在run函数中使用go关键字创建了另外一个goroutine去执行printHello()
 此时 main goroutine 还在继续往下执行，程序中此时存在两个并发执行的 goroutine
 当 main 函数结束时整个程序也就结束了，同时 main goroutine 也结束了，所有由 main goroutine 创建的 goroutine 也会一同退出
 也就是说 main 函数退出太快，另外一个 goroutine 中的函数还未执行完程序就退出了，导致未打印出“hello”。
*/
func main() {
	run() // 这里只会输出end...
}
