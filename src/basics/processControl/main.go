package main

import "fmt"

func main() {
	continueTest()
}

// if
func ifTest() {
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
func forTest() {
	// 基础的for循环
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i // 1+2+3+4
	}
	fmt.Println(sum)

	// 初始语句可以被忽略，但是初始语句后的分号必须要写
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0
	// 循环的初始语句和结束语句都可以省略:类似于C#中的while,在`while`后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。类似于while循环
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// 死循环
	for {
		fmt.Print("hello word!")
	}
}

// switch
func switchTest() {
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

	// switch多个值判断使用特殊写法
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
func gotoTest() {
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

	// 示例二：使用goto简化
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
func breakTest() {
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
func continueTest() {
	// 示例一：continue跳到指定的标签位置
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

	// // 示例二：跳出当次循环，开始下次循环
	// for i := 0; i < 10; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
}
