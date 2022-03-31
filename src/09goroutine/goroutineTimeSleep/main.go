// @title 《Go语言编程》- 使用使用time.Sleep阻塞
// @description
// @author wangpengliang
// @date 2022-03-31 17:25:12

package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("hello")
}

// goroutine使用time.Sleep阻塞
func run() {
	go printHello()
	fmt.Println("end...")
	time.Sleep(time.Second)
}

func main() {
	run() // 先输出:end... 在输出:hello  存在随机性
}
