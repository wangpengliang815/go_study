package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	stringTest()
}

// 字符串
func stringTest() {
	// 判断字符串中存在几个汉字
	name333 := "hello,勒布朗、科比、艾弗森"
	var count int
	for _, c := range name333 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)

	// 字符串字面量语法：可以定义一个多行字符串
	var s1 string = "hello"
	var s2 string = `hello
				     world`
	fmt.Println(s1, s2)

	// 内置函数len：返回字符串的字节数
	s := "abc北京"
	fmt.Printf("字节长度：%d \n", len(s)) // output:9
	// 返回字符串每个字节的值
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	// 获取字符串中字符的个数
	r := []rune(s)
	fmt.Println(len(r))

	// 字符串截取：
	str := "abcdefg"
	t1 := str[1:4] // startIndex=1,endIndex=4
	t2 := str[:4]  // 省略第一个索引从0开始，startIndex=0,endIndex=4
	t3 := str[1:]  // 省略第二个索引，从1开始截取到末尾
	// t4 := str[:10] //runtime error: slice bounds out of range [:10] with length 7 数组越界
	fmt.Println(t1, t2, t3)

	// 字符串拼接
	q1 := "hello"
	q2 := ",world"
	fmt.Println(q1 + q2)
	// %s 代表原样输出
	fmt.Printf("%s%s\n", q1, q2)
	// Sprintf拼接,不会直接打印而是生成一个新的字符串
	var result = fmt.Sprintf("%s%s", q1, q2)
	fmt.Println(result)

	// 字符串分割
	url := "www.baidu.com"
	strArray := strings.Split(url, ".")
	fmt.Println(strArray)
	// 字符串Join
	fmt.Println(strings.Join(strArray, "."))
	// 前缀判断
	fmt.Println(strings.HasPrefix(url, "www"))
	// 后缀判断
	fmt.Println(strings.HasSuffix(url, "com"))

	// 字符串遍历
	p := "abc你好"
	for i := 0; i < len(p); i++ {
		fmt.Printf("%c \n", p[i])
	}
	// 使用range解决非单字节字符显示问题
	for _, v := range p {
		fmt.Printf("%c", v)
	}

	// 字符串修改
	// 修改字符串中的字节：使用[]byte
	ss := "hello world"
	value := []byte(ss) // 转换为[]byte
	value[5] = ','      // 将空格替换为“,”
	fmt.Printf("%s\n", ss)
	fmt.Printf("%s\n", value)

	// 修改字符串中的字符：用[]rune
	sss := "一梦三两年"
	value2 := []rune(sss) // 转换为[]rune
	value2[2] = '四'
	value2[3] = '五'
	fmt.Println(sss)
	fmt.Println(string(value2))
}

// 内置strings包常见用法
func stringsTest() {
	s := "helloworld 世界你好"
	fmt.Printf("string:%q\n", s)              // 原文格式输出
	fmt.Printf("rune(char):%q\n", []rune(s))  // 输出[]rune切片
	fmt.Printf("rune(hex):%x\n", []rune(s))   // 采用16进制数表示
	fmt.Printf("bytes(hex):% x\n", []byte(s)) // 输出[]byte切片
	// 判断一个字符串中是否包含某个子串
	fmt.Printf("%t\n", strings.Contains(s, "world"))
	// 检查字符串是不是以某个子串开始
	fmt.Printf("%t\n", strings.HasPrefix(s, "hello"))
	// 检查字符串是不是以某个子串结束
	fmt.Printf("%t\n", strings.HasSuffix(s, "world"))

	// Contains Vs ContainsAny
	fmt.Println(strings.Contains("failure", "a & o")) // false
	fmt.Println(strings.Contains("foo", ""))          // true
	fmt.Println(strings.Contains("", ""))             // true

	fmt.Println(strings.ContainsAny("failure", "a & o")) // true
	fmt.Println(strings.ContainsAny("foo", ""))          // false
	fmt.Println(strings.ContainsAny("", ""))             // false
	fmt.Println(strings.ContainsAny("好树结好果", "好树"))      // true

	// 获取子串索引
	fmt.Println(strings.Index("hello world", "world")) // 6
	fmt.Println(strings.Index("hello world", "hi"))    // -1
	// lastIndex返回匹配到的最后一个字符串的索引
	fmt.Println(strings.LastIndex("hello world,world", "world")) // 12
	// 使用 `IndexRune` 函数查找中文字符出现的位置
	fmt.Println(strings.IndexRune("一梦三两年", '三'))    // 12
	fmt.Printf("rune(char):%q\n", []rune("一梦三两年"))  // 输出[]rune切片
	fmt.Printf("bytes(hex):% x\n", []byte("一梦三两年")) // 输出[]byte切片
}
