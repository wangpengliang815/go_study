package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

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

	// 变量声明第四种方式：使用new()函数
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*p) // "0"
	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"

	// 指针
	x := 1           // 声明int类型变量x
	p1 := &x         // &x用于获取变量x的内存地址，返回一个指向x的指针p
	fmt.Println(*p1) // *p用户获取指针p指向变量的值
	*p1 = 2          // 可以重新给*p指针赋值
	fmt.Println(x)   // "2"
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
