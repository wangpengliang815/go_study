// @title 《Go语言编程》-接口定义、接口值接收者和指针接收者区别、一个类型实现多接口、多类型实现同一接口
// @description
// @author wangpengliang
// @date 2022-03-29 15:49:32

package main

import "fmt"

// Sayer 接口
type Sayer interface {
	Say()
}

// Mover 接口
type Mover interface {
	Move()
}

type Dog struct {
	name string
}

// Move 使用值接收者定义Move方法实现Mover接口
func (d Dog) Move() {
	fmt.Println("狗会动")
}

// 实现Sayer接口
func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

type Cat struct{}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}

// Car 汽车结构体类型
type Car struct {
	brand string
}

// Car类型实现Mover接口
func (c Car) Move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func dogTest() {
	var x Mover // 声明一个Mover类型的变量x

	var d1 = Dog{} // d1是Dog类型
	x = d1         // 可以将d1赋值给变量x
	x.Move()

	var d2 = &Dog{} // d2是Dog指针类型
	x = d2          // 也可以将d2赋值给变量x
	x.Move()
}

func catTest() {
	var x Mover

	// var c1 = Cat{} // c1是Cat类型
	// x = c1         // 编译错误：不能将c1当成Mover类型

	var c2 = &Cat{} // c2是*Cat类型
	x = c2          // 可以将c2当成Mover类型
	x.Move()
}

// 一个类型实现多个接口
func multiInterfaceTest() {
	var d = Dog{name: "旺财"}

	var s Sayer = d
	var m Mover = d

	s.Say()  // 对Sayer类型调用Say方法
	m.Move() // 对Mover类型调用Move方法}
}

// 多种类型实现同一接口
func multiTypeTest() {
	var x Mover

	x = Dog{name: "旺财"}
	x.Move()

	x = Car{brand: "宝马"}
	x.Move()
	/*
		狗会动
		宝马速度70迈
	*/
}
