// @title 《Go语言编程》-结构体定义、初始化、匿名结构体
// @description
// @author wangpengliang
// @date 2022-03-27 19:17:07
package main

import (
	"fmt"
	"unsafe"
)

// Go中使用struct来实现面向对象,比如如下代码:定义表示矩形的结构体
type Rect struct {
	width, height float64
}

// struct最基础的结构实例化方式
func structInstance1() {
	var rect Rect
	fmt.Printf("%p %T \n", &rect, rect)
	rect.height = 13.4
	rect.width = 21.
	fmt.Printf("%p %T \n", &rect, rect)
	fmt.Println(rect.height, rect.width)
}

// 使用new关键字对struct进行实例化,得到的是结构体的地址
func structInstance2() {
	var rect = new(Rect)
	// Go语言中支持对结构体指针直接使用.来访问结构体的成员
	rect.height = 13.4
	rect.width = 21.4
	fmt.Printf("%T %p\n", rect, rect)
}

// 使用&对struct进行取址操作相当于对该结构体类型进行了一次new实例化操作
func structInstance3() {
	var rect = &Rect{}
	rect.height = 13.4
	rect.width = 21.4
	fmt.Printf("%T %p\n", rect, rect)
}

// 没有初始化的结构体，其成员变量都是其对应类型的零值
func structInit1() {
	var rect Rect
	fmt.Printf("%#v \n", rect)
}

// 使用键值对初始化.使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值
func structInit2() {
	rect := Rect{width: 10.5, height: 3.5}
	fmt.Printf("%#v \n", rect)

	// 对结构体指针进行键值初始化
	rect2 := &Rect{width: 10.5, height: 3.5}
	fmt.Printf("%#v \n", rect2)

	// 如果字段没有初始值,可以省略,那么被省略的值就是该字段的零值
	rect3 := Rect{width: 100.0}
	fmt.Printf("%#v \n", rect3)
}

//使用值的列表初始化.初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值.注意：必须初始化结构体的所有字段、初始值的填充顺序必须与字段在结构体中的声明顺序一致、该方式不能和键值初始化方式混用
func structInit3() {
	rect := Rect{10.5, 3.5}
	fmt.Printf("%#v \n", rect)

	rect2 := &Rect{10.5, 3.5}
	fmt.Printf("%#v \n", rect2)
}

// 空结构体不占用内存
func structEmpty() {
	var a struct{}
	fmt.Println(unsafe.Sizeof(a)) // 0
}

// 匿名结构体,常见于只用一次的场景
func structAnonymity() {
	// 匿名结构体
	var user struct {
		name string
		age  int
	}
	user.name = "wangpengliang"
	user.age = 18
	fmt.Printf("%v \n", user)
}
