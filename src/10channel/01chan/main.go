// @title 《Go语言编程》- channel、无缓冲、带缓冲、多返回值、for range接收消息
// @description
// @author wangpengliang
// @date 2022-04-01 15:20:49
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 无缓冲的通道又称为阻塞的通道：无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段
func chan1() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功") // fatal error: all goroutines are asleep - deadlock!
}

// 解决无缓冲通道死锁问题,使用匿名函数接收通道值
func chan2() {
	value := 10
	ch := make(chan int)
	go func(chan int) {
		value := <-ch
		fmt.Printf("接收成功,value=%d \n", value)
	}(ch)
	ch <- value
	close(ch)
	fmt.Printf("发送成功,value=%d \n", value)
}

// 有缓冲的信道:只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量
func chan3() {
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功")
}

func receive(ch chan int) {
	// 循环接收值，直到通道被关闭后退出
	for {
		v, ok := <-ch
		if !ok {
			break
		} else {
			wg.Done()
			fmt.Printf("v=%v ok=%v \n", v, ok)
		}
	}
	close(ch)
}

// 多返回值：循环从通道ch中接收所有值，直到通道被关闭后退出
func chan4() {
	ch := make(chan int, 10)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		ch <- i
	}
	close(ch)

	go func(chan int) {
		// 循环接收值，直到通道被关闭后退出
		for {
			v, ok := <-ch
			if !ok {
				break
			} else {
				wg.Done()
				fmt.Printf("v=%v ok=%v \n", v, ok)
			}
		}
	}(ch)
	wg.Wait()
}

// 使用for range 循环接收值
func chan5() {
	ch := make(chan int, 10)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		ch <- i
	}
	close(ch)

	go func() {
		for v := range ch {
			fmt.Println(v)
			wg.Done()
		}
	}()
	wg.Wait()
}

func main() {
	chan5()
}
