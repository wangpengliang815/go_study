package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

func main() {
	stringsTest()
}

// 浮点数类型float
func floatTest() {
	var f float64 = 3.1415926
	fmt.Printf("%f\n", f)              //默认保留6位小数
	fmt.Printf("%.2f\n", f)            //保留2位小数
	fmt.Printf("值:%v--类型:%T \n", f, f) // 浮点数默认类型是float64
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

// 科学计数法表示浮点类型
func float64ScientificTest() {
	n1 := 5.1234e2
	n2 := 5.1234e2
	n3 := 5.1234e-2
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
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
	s := "helloworld"
	fmt.Printf("string:%q\n", s)
	fmt.Printf("rune(char):%q\n", []rune(s))
	fmt.Printf("rune(hex):%x\n", []rune(s))
	fmt.Printf("bytes(hex):% x\n", []byte(s))

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
}

// 数组
func arrayTest() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a[0])
	// invalid array index 10 (out of bounds for 3-element array),数组越界
	// fmt.Println(a[10])

	b := [...]int{1, 2, 3} //如果数组的长度位置出现的是“…”省略号，表示数组的长度是根据初始化值的个数来计算
	fmt.Println(b[0])

	// 数组遍历
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	// 使用range遍历数组
	for i, item := range b { //range returns both the index and value
		fmt.Printf("%d the element of is %v \n", i, item)
	}

	// 如果需要值并希望忽略索引，可以通过使用_ blank标识符替换索引来实现
	for _, item := range b { //range returns both the index and value
		fmt.Printf("element of is %v \n", item)
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
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := make([]int, 5)
	fmt.Println(slice)

	// 创建一个整型切片
	// 其长度为 3 个元素，容量为 5 个元素
	slice2 := make([]int, 3, 5)
	fmt.Println(slice2)

	//1、make
	a := make([]int32, 0, 5)
	//2、[]int32{}
	b := []int32{1, 2, 3}
	//3、new([]int32)
	c := *new([]int32)
	fmt.Println(a, b, c)
}
