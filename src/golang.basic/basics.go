package main

import (
	"bufio"
	"fmt"
	"os"
)

// 全局声明变量,如果没有使用不会编译报错
var (
	author  string
	address string
)

func main() {
	variableTest()
}

// 变量
func variableTest() {
	// 变量声明第一种：指定变量类型但不赋值，使用默认值
	var name string // 省略表达式，使用字符串对应的零值初始化（空字符串）
	var age int     // 省略表达式，使用数值类型对应的零值初始化（0）
	fmt.Println(name, age)

	// 批量定义变量,需要注意如果定义的是局部变量，则必须使用。否则编译报错
	var (
		name1 string
		age1  int
	)
	fmt.Println(name1, age1)

	// 变量声明第二种方式：类型推导
	var address = "beijing"
	fmt.Println(address)

	// 变量声明第三种方式：简短声明，只能在函数内使用
	sex := "男"
	fmt.Println(sex)

	// 多变量声明
	var n1, n2, n3 string
	n1, n2, n3 = "1", "2", "3"
	fmt.Println(n1, n2, n3)
	// 或者直接赋值
	var n4, n5, n6 = "1", "2", "3"
	fmt.Println(n4, n5, n6)

	// 变量声明第四种方式：使用 new()函数
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
