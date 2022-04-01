// @title 《Go语言编程》- 闭包
// @description
// @author wangpengliang
// @date 2022-04-01 09:42:53

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printHello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello:", i)
}

// 调用正常输出
func closure1() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printHello(i)
	}
	wg.Wait()
}

// 使用匿名函数发现多次输出同一个值
func closure2() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("%p", &i)
			// 这里匿名函数的i是由外层循环提供的其实就是一个闭包
			fmt.Println("hello:", i)
		}()
	}
	wg.Wait()
}

// 解决闭包问题将参数i传入匿名函数中使用新变量接收
func closure3() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("hello:", i)
		}(i)
	}
	wg.Wait()
}

func main() {
	closure3()
}
