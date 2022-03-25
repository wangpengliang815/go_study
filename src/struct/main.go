// @title 《Go语言编程》-结构体
// @description
// @author wangpengliang
// @date 2022-03-25 11:59:36
package main

import (
	"fmt"
)

// 定义结构体
type Rect struct {
	width, height float64
}

// 定义Rect成员方法
func (r *Rect) Area() float64 {
	return r.width * r.height
}

// 定义Rect构造函数
func NewRect(width, height float64) *Rect {
	return &Rect{width, height}
}

type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println("Base.Foo")
}
func (base *Base) Bar() {
	fmt.Println("Base.Bar")
}

type Foo struct {
	Base
}

func (foo *Foo) Bar() {
	foo.Base.Bar()
}

func main() {
	r := &Foo{}
	r.Bar()
	// var r1 Rect
	// r1.width = 3.8
	// r1.height = 2.8
	// result1 := r1.Area()
	// fmt.Println(result1)

	// r2 := new(Rect)
	// r2.width = 3.8
	// r2.height = 2.8
	// result2 := r2.Area()
	// fmt.Println(result2)

	// r3 := &Rect{}
	// r3.width = 3.8
	// r3.height = 2.8
	// result3 := r3.Area()
	// fmt.Println(result3)

	// r4 := &Rect{3.8, 2.8}
	// fmt.Println(r4.Area())

	// r5 := &Rect{width: 3.8, height: 2.8}
	// fmt.Println(r5.Area())

	// // 匿名结构体
	// var user struct {
	// 	name string
	// 	age  int
	// }
	// user.name = "wangpengliang"
	// user.age = 18
	// fmt.Printf("%v \n", user)

	// // 创建指针类型结构体
	// var p1 = new(person)
	// fmt.Println(p1.address)
	// fmt.Printf("%T\n", p1)
	// fmt.Printf("p1=%#v\n", p1)
}
