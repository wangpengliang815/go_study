// @title 《Go语言编程》- fmt包使用
// @description
// @author wangpengliang
// @date 2022-04-02 16:39:40
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmtFprint()
}

// print
func fmtPrint() {
	fmt.Print("在终端打印该信息")
	name := "wangpengliang"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
}

// fprint:往文件写入内容
func fmtFprint() {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./fprintTest.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "wangpengliang"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写入信息：%s", name)
}

// sprint:把传入的数据生成并返回一个字符串
func fmtSprint() {
	s := "wangpengliang"
	s1 := fmt.Sprint(s)

	s2 := fmt.Sprintf("name:%s,age:%d", "wangpengliang", 18)

	fmt.Println(s1, s2)
}

// 根据 format 参数生成格式化字符串并返回一个包含该字符串的错误
func errorof() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Println(w)
}

// 通用占位符
func generalPlaceholder() {
	fmt.Printf("%v\n", 100)   // %v：默认格式输出
	fmt.Printf("%v\n", false) // %v：默认格式输出
	o := struct{ name string }{"wangpengliang"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o) // %+v：类似%v，但输出结构体时会添加字段名
	fmt.Printf("%#v\n", o) // %#v：值的Go语法表示
	fmt.Printf("%T\n", o)  // %T：打印值的类型
	fmt.Printf("100%%\n")  // %%：百分号
}

// 整型占位符
func intPlaceholder() {
	n := 65
	fmt.Printf("%b\n", n) // %b：二进制输出
	fmt.Printf("%c\n", n) // %c：输出对应的unicode编码
	fmt.Printf("%d\n", n) // %d:十进制输出
	fmt.Printf("%o\n", n) // %o:八进制输出
	fmt.Printf("%x\n", n) // %x:十六进制输出,使用a-f
	fmt.Printf("%X\n", n) // %X:十六进制输出，使用A-F
}

// 浮点数与复数占位符
func floatPlaceholder() {
	f := 12.34
	fmt.Printf("%b\n", f) // %b：无小数部分、二进制指数的科学计数法
	fmt.Printf("%e\n", f) // %e: 科学计数法,如-1234.456e+78
	fmt.Printf("%E\n", f) // %E: 科学计数法,如-1234.456E+78
	fmt.Printf("%f\n", f) // %f: 有小数部分但无指数部分，如123.456
	fmt.Printf("%g\n", f) // %g: 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	fmt.Printf("%G\n", f) // %G: 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
}

// 字符串与[]byte占位符
func stringPlaceholder() {
	s := "wangpengliang"
	fmt.Printf("%s\n", s) // %s：直接输出字符串或者[]byte
	fmt.Printf("%q\n", s) // %q: 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	fmt.Printf("%x\n", s) // %x: 每个字节用两字符十六进制数表示（使用a-f)
	fmt.Printf("%X\n", s) // %X: 每个字节用两字符十六进制数表示（使用A-F)
}

// 指针占位符
func pointerPlaceholder() {
	a := 10
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)
}

// 宽度标识符
func breadth() {
	n := 12.34
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%9.2f\n", n)
	fmt.Printf("%9.f\n", n)
}

// 其他flag
func otherPlaceholder() {
	s := "wangpengliang"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)
}

// Scan获取用户输入,空格分隔
func fmtScan() {
	fmt.Println("please input you name")
	var input string
	fmt.Scanln(&input) //读取键盘输入，通过操作地址赋值给input.阻塞式
	fmt.Println("you name is：" + input)

	fmt.Println("please input you name,age,sex")
	var (
		name string
		age  int
		sex  bool
	)
	fmt.Scan(&name, &age, &sex)
	fmt.Printf("扫描结果 name:%s age:%d sex:%t \n", name, age, sex)
}

// Scanf获取用户输入,指定输入格式
func fmtScanf() {
	var (
		name string
		age  int
		sex  bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &sex)
	fmt.Printf("扫描结果 name:%s age:%d sex:%t \n", name, age, sex)
}

// Scanln
func fmtScanlnf() {
	var (
		name string
		age  int
		sex  bool
	)
	fmt.Scanln(&name, &age, &sex)
	fmt.Printf("扫描结果 name:%s age:%d sex:%t \n", name, age, sex)
}

// bufio：https://golang.google.cn/pkg/bufio
func bufioTest() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
