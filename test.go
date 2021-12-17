package main

import (
	"fmt"
)

func main() {
	var isCheck = false // 省略变量类型，根据推断变量类型
	var name string     //省略表达式，使用字符串对应的零值初始化（空字符串）
	var age int         //省略表达式，使用数值类型对应的零值初始化（0）
	fmt.Println(isCheck)
	fmt.Println(name)
	fmt.Println(age)

	// 简单声明变量方式
	name2 := "lizimeng"
	fmt.Println(name2)
	name3 := name2 + "hello"
	fmt.Println(name3)

	"const"关键字可以出现在任何"var"关键字出现的地方
	区别是常量必须有初始值
	const n = 20
	// 常量表达式可以执行任意精度数学计算
	const d = 10000 / n
	fmt.Println(d)

	// 数值型常量没有具体类型，除非指定一个类型
	// 比如显式类型转换
	fmt.Println(int64(d))

	// 数值型常量可以在程序的逻辑上下文中获取类型，比如变量赋值或者函数调用，例如，对于math包中的Sin函数,它需要一个float64类型的变量
	fmt.Println(math.Sin(n))

	// unsafe.Sizeof(n1)是unsafe包的函数，可以返回变量占用的字节数
	var a int8 = 120
	fmt.Printf("%T\n", a)
	fmt.Println(unsafe.Sizeof(a))

	// int不同长度转换
	var num1 int8
	num1 = 127
	num2 := int32(num1)
	fmt.Printf("值:%v 类型%T", num2, num2)
}
