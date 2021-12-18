package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"runtime"
	"unsafe"
)

func main() {
	variableTest()

	constTest()

	// 返回变量占用的字节数
	var value int8 = 120
	fmt.Printf("%T\n", value)
	fmt.Println(unsafe.Sizeof(value))

	// int不同长度转换
	var num1 int8 = 127
	num2 := int32(num1)
	fmt.Printf("值:%v 类型%T", num2, num2)

	numberLiteralsSyntaxTest()

	convertOutputTest()

	floatTest()

	floatPrecisionLossTest()

	floatPrecisionLossSolveTest()
}

// @title variableTest
// @description   变量
// @auth     WangPengLiang
func variableTest() {
	printRunFuncName()
	var isCheck = false // 省略变量类型，根据value推断变量类型
	var name string     //省略表达式，使用字符串对应的零值初始化（空字符串）
	var age int         //省略表达式，使用数值类型对应的零值初始化（0）
	fmt.Println(isCheck, name, age)

	// 简单声明变量方式
	name2 := "lizimeng"
	fmt.Println(name2)
	// 字符串拼接
	name3 := name2 + "hello"
	fmt.Println(name3)
}

// @title constTest
// @description   常量
// @auth     WangPengLiang
func constTest() {
	printRunFuncName()
	const n = 20        //"const"关键字可以出现在任何"var"关键字出现的地方,区别是常量必须有初始值
	const d = 10000 / n // 常量表达式可以执行任意精度数学计算
	fmt.Println(d)
	fmt.Println(int64(d))    // 数值型常量没有具体类型，除非指定一个类型，比如显式类型转换
	fmt.Println(math.Sin(n)) // 数值型常量可以在程序的逻辑上下文中获取类型，比如变量赋值或者函数调用，例如，对于math包中的Sin函数,它需要一个float64类型的变量
}

// @title numberLiteralsSyntaxTest
// @description   数字字面量语法
// @auth     WangPengLiang
func numberLiteralsSyntaxTest() {
	printRunFuncName()
	v1 := 0b00101101 //代表二进制的 101101，相当于十进制的 45
	fmt.Printf("值:%v 类型%T \n", v1, v1)
	v2 := 0o377 //代表八进制的377，相当于十进制的 255
	fmt.Printf("值:%v 类型%T \n", v2, v2)
	v3 := 0x1p-2 //代表十六进制的 1 除以 2²，也就是 0.25
	fmt.Printf("值:%v 类型%T \n", v3, v3)
	v4 := 123_456 // 使用“_”分隔数字
	fmt.Printf("值:%v 类型%T \n", v4, v4)
}

// @title convertOutputTest
// @description   以不同进制显示
// @auth     WangPengLiang
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
}

// @title floatTest
// @description   浮点数类型float
// @auth     WangPengLiang
func floatTest() {
	printRunFuncName()
	var f float64 = 3.1415926
	fmt.Printf("%f\n", f)              //默认保留6位小数
	fmt.Printf("%.2f\n", f)            //保留2位小数
	fmt.Printf("值:%v--类型:%T \n", f, f) // 浮点数默认类型是float64
}

// @title floatPrecisionLossTest
// @description   float精度丢失问题
// @auth     WangPengLiang
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
	fmt.Println(decimal.NewFromFloat(f1).Add(decimal.NewFromFloat(float64(i1)))) // "2.1 + 3": float和int相加,output: "5.1"

	var f2 float64 = 2.1
	var f3 float64 = 3.1
	fmt.Println(decimal.NewFromFloat(f2).Add(decimal.NewFromFloat(f3))) //2.1 + 3.1 (float和float相加), output: "5.2"

	var f4 float64 = 2
	var f5 float64 = 3
	fmt.Println(decimal.NewFromFloat(f4).Add(decimal.NewFromFloat(f5))) //2 + 3 (int和int相加 可以直接相加（d1 = num1+num2）), output: "5"

	// 	减法 Sub
	var f7 float64 = 3.1
	var f8 int = 2
	d1 := decimal.NewFromFloat(num1).Sub(decimal.NewFromFloat(float64(num2))) // 	3.1 - 2 float和int相减
	fmt.Println(d1)                                                           // output: "1.1"

	var num1 float64 = 2.1
	var num2 float64 = 3.1
	d1 := decimal.NewFromFloat(num1).Sub(decimal.NewFromFloat(num2)) // 	2.1 - 3.1 (float和float相减)
	fmt.Println(d1)                                                  // output: "-1"

	var num1 int = 2
	var num2 int = 3
	d1 := decimal.NewFromFloat(float64(num1)).Sub(decimal.NewFromFloat(float64(num2))) // 	2 - 3 (int和int相减(d1 = num1 - num2))
	fmt.Println(d1)                                                                    // output: "-1"
	// 	乘法 Mul

	var num1 float64 = 3.1
	var num2 int = 2
	d1 := decimal.NewFromFloat(num1).Mul(decimal.NewFromFloat(float64(num2))) // 3.1 * 2 float和int相乘
	fmt.Println(d1)                                                           // output: "6.2"

	var num1 float64 = 2.1
	var num2 float64 = 3.1
	d1 := decimal.NewFromFloat(num1).Mul(decimal.NewFromFloat(num2)) // 	2.1 * 3.1 (float和float相乘)
	fmt.Println(d1)                                                  // output: "6.51"

	var num1 int = 2
	var num2 int = 3

	d1 := decimal.NewFromFloat(float64(num1)).Mul(decimal.NewFromFloat(float64(num2))) // 2 * 3 int和int相乘(d1 = num1 * num2)
	fmt.Println(d1)                                                                    // output: "6"
	// 	除法 Div

	var num1 int = 2
	var num2 int = 3
	d1 := decimal.NewFromFloat(float64(num1)).Div(decimal.NewFromFloat(float64(num2))) // 	2 ➗ 3 = 0.6666666666666667
	fmt.Println(d1)                                                                    // output: "0.6666666666666667"

	// 		float64 和 int 相除
	var num1 float64 = 2.1
	var num2 int = 3
	d1 := decimal.NewFromFloat(num1).Div(decimal.NewFromFloat(float64(num2)))
	fmt.Println(d1) // output: "0.7"

	// 	float64 和 float64 相除
	var num1 float64 = 2.1
	var num2 float64 = 0.3
	d2 := decimal.NewFromFloat(num1).Div(decimal.NewFromFloat(num2)) // 		2.1 / 0.3 = 7
	fmt.Println(d2)                                                  // output: "7"

	// int 和 int 相除 并保持3位精度
	var num1 int = 2
	var num2 int = 3
	decimal.DivisionPrecision = 3
	d1 := decimal.NewFromFloat(float64(num1)).Div(decimal.NewFromFloat(float64(num2))) // 		2 ➗ 3 = 0.667
	fmt.Println(d1)                                                                    // output: "0.667"
}

// 获取当前运行函数名称
func printRunFuncName() {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	fmt.Println("FuncName:", f.Name()+"========================================================>")
}
