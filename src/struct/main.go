package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
}

func main() {
	// 结构体初始化
	var p person
	p.name = "wangpengliang"
	p.age = 18
	p.address = "北京"
	fmt.Printf("p=%v \n", p)
	fmt.Printf("p=%#v \n", p)

	// 匿名结构体
	var user struct {
		name string
		age  int
	}
	user.name = "wangpengliang"
	user.age = 18
	fmt.Printf("%v \n", user)

	// 创建指针类型结构体
	var p1 = new(person)
	fmt.Println(p1.address)
	fmt.Printf("%T\n", p1)
	fmt.Printf("p1=%#v\n", p1)
}
