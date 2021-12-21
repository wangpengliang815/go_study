package main

import (
	"bufio"
	"fmt"
	"os"
	"unsafe"
)

func main() {
	// 枚举
	enumTest()
}

// 变量
func variableTest() {
	// 变量声明第一种：指定变量类型但不赋值，使用默认值
	var name string // 省略表达式，使用字符串对应的零值初始化（空字符串）
	var age int     // 省略表达式，使用数值类型对应的零值初始化（0）
	fmt.Println(name, age)

	// 变量声明第二种方式：类型推断
	var address = "beijing"
	fmt.Println(address)

	// 变量声明第三种方式：简短声明
	sex := "男"
	fmt.Println(sex)

	// 多变量声明
	var name1, name2, name3 string
	name1, name2, name3 = "1", "2", "3"
	fmt.Println(name1, name2, name3)
	// 或者直接赋值
	var name4, name5, name6 = "1", "2", "3"
	fmt.Println(name4, name5, name6)

	// 前者表示所见即所得的(除了回车符)。后者所表示的值中转义符会起作用
	var s1 string = "WangPengLiang"
	var s2 string = `WangPengLiang`
	fmt.Println(s1, s2)
}

// 常量
func constTest() {
	// 显式类型定义,常量名称推荐全大写
	const LENGTH int = 10
	// 隐式类型定义
	const WIDTH = 20
	// 多重定义
	const a, b, c = 1, false, "hello world"
	fmt.Println(a, b, c)

	area := LENGTH * WIDTH
	fmt.Println(area)
}

// Go没有枚举,使用const+iota模拟
func enumTest() {
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
	 第一个 iota 等于0，每当 iota 在新的一行被使用时，它的值都会自动加 1
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

	const (
		a1 = iota    // 0
		a2           // 1
		b1 = "hello" // 独立值hello,iota+=1
		b2           // 如不指定类型和初始化值，则与上一行非空常量右值相同,所以是hello;iota+=1
		c1 = iota    // 恢复自增,4
		c2           // 5
	)
	fmt.Println(a1, a2, b1, b2, c1, c2)

	// TODO: 关于iota的更多用法需要专门学习下
	const (
		A, B = iota, iota << 10 // 0, 0 << 10
		C, D                    // 1, 1 << 10
	)
	fmt.Println(A, B, C, D)
}

// 输入输出
/*
格式化打印占位符：
			%v,原样输出
			%T，打印类型
			%t,bool类型
			%s，字符串
			%f，浮点
			%d，10进制的整数
			%b，2进制的整数
			%o，8进制
			%x，%X，16进制
				%x：0-9，a-f
				%X：0-9，A-F
			%c，打印字符
			%p，打印地址
*/
func outputTest() {
	/*******************************************fmt包****************/
	// 普通输出
	fmt.Print("hello world")
	// 输出后换行
	fmt.Println("hello world2")
	// 格式化输出
	name := "WangPengLiang"
	fmt.Printf("type:%T value:%v", name, name)

	fmt.Println()
	fmt.Println("please input you name")
	var input string
	fmt.Scanln(&input) //读取键盘输入，通过操作地址赋值给input.阻塞式
	fmt.Println("you name is：" + input)

	/*******************************************bufio包：https://golang.google.cn/pkg/bufio/****************/
	fmt.Println("please  again input you name")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	fmt.Println("you name is：：", value)
}

// 数字字面量语法，以不同进制定义变量
func numberLiteralsSyntaxTest() {
	v1 := 0b00101101 //代表二进制的 101101，相当于十进制的 45
	fmt.Printf("value:%v type:%T \n", v1, v1)
	v2 := 0o377 //代表八进制的377，相当于十进制的 255
	fmt.Printf("value:%v type:%T \n", v2, v2)
	v3 := 0x1p-2 //代表十六进制的 1 除以 2²，也就是 0.25
	fmt.Printf("value:%v type:%T \n", v3, v3)
	v4 := 123_456 // 使用“_”分隔数字
	fmt.Printf("value:%v type:%T \n", v4, v4)
}

// 返回变量占用的字节数
func getSizeofTest() {
	var value int8 = 120
	fmt.Printf("%T\n", value)
	fmt.Println(unsafe.Sizeof(value))
}

// 借助 fmt 函数不同进制形式显示整数
func convertOutputTest() {
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

	// int不同长度转换
	var num1 int8 = 127
	num2 := int32(num1)
	fmt.Printf("value:%v type:%T \n", num2, num2)
}
