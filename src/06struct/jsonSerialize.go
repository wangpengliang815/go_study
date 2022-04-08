package main

import (
	"encoding/json"
	"fmt"
)

// Student
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class
type Class struct {
	Title    string
	Students []*Student
}

func jsonSerializeTest() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 3; i++ {
		// 使用&方式初始化结构体
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		// 通过append方式向结构体指针类型切片中动态插入数据
		c.Students = append(c.Students, stu)
	}

	// JSON序列化：结构体=>SON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s %T\n", data, data)

	// JSON反序列化：JSON格式的字符串=>结构体
	c1 := &Class{}
	err = json.Unmarshal([]byte(string(data)), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
