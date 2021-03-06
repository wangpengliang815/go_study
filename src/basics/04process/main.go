// @title 《Go语言编程》-流程控制
// @description
// @author wangpengliang
// @date 2022-03-25 11:19:20
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
}

// if特殊语法：可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断;这种写法可以将返回值与判断放在一行进行处理，而且返回值的作用范围被限制在 if、else 语句组合中。
func specialIfTest() {
	// 这么做好处在于：在编程中，变量的作用范围越小，所造成的问题可能性越小，每一个变量代表一个状态，有状态的地方，状态就会被修改，函数的局部变量只会影响一个函数的执行，但全局变量可能会影响所有代码的执行状态，因此限制变量的作用范围对代码的稳定性有很大的帮助
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

// forRange用于遍历数组、切片、字符串、map 及通道（channel）
func forRangeTest() {
	c := [...]int{1, 2, 3}
	for _, v := range c {
		fmt.Printf("value:%d\n", v)
	}
}

// switch
func switchTest() {
	// 基础的switch
	sex := "男"
	switch sex {
	case "男":
		fmt.Println("男的")
	case "女":
		fmt.Println("女的")
	default:
		fmt.Println("无效的")
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
		fmt.Println("好好学习")
	case age > 25 && age < 35:
		fmt.Println("好好工作")
	case age > 60:
		fmt.Println("好好享受")
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

// TODO：需要结合实际场景再看看流程语句中的标记
// goto C#中也有只是不推荐使用,只在特定场景下才考虑使用。因为goto可以无条件地转移到过程中指定的行会造成程序流程的混乱，使理解和调试程序都产生困难
func gotoTest() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			// i==5时跳到指定的标签LOOP处
			goto LOOP
		} else {
			fmt.Printf("a: %d\n", i)
		}
	}
LOOP:
}

// break
func breakTest() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

// continue
func continueTest() {
	for i := 0; i < 3; i++ {
	forloop1:
		for j := 0; j < 3; j++ {
			if i == 2 {
				continue forloop1 // 如果加了标签表示开启标签对应的代码块
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
