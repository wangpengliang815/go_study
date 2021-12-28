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
	processControlContinue()
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

// if
func processControlIf() {
	score := 60
	if score >= 90 {
		fmt.Println("A")
	} else if score > 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// if 特殊写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断;这种写法可以将返回值与判断放在一行进行处理，而且返回值的作用范围被限制在 if、else 语句组合中。
	// 在编程中，变量的作用范围越小，所造成的问题可能性越小，每一个变量代表一个状态，有状态的地方，状态就会被修改，函数的局部变量只会影响一个函数的执行，但全局变量可能会影响所有代码的执行状态，因此限制变量的作用范围对代码的稳定性有很大的帮助
	if score := 60; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// for
func processControlFor() {
	// 基础的for循环
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// 初始语句可以被忽略，但是初始语句后的分号必须要写
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// 循环的初始语句和结束语句都可以省略:类似于C#中的while,在`while`后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 无限循环
	for {
		fmt.Print("hello word")
	}
}

// switch
func processControlSwitch() {
	// 基础的switch
	sex := "男"
	switch sex {
	case "男":
		fmt.Println("男性")
	case "女":
		fmt.Println("女性")
	default:
		fmt.Println("无效的输入！")
	}

	// switch多个值判断也使用特殊写法
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	// switch 使用表达式
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	// fallthrough 语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

// goto
func processControlGoto() {
	// 示例一：内层循环打印到2时结束，外层循环也随即结束
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
	// 使用goto简化
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

// break
func processControlBreak() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1 // 退出到定义的标签位置
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

// continue
func processControlContinue() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}
