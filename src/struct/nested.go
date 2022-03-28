// @title 《Go语言编程》-嵌套结构体、匿名字段结构体、结构体字段冲突
// @description
// @author wangpengliang
// @date 2022-03-28 13:41:04

package main

import "fmt"

// 地址
type Address struct {
	province   string
	city       string
	createTime string
}

// 邮箱
type Email struct {
	account    string
	createTime string
}

// 用户
type User struct {
	name    string
	gender  string
	address Address // 该字段为具名字段
	Address         // 嵌套的Address结构体也可以采用匿名字段的方式
	Email
}

func nestedStructTest() {
	// 具名字段赋值
	user := User{
		name:   "wangpengliang",
		gender: "男",
		address: Address{
			province: "山西",
			city:     "长治",
		},
	}
	fmt.Printf("user=%#v\n", user)

	// 匿名字段
	user1 := User{
		name:   "wangpengliang",
		gender: "男",
		Address: Address{
			province: "山西",
			city:     "长治",
		},
	}
	fmt.Printf("user1=%#v\n", user1)

	var user2 User
	user2.name = "wangpengliang"
	user2.gender = "男"
	// user2.createTime = "2019" //ambiguous selector user2.createTime
	user2.Address.createTime = "2000" //指定Address结构体中的createTime
	user2.Email.createTime = "2000"   //指定Email结构体中的createTime
	fmt.Printf("user2=%#v\n", user2)

	/*
		user=main.User{name:"wangpengliang", gender:"男", address:main.Address{province:"山西", city:"长治", createTime:""}, Address:main.Address{province:"", city:"", createTime:""}, Email:main.Email{account:"", createTime:""}}
		user1=main.User{name:"wangpengliang", gender:"男", address:main.Address{province:"", city:"", createTime:""}, Address:main.Address{province:"山西", city:"长治", createTime:""}, Email:main.Email{account:"", createTime:""}}
		user2=main.User{name:"wangpengliang", gender:"男", address:main.Address{province:"", city:"", createTime:""}, Address:main.Address{province:"", city:"", createTime:"2000"}, Email:main.Email{account:"", createTime:"2000"}}
	*/
}
