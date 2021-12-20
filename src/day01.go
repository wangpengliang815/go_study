package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"runtime"
	"unsafe"
)

func main() {
	// 变量
	variableTest()

	// 常量
	constTest()

	// 数字字面量语法
	numberLiteralsSyntaxTest()

	// 返回变量占用的字节数
	getSizeofTest()

	// 以不同进制显示
	convertOutputTest()

	// 浮点数类型float
	floatTest()

	// float精度丢失问题
	floatPrecisionLossTest()

	// 使用第三方包解决float精度丢失问题
	floatPrecisionLossSolveTest()

	// 使用科学计数法表示浮点数
	float64ScientificTest()

	// 枚举
	enumTest()

	// 数组
	arrayTest()
}

// 变量
func variableTest() {
	printRunFuncName()
	var name string // 省略表达式，使用字符串对应的零值初始化（空字符串）
	var age int     // 省略表达式，使用数值类型对应的零值初始化（0）
	fmt.Println(name, age)

	// 简单声明变量方式
	name2 := "hello"
	fmt.Println(name2)
	// 字符串拼接
	name3 := name2 + "world"
	fmt.Println(name3)

	// 前者表示所见即所得的(除了回车符)。后者所表示的值中转义符会起作用
	var s1 string = "WangPengLiang"
	var s2 string = `WangPengLiang`
	fmt.Println(s1, s2)
}

// 常量
func constTest() {
	printRunFuncName()
	const n = 20        //"const"关键字可以出现在任何"var"关键字出现的地方,区别是常量必须有初始值
	const d = 10000 / n // 常量表达式可以执行任意精度数学计算
	fmt.Println(d)
	fmt.Println(int64(d))    // 数值型常量没有具体类型，除非指定一个类型，比如显式类型转换
	fmt.Println(math.Sin(n)) // 数值型常量可以在程序的逻辑上下文中获取类型，比如变量赋值或者函数调用，例如，对于math包中的Sin函数,它需要一个float64类型的变量
}

// 数字字面量语法
func numberLiteralsSyntaxTest() {
	printRunFuncName()
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

// 以不同进制显示
func convertOutputTest() {
	printRunFuncName()
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
	printRunFuncName()
	var f float64 = 3.1415926
	fmt.Printf("%f\n", f)              //默认保留6位小数
	fmt.Printf("%.2f\n", f)            //保留2位小数
	fmt.Printf("值:%v--类型:%T \n", f, f) // 浮点数默认类型是float64
}

//  float精度丢失问题
func floatPrecisionLossTest() {
	printRunFuncName()
	d := 1129.6
	fmt.Println(d * 100) //此处结果应该是112960，实际打印的却是112959.99999999999

	m1 := 8.2
	fmt.Printf("值:%v--类型:%T \n", m1, m1)
	m2 := 3.8
	fmt.Printf("值:%v--类型:%T \n", m2, m2)
	fmt.Println(m1 - m2) // 此处结果应该是4.4，实际打印4.3999999999999995
}

// @title floatPrecisionLossSolveTest
// @description   使用第三方包解决float精度丢失问题：go get github.com/shopspring/decimal
// @auth     WangPengLiang
// 使用文档：https://pkg.go.dev/github.com/shopspring/decimal#section-readme
func floatPrecisionLossSolveTest() {
	printRunFuncName()
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
	fmt.Println(decimal.NewFromFloat(f1).Add(decimal.NewFromFloat(float64(i1)))) // "2.1 + 3": float和int相加=>output: "5.1"

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
	fmt.Println(decimal.NewFromFloat(float64(n11)).Div(decimal.NewFromFloat(float64(n12)))) // 2 ➗ 3 = 0.6666666666666667  =>  output: "0.6666666666666667"

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

// Go没有枚举,使用const+iota模拟
func enumTest() {
	const (
		Sunday    = iota // 0
		Monday           // 1,通常省略后续行行表达式。
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Println(Sunday, Monday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	const (
		A, B = iota, iota << 10 // 0, 0 << 10
		C, D                    // 1, 1 << 10
	)
	fmt.Println(A, B, C, D)
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

	// 数组比较：数组可以直接进行比较，当数组内的元素都一样的时候表示两个数组相等。
	q := [...]int{11, 22, 33}
	w := [...]int{11, 22, 33}
	fmt.Println(q == w)
	fmt.Println(q == b)

	// 数组作为函数参数传入用的是值传递的方式，就是拷贝。在函数内改变数组元素并不影响外层数组
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

// 打印当前运行函数名称
func printRunFuncName() {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	fmt.Println("FuncName:", f.Name()+"========================================================>")
}
