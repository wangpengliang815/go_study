// @title 《Go语言编程》-复合数据类型map
// @description
// @author wangpengliang
// @date 2022-03-25 11:17:08
package main

import (
	"fmt"
	"sort"
)

// 使用make创建map
func createMap() {
	users := make(map[string]int, 8)
	users["小王"] = 18
	users["小李"] = 2
	fmt.Println(users)
	fmt.Println(users["小王"])
	fmt.Printf("type of a:%T\n", users)
}

// 声明时候填充map
func createMap2() {
	users := map[string]int{
		"小周": 22,
		"小张": 23,
	}
	fmt.Println(users)
	fmt.Println(users["小王"])
	fmt.Printf("type of a:%T\n", users)
}

// 判断键是否存在：如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
func existMap() {
	users := map[string]int{
		"小周": 22,
		"小张": 23,
	}
	// 如果返回值是bool值通常使用ok接收，约定俗称
	v, ok := users["小张"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}
}

// 遍历map
func traversalMap() {
	users := map[string]int{
		"小周": 22,
		"小张": 23,
	}
	for k, v := range users {
		fmt.Println(k, v)
	}

	// 只想遍历key
	for k := range users {
		fmt.Println(k)
	}
}

// 顺序遍历map
func orderTraversalMap() {
	users := map[string]int{
		"1": 22,
		"2": 23,
		"3": 23,
		"4": 23,
		"5": 23,
		"6": 23,
		"7": 23,
		"8": 23,
	}
	// 这里看到map的迭代顺序是不确定随机的，并且不同的哈希函数实现可能导致不同的遍历顺序
	for k, v := range users {
		fmt.Println(k, v)
	}
	fmt.Println("************排序后：***********************")
	// 如果需要按顺序遍历 key/value，必须显式地对key进行排序
	// 1.取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range users {
		keys = append(keys, key)
	}
	// 2.对切片进行排序
	sort.Strings(keys)
	// 3.按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, users[key])
	}
}

// delete删除某个键值对，key不存在时什么也不做
func delMap() {
	users := map[string]int{
		"小周": 21,
		"小张": 22,
	}
	for k, v := range users {
		fmt.Println(k, v)
	}
	delete(users, "小周")
	for k, v := range users {
		fmt.Println(k, v)
	}
}

// 元素为map类型的切片
func sliceValueMap() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "wangpengliang"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "北京"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

// 值为切片类型的map
func mapValueslice() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
