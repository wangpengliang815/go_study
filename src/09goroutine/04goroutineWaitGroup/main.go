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
var wg sync.WaitGroup

func main() {
	wg.Add(1) // 启动一个goroutine登记+1
	go func() {
		defer wg.Done() // goroutine结束登记-1
		fmt.Println("hello")
	}()
	fmt.Println("end...")
	wg.Wait() // 阻塞等待所有协程执行完,先输出:end... 再输出:hello
}
