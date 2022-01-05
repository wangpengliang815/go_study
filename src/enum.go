package main

import (
	"fmt"
)

// Go没有枚举,使用const+iota模拟
func EnumTest() {
	// 普通常量组
	const (
		Windows = 0
		Linux
		MaxOS
	)
	// 普通常量组如果不指定类型和初始化值，则与上一行非空常量右值相同
	fmt.Println(Windows, Linux, MaxOS)

	/*
	 结合iota实现枚举
	 第一个iota等于0，每当 iota 在新的一行被使用时，它的值都会自动加1
	*/
	const (
		Sunday    = iota // 0
		Monday           // 1,通常省略后续行行表达式。
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	// 插队
	const (
		a1 = iota    // 0
		a2           // 1
		b1 = "hello" // 独立值hello,iota+=1
		b2           // 如不指定类型和初始化值，则与上一行非空常量右值相同,所以是hello;iota+=1
		c1 = iota    // 恢复自增,4
		c2           // 5
	)
	fmt.Println(a1, a2, b1, b2, c1, c2)

	// 使用_忽略某些值
	const (
		d1 = iota
		d2
		_
		d3
	)
	fmt.Println(d1, d2, d3)

	// 定义数量级别
	const (
		_  = iota             // _表示将0忽略
		KB = 1 << (10 * iota) // 表示1左移十位，转换为十进制就是1024
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	// 多个iota一行定义
	const (
		a, b = iota + 1, iota + 2 // a=iota+1 b=iota+2 => 1,2
		c, d                      // c=iota(1)+1 b=iota(2)+1 => 2,3
	)
	fmt.Println(a, b, c, d)
}
