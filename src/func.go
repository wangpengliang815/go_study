package main

import (
	"fmt"
)

// 最简单的函数定义,求两数之和
func Sum(x int, y int) int {
	return x + y
}

// 使用类型简写
func ShortSum(x, y int) int {
	return x + y
}

// 可变参数指的是参数数量不固定,这里假设传入一个切片
func VariableElement(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 可变参数搭配固定参数时，必须放在参数最后一位
func VariableElement2(x int, y ...int) int {
	for _, v := range y {
		x += v
	}
	return x
}

// 无参无返回值的函数定义
func SayHello() {
	fmt.Println("hello world!")
}

// 函数多返回值
func Calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 函数返回值命名，可以在函数体中直接使用，最后通过return返回
func Calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

// 当函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显式返回一个长度为0的切片
func Calc3(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	return []int{0}
}

// 定义函数类型calculation:接受两个int类型参数并且返回一个int类型.凡是满足这个条件的函数都是calculation类型的函数,比如Sum2()/Sub2()像C#中委托的定义
type calculation func(int, int) int

// calculation类型的函数Sum2
func Sum2(x, y int) int {
	return x + y
}

// calculation类型的函数Sub2
func Sub2(x, y int) int {
	return x - y
}

func Sum3(x, y int) int {
	return x + y
}

func Calc4(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func Test1() {
	ret2 := Calc4(10, 20, Sum3)
	fmt.Println(ret2) //30
}
