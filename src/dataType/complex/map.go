package main

import "fmt"

// map
func mapTest() {
	// 正常创建map
	ageMap := make(map[string]int, 8)
	ageMap["小王"] = 18
	ageMap["小李"] = 2
	fmt.Println(ageMap)
	fmt.Println(ageMap["小王"])
	fmt.Printf("type of a:%T\n", ageMap)

	// 声明时候填充map
	ageMap2 := map[string]int{
		"小周": 22,
		"小张": 23,
	}
	fmt.Println(ageMap2)
	fmt.Println(ageMap2["小王"])
	fmt.Printf("type of a:%T\n", ageMap2)

	// 判断键是否存在：如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	// 如果返回值是bool值通常使用ok接收，约定俗称
	v, ok := ageMap2["小张"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}

	// 遍历map
	for k, v := range ageMap2 {
		fmt.Println(k, v)
	}
	// 只想遍历key
	for k := range ageMap2 {
		fmt.Println(k)
	}

	// delete删除某个键值对，不存在时什么也不做
	ageMap3 := map[string]int{
		"小周": 21,
		"小张": 22,
	}
	for k, v := range ageMap3 {
		fmt.Println(k, v)
	}
	delete(ageMap3, "小周")
	for k, v := range ageMap3 {
		fmt.Println(k, v)
	}
}
