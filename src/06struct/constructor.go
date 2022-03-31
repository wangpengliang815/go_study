// @title 《Go语言编程》-结构体的构造函数、方法和接受者(Receiver)
// @description
// @author wangpengliang
// @date 2022-03-28 10:24:45

package main

import "fmt"

// 如果不初始化,字段值为类型零值
type Person struct {
	name    string
	age     int
	address string
	hobby   []string
}

// 构造函数重载第一种方式,使用指定值初始化结构体,返回指针类型
func NewPerson(name, address string, age int, hobby []string) *Person {
	instance := new(Person)
	instance.name = name
	instance.address = address
	instance.age = age
	instance.hobby = hobby
	return instance
}

// 构造函数重载第二种方式,使用&初始化结构体,返回指针类型(Go中不支持函数重载,所以两个构造函数名称必须不同)
func NewPerson2(name, address string, age int) *Person {
	return &Person{
		name:    name,
		address: address,
		age:     age,
	}
}

// Person结构成员方法,所谓方法在go中就是定义了接受者的函数：使用值类型接受者
func (p Person) say1() {
	fmt.Printf("name: %s,age: %d  \n", p.name, p.age)
}

func (p Person) addAge1() {
	p.age = p.age + 1
}

// 使用指针类型接收者
func (p *Person) say2() {
	fmt.Printf("name: %s,age: %d  \n", p.name, p.age)
}

func (p *Person) addAge2() {
	p.age = p.age + 1
}

func constructorTest() {
	// 第一种情况：接受者是struct
	var p1 Person = Person{"zhansan", 16, "beijing", []string{}}
	p1.addAge1()
	p1.say1()

	var p2 *Person = &Person{"lisi", 16, "shanghai", []string{}}
	p2.addAge1()
	p2.say1()

	// 第二种情况：接受者是指针
	var p3 Person = Person{"zhansan", 16, "beijing", []string{}}
	p3.addAge2()
	p3.say2()

	var p4 *Person = &Person{"lisi", 16, "beijing", []string{}}
	p4.addAge2()
	p4.say2()

	// 使用不同构造函数初始化
	a := NewPerson("wangpengliang", "beijing", 18, []string{"java", "go"})
	fmt.Println(a)
	b := NewPerson2("wangpengliang", "beijing", 18)
	fmt.Println(b)
}
