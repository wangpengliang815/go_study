// @title 《Go语言编程》 -使用sync.WaitGroup全局等待组变量优雅的阻塞
// @description
// @author wangpengliang
// @date 2022-03-31 17:25:12

package main

import (
	"fmt"
	"sync"
)

// 声明全局等待组变量
var wait sync.WaitGroup

func printHello() {
	defer wait.Done() // goroutine结束登记-1
	fmt.Println("hello")
}

// 使用sync包
func run() {
	wait.Add(1) // 启动一个goroutine登记+1
	go printHello()
	fmt.Println("end...")
	wait.Wait() // 等待所有线程执行完
}

func main() {
	run() // 先输出:end... 在输出:hello  存在随机性
}
