package main

import "fmt"

// 数组创建,使用类型零值初始化
func createArrayTest() {
	var arr [3]int               // array of 3 integers
	fmt.Println(arr[0])          // print the first element
	fmt.Println(arr[len(arr)-1]) // print the last element, a[2]
}

// 数组创建,使用数字字面值语法
func createArrayTest2() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr[0])
	// fmt.Println(a[10]) invalid array index 10 (out of bounds for 3-element array),数组越界
}

// 如果数组的长度位置出现的是“…”省略号，表示数组的长度是根据初始化值的个数来计算
func createArrayTest3() {
	arr := [...]int{1, 2, 3}
	fmt.Println(arr[0])
}

// 指定数组第1个元素值为100,第9个元素值为200,数组长度就是10
func createArrayTest4() {
	arr := [...]int{1: 100, 9: 200}
	fmt.Println(arr)
}

// 数组遍历
func foreachArrayTest() {
	var arr [3]int = [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("index:%d, value:%d\n", i, arr[i])
	}

	// 使用range遍历数组
	for i, v := range arr {
		fmt.Printf("index:%d, value:%d\n", i, v)
	}

	// 使用range遍历数组:如果需要值但希望忽略索引，可以通过使用_ blank标识符替换索引来实现
	for _, v := range arr {
		fmt.Printf("value:%d\n", v)
	}
}

// 多维数组
func multidimensionalArrayTest() {
	array := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /*  第三行索引为 2 */
	}
	for i, item := range array {
		fmt.Printf("%d the element of is %v \n", i, item)
	}
}

// 数组比较：数组可以直接进行比较，当数组内的元素都一样的时候表示两个数组相等
func compareArrayTest() {
	arr1 := [...]int{11, 22, 33}
	arr2 := [...]int{11, 22, 33}
	arr3 := [...]int{1, 2, 3}

	fmt.Println(arr1 == arr2)
	fmt.Println(arr1 == arr3)
}

// 数组
func arrayTest() {
	// 数组是值类型：意味着当它被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改不会在原始数组中反映
	e := [...]int{11, 22, 33}
	fmt.Println(e[0])
	passByValue(e)
	fmt.Println(e[0])

	// 使用指针后函数内部对数组的更改将反应到原数组上
	r := [...]int{11, 22, 33}
	fmt.Println(r[0])
	passByReference(&r)
	fmt.Println(r[0])
}

// 数组当参数传递,对数组进行了更改不会在原始数组中反映
func passByValue(arr [3]int) {
	arr[0] = 1
}

// 使用指针后函数内部对数组的更改将反应到原数组上
func passByReference(arr *[3]int) {
	arr[0] = 100
}
