// @title 《Go语言编程》- 启动多个goroutine
// @description
// @author wangpengliang
// @date 2022-03-31 17:36:43
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printHello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go printHello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
