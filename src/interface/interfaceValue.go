// @title 《Go语言编程》-接口值、接口类型断言
// @description
// @author wangpengliang
// @date 2022-03-30 16:06:51

package main

import "fmt"

type Mover2 interface {
	Move()
}

type Dog2 struct {
	Name string
}

func (d *Dog2) Move() {
	fmt.Println("狗在跑~")
}

type Car2 struct {
	Brand string
}

func (c *Car2) Move() {
	fmt.Println("汽车在跑~")
}

func interfaceValueTest1() {
	var m Mover2
	// 此时，接口变量m是接口类型的零值,也就是它的类型和值部分都是nil
	fmt.Println(m == nil)

	// 不能对一个空接口值调用任何方法,否则会产生panic
	// m.Move() // panic: runtime error: invalid memory address or nil pointer dereference

	// 将一个*Dog结构体指针赋值给变量m,此时m的动态类型会被设置为*Dog，动态值为结构体变量的拷贝
	m = &Dog2{Name: "旺财"}

	// 给接口变量m赋值为一个*Car类型的值,此时接口值的动态类型为*Car，动态值为nil
	m = new(Car2)

	// 此时接口变量m与nil并不相等，因为它只是动态值的部分为nil，而动态类型部分保存着对应值的类型
	fmt.Println(m == nil)

	// 接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等时才相等
	var (
		x Mover2 = new(Dog2)
		y Mover2 = new(Car2)
	)
	fmt.Println(x == y) // false

	var z interface{} = []int{1, 2, 3}
	fmt.Println(z == z) // panic: runtime error: comparing uncomparable type []int
}

// 单个类型断言
func interfaceValueTest2() {
	var n Mover2 = &Dog2{Name: "旺财"}
	v, ok := n.(*Dog2)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "富贵" // 变量v是*Dog类型
	} else {
		fmt.Println("类型断言失败")
	}
}

// 对传入的空接口类型变量x进行类型断言
func interfaceValueTest3() {
	multipleInterfaceValueAssert(1)
}

func multipleInterfaceValueAssert(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
