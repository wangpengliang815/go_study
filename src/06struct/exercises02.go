package main

import "fmt"

type Test struct {
	name   string
	age    int8
	dreams []string
}

func (t *Test) setDreams(dreams []string) {
	t.dreams = dreams
}

func exercises02Test() {
	t := Test{name: "wangpengliang", age: 18}
	data := []string{"吃饭", "睡觉", "搞钱"}
	t.setDreams(data)

	// 你真的想要修改t.dreams吗？
	data[1] = "不睡觉"
	fmt.Println(t.dreams) // [吃饭 不睡觉 搞钱]
}

// 使用传入的slice的拷贝进行结构体赋值,避免修改了 `slice` 导致结构体内容也被修改
func (t *Test) setDreams1(dreams []string) {
	t.dreams = make([]string, len(dreams))
	copy(t.dreams, dreams)
}

func exercises03Test() {
	t := Test{name: "wangpengliang", age: 18}
	data := []string{"吃饭", "睡觉", "搞钱"}
	t.setDreams1(data)

	// 你真的想要修改t.dreams吗？
	data[1] = "不睡觉"
	fmt.Println(t.dreams) // [吃饭 睡觉 搞钱]
}
