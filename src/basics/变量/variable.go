package main

import (
	"fmt"
)

// 全局声明变量,如果没有使用不会编译报错
var test1 string
var test2 string

var (
	author  string
	address string
)

func main() {
	variableTest()
}

// 变量
func variableTest() {
	// 变量定义方式1：指定变量类型但不赋值，使用默认值
	var name string // 省略表达式，使用字符串对应的零值初始化（空字符串）
	var age int     // 省略表达式，使用数值类型对应的零值初始化（0）
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("**************************************************")

	// 变量定义方式2：使用类型推导
	var address = "北京"
	fmt.Println(address)
	fmt.Println("**************************************************")

	// 变量定义方式3：使用简短声明，只能在函数内使用
	sex := "男"
	fmt.Println(sex)
	fmt.Println("**************************************************")

	// 变量声明方式4：使用 new()函数
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*p) // "0"
	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"
	fmt.Println("**************************************************")

	// 批量定义变量方式1：需要注意如果定义的是局部变量，则必须使用。否则编译报错
	var (
		name1 string
		age1  int
	)
	fmt.Println(name1)
	fmt.Println(age1)
	fmt.Println("**************************************************")

	// 批量定义变量方式2
	var n1, n2, n3 string
	n1, n2, n3 = "1", "2", "3"
	fmt.Println(n1, n2, n3)
	// 或者
	var n4, n5, n6 = "1", "2", "3"
	fmt.Println(n4, n5, n6)
	fmt.Println("**************************************************")

	// 指针
	x := 1           // 声明int类型变量x
	p1 := &x         // &x用于获取变量x的内存地址，返回一个指向x的指针p
	fmt.Println(*p1) // *p用户获取指针p指向变量的值
	*p1 = 2          // 可以重新给*p指针赋值
	fmt.Println(x)   // "2"
	fmt.Println("**************************************************")

	// 匿名变量，_ 代表匿名变量,匿名变量将会被忽略
	result, _ := sayHello()
	fmt.Println(result)
}

func sayHello() (string, int) {
	return "hello", 20
}
