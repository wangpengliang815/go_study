package main

import "fmt"

type student struct {
	name string
	age  int64
}

func main() {
	m := make(map[string]*student)

	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		fmt.Println(&stu)
		fmt.Println(*&stu.age, *&stu.name)

		m[stu.name] = &stu
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
