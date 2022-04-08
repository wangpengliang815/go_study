// @title 《Go语言编程》- 单向通道
// @description
// @author wangpengliang
// @date 2022-04-01 15:21:27

package main

import (
	"fmt"
	"time"
)

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

func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

// 示例一：Producer负责发送,Consumer负责接收
func test() {
	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res) // 25
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

// 示例二：Producer负责发送,Consumer负责接收同时也可以向chan中写值
func test2() {
	ch := Producer2()
	res := Consumer2(ch)
	fmt.Println(res) // 125
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

// 示例三：Producer负责发送返回只读通道
func test3() {
	ch := Producer3()
	res := Consumer3(ch)
	fmt.Println(res) // 25
}

// 默认的channel 是双向的可读可写
func chanType1() {
	var ch chan int = make(chan int)
	go func() {
		x := <-ch
		fmt.Println(x)
	}()
	ch <- 100
	close(ch)
	time.Sleep(time.Second)
}

// 单向写channel,只允许写入
func chanType2() {
	var sendCh chan<- int = make(chan<- int)
	go func() {
		// x := <-sendCh // 因为是单向写channel 这里编译会报错：invalid operation: cannot receive from send-only channel sendCh (variable of type chan<- int)
		// fmt.Println(x)
	}()
	sendCh <- 100
	time.Sleep(time.Second)
}

// 单向读channel,只允许读取
func chanType3() {
	var recvCh <-chan int = make(<-chan int)
	go func() {
		x := <-recvCh
		fmt.Println(x)
	}()
	// 因为是单向读channel 这里编译会报错：invalid operation: cannot send to receive-only channel sendCh (variable of type <-chan int)
	// sendCh <- 100
	time.Sleep(time.Second)
}

// channel类型转换
func convertChannel() {
	// 双向通道
	var ch chan int = make(chan int)

	// 单向写
	var sendCh chan<- int = make(chan<- int)

	// 单向读
	var recvCh <-chan int = make(<-chan int)

	// 双向channel 可以 隐式转换为 任意一种单向channel
	sendCh = ch
	recvCh = ch

	// 单向 channel 不能转换为 双向 channel
	//	ch = recvCh // cannot use recvCh (variable of type <-chan int) as chan int value in assignment
	//	ch = sendCh // cannot use sendCh (variable of type chan<- int) as chan int value in assignment
	fmt.Printf("sendCh:%T,recvTh:%T", sendCh, recvCh)
}

// 这里可以Consumer既可以取也也可以发送
func main() {
	chanType1()
	// convertChannel()
}
