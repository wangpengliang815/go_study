// @title 《Go语言编程》-基础数据类型字符串
// @description
// @author wangpengliang
// @date 2022-03-25 10:58:48
package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 内置函数len：返回字符串的字节数
func getByteCount() {
	str := "abc北京"
	fmt.Printf("字节长度：%d \n", len(str)) // output:9

	// 返回字符串每个字节的值
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
}

// 获取字符个数
func getCharCount() {
	// TODO：这里需要了解一下Utf-8和Unicode的区别
	str := "abc北京"
	// 获取字符串中字符的个数
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i])
	}
	fmt.Println(len(r))
}

// 字符串字面量语法：可以用来定义一个多行字符串
func stringLiteral() {
	var s1 string = "hello"
	var s2 string = `hello
				 world`
	fmt.Println(s1, s2)
}

// 一个 string 类型的值既可以被拆分为一个包含多个字符的序列，也可以被拆分为一个包含多个字节的序列。
func stringRunByte() {
	str := "helloworld 世界你好"
	fmt.Printf("string:%q\n", str)              // 原文格式输出
	fmt.Printf("rune(char):%q\n", []rune(str))  // 输出[]rune切片
	fmt.Printf("rune(hex):%x\n", []rune(str))   // 采用16进制数表示
	fmt.Printf("bytes(hex):% x\n", []byte(str)) // 输出[]byte切片
}

// 字符串截取
func subString() {
	str := "abcdefg"
	t1 := str[1:4] // startIndex=1,endIndex=4
	t2 := str[:4]  // 省略第一个索引从0开始，startIndex=0,endIndex=4
	t3 := str[1:]  // 省略第二个索引，从1开始截取到末尾
	// t4 := str[:10] //runtime error: slice bounds out of range [:10] with length 7 数组越界
	fmt.Println(t1, t2, t3)
}

// 字符串拼接的方式
func joinString() {
	// 第一种：直接使用“+”号拼接
	q1 := "hello"
	q2 := ",world"
	fmt.Println(q1 + q2)

	// 第二种：使用Printf拼接输出
	// %s 代表原样输出
	fmt.Printf("%s%s\n", q1, q2)

	// 第三种：使用Sprintf拼接,不会直接打印而是生成一个新的字符串
	var result = fmt.Sprintf("%s%s", q1, q2)
	fmt.Println(result)

	// 第四种：使用strings.Join
	url := "www.baidu.com"
	strArray := strings.Split(url, ".")
	fmt.Println(strings.Join(strArray, ","))
}

// 字符串分割
func splitStringTest() {
	url := "www.baidu.com"
	strArray := strings.Split(url, ".")
	fmt.Println(strArray)
}

// 字符串前缀后缀判断
func hasPrefixAndHasSuffix() {
	url := "www.baidu.com"
	// 前缀判断
	fmt.Println(strings.HasPrefix(url, "www"))
	// 后缀判断
	fmt.Println(strings.HasSuffix(url, "com"))
}

// 修改字符串中的字节：使用[]byte
func updateStringByte() {
	str := "hello world"
	value := []byte(str) // 转换为[]byte
	value[5] = ','       // 将空格替换为“,”
	fmt.Printf("%s\n", str)
	fmt.Printf("%s\n", value)
}

// 修改字符串中的字符：用[]rune
func updateStringChar() {
	str := "一梦三两年"
	value := []rune(str) // 转换为[]rune
	value[2] = '四'
	value[3] = '五'
	fmt.Println(str)
	fmt.Println(string(value))
}

// 字符串遍历
func foreachString() {
	p := "abc你好"
	for i := 0; i < len(p); i++ {
		fmt.Printf("%c \n", p[i])
	}

	// 使用range解决非单字节字符显示问题
	for _, v := range p {
		fmt.Printf("%c \n", v)
	}
}

// 内置strings包常见用法
func stringsTools() {
	s := "helloworld 世界你好"
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

// 练习题1：获取字符串中存在几个汉字
func exercises1() {
	name := "hello,勒布朗、科比、艾弗森"
	var count int
	for _, c := range name {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)
}

// 练习题2：获取字符串中每个单词出现的次数
func exercises2() {
	s := "how do you do"
	// 以空格分隔
	s2 := strings.Split(s, " ")
	// 定义map用于存储每个单词出现的次数
	m1 := make(map[string]int, 10)
	// 遍历slice
	for _, w := range s2 {
		// 如果map中不存在key则添加
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			// 否则该key的value+1
			m1[w]++
		}
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}

// 练习题3：回文判断
func exercises3() {
	s := "上海自来水来自海上"
	// 创建rune类型切片
	r := make([]rune, 0, len(s))

	for _, v := range s {
		r = append(r, v)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("no")
			return
		}
	}
	fmt.Println("yes")
}
