package main

import "fmt"

func exercises01Test() {
	make := make(map[int]*int)
	slice := []int{0, 1, 2, 3}
	for index, value := range slice {
		make[index] = &value
	}

	for k, v := range make {
		fmt.Println(k, "=>", *v)
	}
	/*
		思考：这里为什么都是3
			1 => 3
			2 => 3
			3 => 3
			0 => 3
	*/
}

// 第一种方式处理
func solveTest1() {
	make := make(map[int]*int)
	slice := []int{0, 1, 2, 3}
	for index, value := range slice {
		// 循环中使用新的变量接收
		item := value
		fmt.Printf("%p \n", &item)
		make[index] = &item
	}
	for k, v := range make {
		fmt.Println(k, "=>", *v)
	}
}

// 第二种方式处理
func solveTest2() {
	make := make(map[int]*int)
	slice := []int{0, 1, 2, 3}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%p \n", &slice[i])
		make[i] = &slice[i]
	}
	for k, v := range make {
		fmt.Println(k, "=>", *v)
	}
}
