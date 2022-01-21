package main

import "fmt"

// Go 语言切片是对数组的抽象。 Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
// Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
// 切片是一种方便、灵活且强大的包装器。切片本身没有任何数据。它们只是对现有数组的引用。
// 切片与数组相比，不需要设定长度，在[]中不用设定值，相对来说比较自由
// 从概念上面来说slice像一个结构体，这个结构体包含了三个元素：
// 指针，指向数组中slice指定的开始位置
// 长度，即slice的长度
// 最大长度，也就是slice开始位置到数组的最后位置的长度
func main() {
	//定义初始化
	s := []int{1, 2, 3}
	fmt.Println(s[0])

	//len() 方法获取长度
	fmt.Println(len(s))
	//cap() 测量切片最长可以达到多少
	fmt.Println(cap(s))

	//打印s指针地址
	fmt.Printf("slice addr %p", s)

	//添加元素
	s = append(s, 4)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	//打印s指针地址
	fmt.Printf("slice addr %p", s)

}
