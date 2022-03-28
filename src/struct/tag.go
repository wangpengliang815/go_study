// @title 《Go语言编程》-结构体标签
// @description
// @author wangpengliang
// @date 2022-03-28 15:43:08

package main

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Company struct {
	ID      int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Address string // json序列化是默认使用字段名作为key
	name    string // 私有不能被json包访问
}

func tagTest() {
	s1 := Company{
		ID:      1,
		Address: "男",
		name:    "沙河娜扎",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("%s\n", data) //json str:{"id":1,"Gender":"男"}
}
