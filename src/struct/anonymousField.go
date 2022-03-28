// @title 《Go语言编程》- 结构体匿名字段
// @description
// @author wangpengliang
// @date 2022-03-28 11:49:35

package main

import "fmt"

// 匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个
type Book struct {
	string
	float64
}

func anonymousFieldTest() {
	book := Book{
		"go语言编程",
		100.00,
	}
	fmt.Printf("%#v\n", book)              // main.Book{string:"go语言编程", float64:100}
	fmt.Println(book.string, book.float64) //北京 go语言编程 100
}
