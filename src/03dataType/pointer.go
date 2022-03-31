// @title 《Go语言编程》-复合数据类型指针
// @description
// @author wangpengliang
// @date 2022-03-25 11:35:45
package main

import "fmt"

func pointer() {
	// 指针的&和*
	x := 1          // 声明int类型变量x
	p := &x         // &x用于获取变量x的内存地址，返回一个指向x的指针p
	fmt.Println(*p) // *p用户获取指针p指向变量的值
	*p = 2          // 可以重新给*p指针赋值
	fmt.Println(x)  // "2"

	// 引用类型不仅要声明，而且需要分配内存空间
	// var a *int
	// *a = 100
	// fmt.Println(*a)

	// var b map[string]int
	// b["wangpengliang"] = 100
	// fmt.Println(b)

	// new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false

	// make函数只用于 slice、map以及 channel 的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型
	var c map[string]int = make(map[string]int, 10)
	c["山西一枝花"] = 100
	fmt.Println(c)
}
