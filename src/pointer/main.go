package main

import (
	"fmt"
)

/*指针是存储另一个变量的内存地址的变量。
我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址。*/
func main() {
	a := 10
	fmt.Println(a)
	change(&a)
	fmt.Println(a)
}

//使用指针传递函数的参数
func change(val *int) {
	*val = 58
}
