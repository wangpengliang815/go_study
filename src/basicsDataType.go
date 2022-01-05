package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
	"unsafe"

	"github.com/shopspring/decimal"
)

// 全局常量定义，一旦定义不能修改
const PI = 3.141592653589793

func main() {
	stringTest()
}

// 数字字面量语法，以不同进制定义整形变量
func numberLiteralsSyntaxTest() {
	// 代表二进制的 101101，相当于十进制的 45
	v1 := 0b00101101
	fmt.Printf("value:%v type:%T \n", v1, v1)

	// 代表八进制的377，相当于十进制的 255
	v2 := 0o377
	fmt.Printf("value:%v type:%T \n", v2, v2)

	// 代表十六进制的 1 除以 2²，也就是 0.25
	v3 := 0x1p-2
	fmt.Printf("value:%v type:%T \n", v3, v3)

	// 使用“_”分隔数字
	v4 := 123_456
	fmt.Printf("value:%v type:%T \n", v4, v4)
}

// 返回变量占用的字节数
func getSizeofTest() {
	var value int8 = 120
	fmt.Printf("%T\n", value)
	fmt.Println(unsafe.Sizeof(value))
}

// 借助fmt函数不同进制形式显示整数
func convertOutputTest() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) //%d 表示十进制
	fmt.Printf("%b \n", a) //%b 表示二进制

	// 八进制以0开头
	var b int = 077
	fmt.Printf("%o \n", b) //%o 表示八进制
	fmt.Printf("%d \n", b) //%o 表示八进制

	// 十六进制 以 0x 开头
	var c int = 0xff
	fmt.Printf("%x \n", c) //%x 表示十六进制
	fmt.Printf("%X \n", c) //%X 表示大写的十六进制
	fmt.Printf("%d \n", c) //%d 表示十进制
	fmt.Printf("%T \n", c) //%T 查看变量类型

	// int不同长度转换
	var num1 int8 = 127
	num2 := int32(num1)
	fmt.Printf("value:%v type:%T \n", num2, num2)
}

// 浮点数类型float
func floatTest() {
	var f float64 = 3.1415926
	fmt.Printf("%f\n", f)              //默认保留6位小数
	fmt.Printf("%.2f\n", f)            //保留2位小数
	fmt.Printf("值:%v--类型:%T \n", f, f) // 浮点数默认类型是float64

	fmt.Println(math.MaxFloat32) //输出float32最大值
	fmt.Println(math.MaxFloat64) //输出float64最大值

	// 科学计数法表示浮点类型
	n1 := 5.1234e2
	n2 := 5.1234e2
	n3 := 5.1234e-2
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
}

//  float精度丢失问题
func floatPrecisionLossTest() {
	d := 1129.6
	fmt.Println(d * 100) //此处结果应该是112960，实际打印的却是112959.99999999999

	m1 := 8.2
	fmt.Printf("值:%v--类型:%T \n", m1, m1)
	m2 := 3.8
	fmt.Printf("值:%v--类型:%T \n", m2, m2)
	fmt.Println(m1 - m2) // 此处结果应该是4.4，实际打印4.3999999999999995
}

