// @title 《Go语言编程》-变量
// @description  包括：变量声明、变量初始化、多变量声明、匿名变量、变量赋值
// @author wangpengliang
// @date 2022-03-25 10:45:00
package main

import (
	"fmt"
)

// 全局声明变量,如果没有使用不会编译报错
var t1 string
var t2 string
var (
	t3 string
	t4 string
)

func main() {
	varDeclar()
}

// 变量声明
func varDeclar() {
	// 指定变量类型但不赋值，使用默认值
	var v1 int      // 省略表达式，使用数值类型对应的零值初始化（0）
	var v2 string   // 省略表达式，使用字符串对应的零值初始化（空字符串）
	var v3 [10]int  // 会创建长度为10的数据
	var v4 []int    // 数组切片
	var v5 struct { // {0}
		f int
	}
	var v6 *int            // 指针,因未被初始化所以默认值为nil
	var v7 map[string]int  // map,key为string类型，value为int类型,默认值为map[]
	var v8 func(a int) int // 函数类型,因未被初始化所以默认值为nil,直接调用会抛出空指针异常
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8)
}

// 多变量声明
func varBatchDeclar() {
	// 多变量声明使用此方式避免重复写var
	var (
		v1 int
		v2 string
	)
	fmt.Println(v1, v2)

	// 或者这样
	var n1, n2, n3 string
	n1, n2, n3 = "1", "2", "3"
	fmt.Println(n1, n2, n3)

	// 再或者这样
	var n4, n5, n6 = "1", "2", "3"
	fmt.Println(n4, n5, n6)
}

// 变量初始化
func varInit() {
	var v1 int = 10 // 初始化方式1
	var v2 = 10     // 初始化方式2，使用类型推导。编译器可以自动推导出v2的类型
	v3 := 10        // 初始化方式3，使用简短声明。编译器可以自动推导出v3的类型.只能在函数内使用
	var v4 int      // 初始化方式4，使用类型对应零值初始化
	fmt.Println(v1, v2, v3, v4)
}

// 变量赋值
func varAssignment() {
	// 最简单的方式没什么好说的
	var value int
	value = 10
	fmt.Println(value)

	// 支持多重赋值,在不支持多重赋值的语言中,交换两个变量的内容需要引入一个中间变量
	var i int = 10
	var j int = 20
	i, j = j, i
	fmt.Println(i, j)
}

// 匿名变量测试方法
func getName() (firstName, lastName string) {
	return "wangpengliang", "lizimeng"
}

// 匿名变量
func varAnonymity() {
	// 匿名变量，_ 代表匿名变量,匿名变量将会被忽略
	result, _ := getName()
	fmt.Println(result)
}
