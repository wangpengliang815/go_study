package main

import "fmt"

type books struct {
	title   string
	author  string
	book_id int
}

//   语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
//   结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
func main() {
	b := books{"标题", "标题作者", 1}
	fmt.Println(b.title, b.author, b.book_id)

}
