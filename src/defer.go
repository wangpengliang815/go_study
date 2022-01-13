package main

import (
	"fmt"
)

// defer将后面的语句延迟到函数即将退出时逆序执行,一般常用于资源释放
func DeferTest() {
	fmt.Println("start")
	defer fmt.Println("...")
	fmt.Println("end")
}

// 函数中存在多个defer时,逆序执行后进先出
func DeferTest2() {
	fmt.Println("start")
	defer fmt.Println("1111")
	defer fmt.Println("2222")
	defer fmt.Println("3333")
	fmt.Println("end")
}

// defer执行时机
func F1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func F2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func F3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func F4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
