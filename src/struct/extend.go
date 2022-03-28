// @title 《Go语言编程》-结构体继承
// @description
// @author wangpengliang
// @date 2022-03-28 14:03:11

package main

import "fmt"

// Animal
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动 \n", a.name)
}

func (a *Animal) eat() {
	fmt.Printf("%s会吃 \n", a.name)
}

// Dog
type Dog struct {
	Feet    int8
	*Animal // 通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s叫声: 汪汪汪~\n", d.name)
}

func (d *Dog) eat() {
	fmt.Printf("%s吃骨头\n", d.name)
}

// Cat
type Cat struct {
	Feet    int8
	*Animal // 通过嵌套匿名结构体实现继承
}

func (d *Cat) miao() {
	fmt.Printf("%s叫声: 喵喵喵~\n", d.name)
}

func (d *Cat) eat() {
	fmt.Printf("%s吃鱼\n", d.name)
}

func extendTest() {
	dog := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	dog.wang()
	dog.eat()
}