// 使用第三方包解决float精度丢失问题：go get github.com/shopspring/decimal
// 文档：https://pkg.go.dev/github.com/shopspring/decimal#section-readme
func floatPrecisionLossSolveTest() {
	a := decimal.NewFromFloat(1129.6)
	b := decimal.NewFromInt(100)
	fmt.Println(a.Mul(b)) // output:112960

	c := decimal.NewFromFloat(8.2)
	d := decimal.NewFromFloat(3.8)
	fmt.Println(c.Sub(d)) // output:4.4

	// 初始化一个变量
	d0 := decimal.NewFromFloat(0)
	// 设置精度 为三位 四舍五入的精度
	decimal.DivisionPrecision = 3
	fmt.Println(d0)

	// 加法 Add
	var f1 float64 = 2.1
	var i1 int = 3
	fmt.Println(decimal.NewFromFloat(f1).Add(decimal.NewFromFloat(float64(i1)))) // 2.1 + 3: float和int相加=>output: "5.1"

	var f2 float64 = 2.1
	var f3 float64 = 3.1
	fmt.Println(decimal.NewFromFloat(f2).Add(decimal.NewFromFloat(f3))) // 2.1 + 3.1 (float和float相加)=>output: "5.2"

	var f4 float64 = 2
	var f5 float64 = 3
	fmt.Println(decimal.NewFromFloat(f4).Add(decimal.NewFromFloat(f5))) // 2 + 3 (int和int相加 可以直接相加 d1 = num1+num2)=> output: "5"

	// 	减法 Sub
	var f7 float64 = 3.1
	var f8 int = 2
	d1 := decimal.NewFromFloat(f7).Sub(decimal.NewFromFloat(float64(f8))) // 3.1 - 2 float和int相减 => output: "1.1"
	fmt.Println(d1)

	var n1 float64 = 2.1
	var n2 float64 = 3.1
	fmt.Println(decimal.NewFromFloat(n1).Sub(decimal.NewFromFloat(n2))) // 2.1 - 3.1 (float和float相减)=>output: "-1"

	var n3 int = 2
	var n4 int = 3
	fmt.Println(decimal.NewFromFloat(float64(n3)).Sub(decimal.NewFromFloat(float64(n4)))) // 2 - 3 (int和int相减 d1 = num1 - num2) => output: "-1"

	// 	乘法 Mul
	var n5 float64 = 3.1
	var n6 int = 2
	fmt.Println(decimal.NewFromFloat(n5).Mul(decimal.NewFromFloat(float64(n6)))) // 3.1 * 2 float和int相乘 => output: "6.2"

	var n7 float64 = 2.1
	var n8 float64 = 3.1
	fmt.Println(decimal.NewFromFloat(n7).Mul(decimal.NewFromFloat(n8))) // 2.1 * 3.1 (float和float相乘) => output: "6.51"

	var n9 int = 2
	var n10 int = 3
	fmt.Println(decimal.NewFromFloat(float64(n9)).Mul(decimal.NewFromFloat(float64(n10)))) // 2 * 3 int和int相乘(d1 = num1 * num2) =>  output: "6"

	// 	除法 Div
	var n11 int = 2
	var n12 int = 3
	fmt.Println(decimal.NewFromFloat(float64(n11)).Div(decimal.NewFromFloat(float64(n12)))) // 2 / 3 = 0.6666666666666667  =>  output: "0.6666666666666667"

	// 	float64 和 int 相除
	var num13 float64 = 2.1
	var num14 int = 3
	fmt.Println(decimal.NewFromFloat(num13).Div(decimal.NewFromFloat(float64(num14)))) // output: "0.7"

	// 	float64 和 float64 相除
	var num15 float64 = 2.1
	var num16 float64 = 0.3
	fmt.Println(decimal.NewFromFloat(num15).Div(decimal.NewFromFloat(num16))) // output: "7"

	// int 和 int 相除 并保持3位精度
	var num17 int = 2
	var num18 int = 3
	decimal.DivisionPrecision = 3
	result := decimal.NewFromFloat(float64(num17)).Div(decimal.NewFromFloat(float64(num18))) // 2/3 = 0.667 =>  output: "0.667"
	fmt.Println(result)
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

	// 字符串字面量
	var s1 string = "hello"
	var s2 string = `hello`
	fmt.Println(s1, s2)

	// 内置函数len:返回字符串的字节数
	s := "abc北京"
	fmt.Printf("字节长度：%d \n", len(s)) // output:9
	// 返回字符串每个字节的值
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	//获取字符串中字符的个数
	r := []rune(s)
	fmt.Println(len(r))

	// 字符串截取：
	str := "abcdefg"
	t1 := str[1:4] //startIndex=1,endIndex=4
	t2 := str[:4]  //省略第一个索引从0开始，startIndex=0,endIndex=4
	t3 := str[1:]  //省略第二个索引，从1开始截取到末尾
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

// 常量
func constTest() {
	// 显式类型定义,常量名称推荐全大写
	const LENGTH int = 10
	// 隐式类型定义,其实也是使用类型推导
	const WIDTH = 20
	// 多重定义
	const a, b, c = 1, false, "hello world"
	fmt.Println(a, b, c)

	area := LENGTH * WIDTH
	fmt.Println(area)
}

// Go没有枚举,使用const+iota模拟
func enumTest() {
	// 普通常量组
	const (
		Windows = 0
		Linux
		MaxOS
	)
	// 普通常量组如果不指定类型和初始化值，则与上一行非空常量右值相同
	fmt.Println(Windows, Linux, MaxOS)

	/*
	 结合iota实现枚举
	 第一个 iota 等于0，每当 iota 在新的一行被使用时，它的值都会自动加 1
	*/
	const (
		Sunday    = iota // 0
		Monday           // 1,通常省略后续行行表达式。
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	// 插队
	const (
		a1 = iota    // 0
		a2           // 1
		b1 = "hello" // 独立值hello,iota+=1
		b2           // 如不指定类型和初始化值，则与上一行非空常量右值相同,所以是hello;iota+=1
		c1 = iota    // 恢复自增,4
		c2           // 5
	)
	fmt.Println(a1, a2, b1, b2, c1, c2)

	// 使用_忽略某些值
	const (
		d1 = iota
		d2
		_
		d3
	)
	fmt.Println(d1, d2, d3)

	// 定义数量级别
	const (
		_  = iota             // _表示将0忽略
		KB = 1 << (10 * iota) // 表示1左移十位，转换为十进制就是1024
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	// 多个iota一行定义
	const (
		a, b = iota + 1, iota + 2 // a=iota+1 b=iota+2=> 1,2
		c, d                      // c=iota(1)+1 b=iota(2)+1=> 2,3
	)
	fmt.Println(a, b, c, d)
}
