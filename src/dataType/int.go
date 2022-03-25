// @title 《Go语言编程》-基础类型整型
// @description  包括：数字字面量语法、unsafe.Sizeof()
// @author wangpengliang
// @date 2022-03-25 10:54:55
package main

import (
	"fmt"
	"unsafe"
)

// 整型：数字字面量语法允许以不同进制定义整形变量
func numberLiteralsSyntax() {
	// 代表二进制的 101101，相当于十进制的 45
	v1 := 0b00101101
	fmt.Printf("value:%v type:%T \n", v1, v1)

	// 代表八进制的377，相当于十进制的 255
	v2 := 0o377
	fmt.Printf("value:%v type:%T \n", v2, v2)

	// 代表十六进制的 1 除以 2²，也就是 0.25
	v3 := 0x1p-2
	fmt.Printf("value:%v type:%T \n", v3, v3)

	// 使用“_”分隔数字
	v4 := 123_456
	fmt.Printf("value:%v type:%T \n", v4, v4)
}

// 返回变量占用的字节数
func getSizeof() {
	var value1 int8 = 120  // 1
	var value2 int16 = 120 // 2
	var value3 int32 = 120 // 4
	var value4 int64 = 120 // 8
	fmt.Println(unsafe.Sizeof(value1), unsafe.Sizeof(value2), unsafe.Sizeof(value3), unsafe.Sizeof(value4))

	var name string = "王"
	fmt.Println(unsafe.Sizeof(name))

	var name1 string = "王王王王王王王王王王王王王王王"
	fmt.Println(unsafe.Sizeof(name1))
}

// 借助fmt函数不同进制形式显示整数
func convertOutput() {
	// 十进制
	var a int = 10
	fmt.Printf("十进制:%d 二进制:%b \n", a, a) // %d表示十进制,%b表示二进制

	// 八进制以0开头
	var b int = 077
	fmt.Printf("八进制:%o 二进制:%d\n", b, b) // %o表示八进制,%d表示十进制

	// 十六进制 以 0x 开头
	var c int = 0xff
	fmt.Printf("十六进制:%x 十六进制大写:%X 十进制:%d 变量类型:%T \n", c, c, c, c) //%x 表示十六进制,%X 表示大写的十六进制,%d 表示十进制,%T 查看变量类型

	// int不同长度转换
	var num1 int8 = 127
	num2 := int32(num1)
	fmt.Printf("value:%v type:%T \n", num2, num2)
}
