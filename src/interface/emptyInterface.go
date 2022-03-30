// @title 《Go语言编程》-空接口
// @description
// @author wangpengliang
// @date 2022-03-30 09:58:29

package main

import "fmt"

// Any 不包含任何方法的空接口类型
type Any interface{}

type Person struct{}

// 空接口作为函数参数
func print(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 空接口类型的变量可以存储任意类型的值
func emptyInterfaceTest() {
	// 声明空接口类型变量
	var x Any

	x = "你好" // 字符串型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = 100 // int型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = true // 布尔型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = Person{} // 结构体类型
	fmt.Printf("type:%T value:%v\n", x, x)
}

func emptyInterfaceTest2() {
	// 不声明空接口类型直接使用即可
	var x interface{}

	x = "你好" // 字符串型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = 100 // int型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = true // 布尔型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = Person{} // 结构体类型
	fmt.Printf("type:%T value:%v\n", x, x)
}

func emptyInterfaceTest3() {
	print(100)
	print("100")
}

// 空接口作为map的值可以保存任意值的字典
func emptyInterfaceTest4() {
	var person = make(map[string]interface{})
	person["name"] = "wangpengliang"
	person["age"] = 18
	person["married"] = false
	fmt.Println(person)
}
