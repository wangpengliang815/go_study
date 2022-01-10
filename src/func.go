package main

import (
	"fmt"
)

// 最简单的函数定义,求两数之和.带参数且有返回值
func Sum1(x int, y int) int {
	return x + y
}

// 函数定义2：有参数但无返回值
func Sum2(x int, y int) {
	fmt.Println(x + y)
}

// 函数定义3：无参数且无返回值
func Sum3() {
	fmt.Println("hello world")
}

// 函数定义4：使用参数类型简写
func Sum4(x, y int) int {
	return x + y
}

// 函数定义5：函数返回值命名,可以在函数体中直接使用,通过return返回
func Sum5(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

// 函数多返回值
func Sum6(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 可变参数指的是参数数量不固定
func Sum7(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 可变参数搭配固定参数时,必须放在参数最后一位
func Sum8(x int, y ...int) int {
	for _, v := range y {
		x += v
	}
	return x
}

// 当函数返回值类型为slice时,nil可以看做是一个有效的slice，没必要显式返回一个长度为0的切片
func Sum9(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	return []int{0}
}

// 定义函数类型calculation:接受两个int类型参数并且返回一个int类型.凡是满足这个条件的函数都是calculation类型的函数,比如Sum2()/Sub2()像C#中委托的定义
type calculation func(int, int) int

// calculation类型的函数Csum
func Csum(x, y int) int {
	return x + y
}

// calculation类型的函数Csub
func Csub(x, y int) int {
	return x - y
}

func Calc(x, y int, sum func(int, int) int) int {
	return sum(x, y)
}

func Test1() {
	ret2 := Calc(10, 20, Sum1)
	fmt.Println(ret2) //30
}
