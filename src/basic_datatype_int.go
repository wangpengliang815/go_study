package main

import (
	"fmt"
	"unsafe"
)

func main() {
	stringTest()
}

// 整型：数字字面量语法允许以不同进制定义整形变量
func numberLiteralsSyntaxTest() {
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
func getSizeofTest() {
	var value int8 = 120
	fmt.Printf("%T\n", value)
	fmt.Println(unsafe.Sizeof(value))
}

// 借助fmt函数不同进制形式显示整数
func convertOutputTest() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) //%d 表示十进制
	fmt.Printf("%b \n", a) //%b 表示二进制

	// 八进制以0开头
	var b int = 077
	fmt.Printf("%o \n", b) //%o 表示八进制
	fmt.Printf("%d \n", b) //%o 表示八进制

	// 十六进制 以 0x 开头
	var c int = 0xff
	fmt.Printf("%x \n", c) //%x 表示十六进制
	fmt.Printf("%X \n", c) //%X 表示大写的十六进制
	fmt.Printf("%d \n", c) //%d 表示十进制
	fmt.Printf("%T \n", c) //%T 查看变量类型

	// int不同长度转换
	var num1 int8 = 127
	num2 := int32(num1)
	fmt.Printf("value:%v type:%T \n", num2, num2)
}
