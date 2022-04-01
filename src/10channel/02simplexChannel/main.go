// @title 《Go语言编程》- 单向通道
// @description
// @author wangpengliang
// @date 2022-04-01 15:21:27

package main

import (
	"fmt"
)

// Producer 返回一个通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
func Producer() chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()
	return ch
}

// Consumer 从通道中接收数据进行计算
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func Producer2() chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()
	return ch
}

// Consumer 从通道中接收数据进行计算同时可以向通道中写值
func Consumer2(ch chan int) int {
	sum := 0
	ch <- 100
	for v := range ch {
		sum += v
	}
	return sum
}

// Producer3 返回一个接收通道
func Producer3() <-chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch)
	}()

	return ch
}

// Consumer3 参数为接收通道
func Consumer3(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

// 这里可以Consumer既可以取也也可以发送
func main() {
	// 示例一：Producer负责发送,Consumer负责接收
	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res) // 25

	// 示例二：Producer负责发送,Consumer负责接收同时也可以向chan中写值
	ch2 := Producer2()
	res2 := Consumer2(ch2)
	fmt.Println(res2) // 125

	// 示例三：Producer负责发送返回只读通道
	ch3 := Producer3()
	res3 := Consumer3(ch3)
	fmt.Println(res3) // 25
}
