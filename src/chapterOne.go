package main

import (
	"fmt"
	"math"
	"runtime"
	"unsafe"
)

func main() {
	variableTest()

	constTest()

	// 返回变量占用的字节数
	var value int8 = 120
	fmt.Printf("%T\n", value)
	fmt.Println(unsafe.Sizeof(value))

	// int不同长度转换
	var num1 int8 = 127
	num2 := int32(num1)
	fmt.Printf("值:%v 类型%T", num2, num2)

	numberLiteralsSyntaxTest()

	convertOutputTest()

	var s1 string = "wangpengliang"
	var s2 string = `wangpengliang`
	fmt.Println(s1, s2)
}

// @title variableTest
// @description   变量
// @auth     WangPengLiang
func variableTest() {
	printRunFuncName()
	var isCheck = false // 省略变量类型，根据value推断变量类型
	var name string     //省略表达式，使用字符串对应的零值初始化（空字符串）
	var age int         //省略表达式，使用数值类型对应的零值初始化（0）
	fmt.Println(isCheck, name, age)

	// 简单声明变量方式
	name2 := "lizimeng"
	fmt.Println(name2)
	// 字符串拼接
	name3 := name2 + "hello"
	fmt.Println(name3)
}

// @title constTest
// @description   常量
// @auth     WangPengLiang
func constTest() {
	printRunFuncName()
	const n = 20        //"const"关键字可以出现在任何"var"关键字出现的地方,区别是常量必须有初始值
	const d = 10000 / n // 常量表达式可以执行任意精度数学计算
	fmt.Println(d)
	fmt.Println(int64(d))    // 数值型常量没有具体类型，除非指定一个类型，比如显式类型转换
	fmt.Println(math.Sin(n)) // 数值型常量可以在程序的逻辑上下文中获取类型，比如变量赋值或者函数调用，例如，对于math包中的Sin函数,它需要一个float64类型的变量
}

// @title numberLiteralsSyntaxTest
// @description   数字字面量语法
// @auth     WangPengLiang
func numberLiteralsSyntaxTest() {
	printRunFuncName()
	v1 := 0b00101101 //代表二进制的 101101，相当于十进制的 45
	fmt.Printf("值:%v 类型%T \n", v1, v1)
	v2 := 0o377 //代表八进制的377，相当于十进制的 255
	fmt.Printf("值:%v 类型%T \n", v2, v2)
	v3 := 0x1p-2 //代表十六进制的 1 除以 2²，也就是 0.25
	fmt.Printf("值:%v 类型%T \n", v3, v3)
	v4 := 123_456 // 使用“_”分隔数字
	fmt.Printf("值:%v 类型%T \n", v4, v4)
}

// @title convertOutputTest
// @description   以不同进制显示
// @auth     WangPengLiang
func convertOutputTest() {
	printRunFuncName()
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a) // 占位符%b表示二进制

	// 八进制以0开头
	var b int = 077

	fmt.Printf("%o \n", b)

	// 十六进制 以 0x 开头
	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)
	fmt.Printf("%d \n", c)
}

// 打印当前运行函数名称
func printRunFuncName() {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	fmt.Println("FuncName:", f.Name()+"========================================================>")
}
