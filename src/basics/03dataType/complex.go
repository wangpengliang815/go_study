// @title 《Go语言编程》-基础数据类型复数(complex)
// @description
// @author wangpengliang
// @date 2022-03-25 10:58:06
package main

import "fmt"

// 复数:复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位,用的比较少不做详细记录
func complexDeclar() {
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}
