package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
	"unsafe"
)

func main() {
	sliceTest()
}

// 数字字面量语法，以不同进制定义变量
func numberLiteralsSyntaxTest() {
	v1 := 0b00101101 //代表二进制的 101101，相当于十进制的 45
	fmt.Printf("value:%v type:%T \n", v1, v1)
	v2 := 0o377 //代表八进制的377，相当于十进制的 255
	fmt.Printf("value:%v type:%T \n", v2, v2)
	v3 := 0x1p-2 //代表十六进制的 1 除以 2²，也就是 0.25
	fmt.Printf("value:%v type:%T \n", v3, v3)
	v4 := 123_456 // 使用“_”分隔数字
	fmt.Printf("value:%v type:%T \n", v4, v4)
}

// 返回变量占用的字节数
func getSizeofTest() {
	var value int8 = 120
	fmt.Printf("%T\n", value)
	fmt.Println(unsafe.Sizeof(value))
}

// 借助 fmt 函数不同进制形式显示整数
func convertOutputTest() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a) // 占位符%b表示二进制

	// 八进制以0开头
	var b int = 077
	fmt.Printf("%o \n", b)

	// 十六进制 以 0x 开头
	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)
	fmt.Printf("%d \n", c)

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
	q3 := q1 + q2
	fmt.Println(q3)

	// 字符串遍历
	p := "abc你好"
	for i := 0; i < len(p); i++ {
		fmt.Printf("%c", p[i])
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

// 数组
func arrayTest() {
	// 数组创建，使用类型零值初始化
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// 数组创建使用数字字面值语法
	var b [3]int = [3]int{1, 2, 3}
	fmt.Println(b[0])
	// fmt.Println(a[10]) invalid array index 10 (out of bounds for 3-element array),数组越界

	//如果数组的长度位置出现的是“…”省略号，表示数组的长度是根据初始化值的个数来计算
	c := [...]int{1, 2, 3}
	fmt.Println(c[0])

	// 数组遍历
	for i := 0; i < len(a); i++ {
		fmt.Printf("key:%d, value:%d\n", i, a[i])
	}

	// 使用range遍历数组
	for i, v := range a {
		fmt.Printf("key:%d, value:%d\n", i, v)
	}

	// 使用range遍历数组:如果需要值并希望忽略索引，可以通过使用_ blank标识符替换索引来实现
	for _, v := range a {
		fmt.Printf("value:%d\n", v)
	}

	// go同样支持多维数组
	array := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /*  第三行索引为 2 */
	}
	for i, item := range array {
		fmt.Printf("%d the element of is %v \n", i, item)
	}

	// 数组比较：数组可以直接进行比较，当数组内的元素都一样的时候表示两个数组相等。
	q := [...]int{11, 22, 33}
	w := [...]int{11, 22, 33}
	fmt.Println(q == w)
	fmt.Println(q == b)

	// 数组是值类型：意味着当它被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改不会在原始数组中反映
	e := [...]int{11, 22, 33}
	fmt.Println(e[0])
	arrayTest2(e)
	fmt.Println(e[0])

	// 使用指针后函数内部对数组的更改将反应到原数组上
	r := [...]int{11, 22, 33}
	fmt.Println(r[0])
	arrayTest3(&r)
	fmt.Println(r[0])
}

func arrayTest2(arr [3]int) {
	arr[0] = 1
}

func arrayTest3(arr *[3]int) {
	arr[0] = 100
}

// Slice切片
func sliceTest() {
	// 切片的定义
	var s1 []int    //定义存放int类型的切片
	var s2 []string //定义存放string类型的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	// 切片初始化
	s1 = []int{1, 3, 4, 5, 67, 88}
	s2 = []string{"北京", "上海", "山西"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Printf("len(s1):%d cap(s1):%d \n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d \n", len(s2), cap(s2))

	// 由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a1)
	fmt.Println(a1[0:4]) // =>[1 2 3 4]基于数组得到切片,从0开始到第4个结束（不包含4）.原则：左包含右不包含
	fmt.Println(a1[:4])  // =>[1 2 3 4] 省略第一个参数，默认从0开始
	fmt.Println(a1[3:])  // =>[4 5 6 7 8 9]省略第二个参数，默认到len(a1)结束
	fmt.Println(a1[:])   // =>[1 2 3 4 5 6 7 8 9] 两个参数都省略，默认从0开始到len(a1-1)结束

	// 切片的长度和容量
	s3 := a1[3:] //[4 5 6 7 8 9]
	fmt.Println(s3)
	// 切片的长度是元素的个数，切片的容量是底层数组从切片的第一个元素到最后一个元素
	fmt.Printf("len(s3):%d cap(s3):%d \n", len(s3), cap(s3))

	s4 := a1[4:8] //[5 6 7 8]
	fmt.Println(s4)
	fmt.Printf("len(s4):%d cap(s4):%d \n", len(s4), cap(s4))

	//由切片得到切片
	s5 := s3[2:4]
	fmt.Println(s5)
	fmt.Printf("len(s5):%d cap(s5):%d \n", len(s5), cap(s5))

	//// 使用make创建一个长度和容量都为5的切片
	//slice1 := make([]string, 5)
	////使用make创建一个长度5，容量为10的切片
	//slice2 := make([]string, 5, 10)
	//fmt.Println(slice1, slice2)
	//// fmt.Println(slice2[6]) // 虽然创建的切片对应底层数组的大小为 10，但是不能访问索引值 5 以后的元素
	//
	//// 通过切片字面量创建
	//slice3 := []int{1, 2, 4, 4}
	//fmt.Println(slice3)
	//
	//// 创建空切片
	//slice4 := []int{}
	//slice5 := make([]int, 0)
	//fmt.Println(slice4, slice5)
	//
	//// 通过切片创建切片
	//slice6 := []int{10, 20, 30, 40, 50}
	//slice7 := slice6[1:3]
	//fmt.Println(slice6, slice7)
}
