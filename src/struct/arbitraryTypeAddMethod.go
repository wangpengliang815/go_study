// @title 《Go语言编程》- 给任意类型添加方法
// @description
// @author wangpengliang
// @date 2022-03-28 11:49:35

package main

import "fmt"

// 将int定义为自定义类型
type MyInt int

// SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) sayHello() {
	fmt.Println("Hello, 我是一个int。")
}

func arbitraryTypeAddMethodTest() {
	var value MyInt
	value.sayHello()
	value = 100
	fmt.Printf("%#v  %T\n", value, value) //100  main.MyInt
}
