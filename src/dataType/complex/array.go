package main

import "fmt"

// 数组
func arrayTest() {
	// 数组创建，使用类型零值初始化
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// 数组创建使用数字字面值语法
	var b [3]int = [3]int{1, 2, 3}
	fmt.Println(b[0])
	// fmt.Println(a[10]) invalid array index 10 (out of bounds for 3-element array),数组越界

	//如果数组的长度位置出现的是“…”省略号，表示数组的长度是根据初始化值的个数来计算
	c := [...]int{1, 2, 3}
	fmt.Println(c[0])

	// 指定数组第1个元素值为100,第9个元素值为200,数组长度就是10
	d := [...]int{1: 100, 9: 200}
	fmt.Println(d)

	// 数组遍历
	for i := 0; i < len(a); i++ {
		fmt.Printf("key:%d, value:%d\n", i, a[i])
	}

	// 使用range遍历数组
	for i, v := range a {
		fmt.Printf("key:%d, value:%d\n", i, v)
	}

	// 使用range遍历数组:如果需要值但希望忽略索引，可以通过使用_ blank标识符替换索引来实现
	for _, v := range a {
		fmt.Printf("value:%d\n", v)
	}

	// go同样支持多维数组
	array := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /*  第三行索引为 2 */
	}
	for i, item := range array {
		fmt.Printf("%d the element of is %v \n", i, item)
	}

	// 数组比较：数组可以直接进行比较，当数组内的元素都一样的时候表示两个数组相等。
	q := [...]int{11, 22, 33}
	w := [...]int{11, 22, 33}
	fmt.Println(q == w)
	fmt.Println(q == b)

	// 数组是值类型：意味着当它被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改不会在原始数组中反映
	e := [...]int{11, 22, 33}
	fmt.Println(e[0])
	arrayTest2(e)
	fmt.Println(e[0])

	// 使用指针后函数内部对数组的更改将反应到原数组上
	r := [...]int{11, 22, 33}
	fmt.Println(r[0])
	arrayTest3(&r)
	fmt.Println(r[0])
}

func arrayTest2(arr [3]int) {
	arr[0] = 1
}

func arrayTest3(arr *[3]int) {
	arr[0] = 100
}
