// @title 《Go语言编程》-常量
// @description  包括：常量定义、常量组定义、iota、枚举的实现
// @author wangpengliang
// @date 2022-03-25 10:45:45
package main

import (
	"fmt"
)

// 全局常量定义，一旦定义不能修改,如果没有使用不会编译报错
const PI = 3.141592653589793

func main() {
	constDeclar()
}

// 常量定义
func constDeclar() {
	const PI float64 = 3.14159265358979323846

	// 显式类型定义,常量名称推荐全大写
	const LENGTH int = 10

	// 隐式类型定义,其实也是使用类型推导
	const WIDTH = 20

	// 常量的多重定义
	const (
		size int64 = 1024
		eof        = -1 // 无类型整型常量
	)
	// 或者这样
	const a, b, c = 1, false, "hello world"
}

// Go没有枚举,需要使用const+iota模拟
func enum() {
	// 普通常量组
	const (
		Windows = 0
		Linux
		MaxOS
	)
	// 普通常量组如果不指定类型和初始化值，则与上一行非空常量右值相同
	fmt.Println(Windows, Linux, MaxOS)
	fmt.Println("**************************************************")

	// 结合iota实现枚举:第一个iota等于0，每当 iota 在新的一行被使用时，它的值都会自动加1
	const (
		Sunday    = iota // 0
		Monday           // 1,通常省略后续行表达式
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println("**************************************************")

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
	fmt.Println("**************************************************")

	// 使用_忽略某些值
	const (
		d1 = iota
		d2
		_
		d3
	)
	fmt.Println(d1, d2, d3)
	fmt.Println("**************************************************")

	// 定义数量级别
	const (
		_  = iota             // _表示将0忽略
		KB = 1 << (10 * iota) // 表示1左移十位，转换为十进制就是1024
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println("**************************************************")

	// 多个iota一行定义
	const (
		a, b = iota + 1, iota + 2 // a=iota+1 b=iota+2 => 1,2
		c, d                      // c=iota(1)+1 b=iota(2)+1 => 2,3
	)
	fmt.Println(a, b, c, d)
}
