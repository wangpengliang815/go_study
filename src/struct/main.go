// @title 《Go语言编程》-结构体
// @description
// @author wangpengliang
// @date 2022-03-25 11:59:36
package main

import (
	"fmt"
	"unsafe"
)

// func main() {
// 	structEmpty()
// }

// 定义表示矩形的结构体
type Rect struct {
	width, height float64
}

// 最基础的结构实例化方式
func structInstantiation1() {
	var rect Rect
	rect.height = 13.4
	rect.width = 21.4
	fmt.Printf("%T \n", rect)
	fmt.Println(rect.height, rect.width)
}

// 使用new关键字对结构体进行实例化,得到的是结构体的地址
func structInstantiation2() {
	var rect = new(Rect)
	rect.height = 13.4
	rect.width = 21.4
	fmt.Printf("%T \n", rect)
	fmt.Println(rect.height, rect.width)
}

// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
func structInstantiation3() {
	var rect = &Rect{}
	rect.height = 13.4
	rect.width = 21.4
	fmt.Printf("%T \n", rect)
	fmt.Println(rect.height, rect.width)
}

// 没有初始化的结构体，其成员变量都是对应其类型的零值
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

/*
使用值的列表初始化.初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值
使用这种格式初始化时，需要注意：
 1. 必须初始化结构体的所有字段。
 2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
 3. 该方式不能和键值初始化方式混用。
*/
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

type test struct {
	a int8
	b int8
	c int8
	d int8
}

// struct内存分配
func structMemoryAllocation() {
	rect := Rect{30.2, 43.5}
	fmt.Printf("rect.width %p \n", &rect.width)
	fmt.Printf("rect.heigh %p \n", &rect.height)

	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
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

// func main() {
// r := &Foo{}
// r.Bar()
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
// }
