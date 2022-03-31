// @title 《Go语言编程》-反射
// @description
// @author wangpengliang
// @date 2022-03-31 13:36:54

package main

import (
	"fmt"
	"reflect"
)

// 反射获取结构体
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写,小写表示不可导出反射无法读取到)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

type myInt int64

func main() {
	getTypeAndKindTest()
}

// reflect.TypeOf()可以获取任意值的类型对象（reflect.Type）
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

// reflect.ValueOf()可以获取reflect.Value类型，其中包含了原始值的值信息
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置值时如果函数参数传递的是值拷贝
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("value:%v,kind:%v", v, v.Kind())
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本,reflect包会引发panic
	}
}

// 通过反射设置值时必须传递变量地址才能修改变量值
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// 通过反射获取对象的类型reflect.Type
func getType() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
	var c myInt = 200
	reflectType(c) // type:main.myInt
}

// 通过反射获取对象的值reflect.Value
func getValue() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
}

// 通过反射设置变量的值
func setValue() {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}

// IsNil()常被用于判断指针是否为空,IsValid()常被用于判定返回值是否有效
func isNilAndisValid() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())

	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())

	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("wang")).IsValid())
}

// 结构体反射
func reflectStruct() {
	stu := student{
		Name:  "wangpengliang",
		Score: 90,
	}

	t := reflect.TypeOf(stu)
	fmt.Println(t.Name(), t.Kind()) // student struct

	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

// 反射获取结构体方法并调用
func printMethod(x interface{}) {
	rType := reflect.TypeOf(x)

	rValue := reflect.ValueOf(x)

	fmt.Println(rType.NumMethod())
	for i := 0; i < rValue.NumMethod(); i++ {
		methodType := rValue.Method(i).Type()
		fmt.Printf("method name:%s\n", rType.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)

		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		rValue.Method(i).Call(args)
	}
}

// 获取类型和种类
func getTypeAndKind(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func getTypeAndKindTest() {
	var a *float32    // 指针
	var b myInt       // 自定义类型
	var c rune        // 类型别名
	getTypeAndKind(a) // type: kind:ptr
	getTypeAndKind(b) // type:myInt kind:int64
	getTypeAndKind(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "wangpengliang",
		age:  18,
	}
	var e = book{title: "《Go语言编程》"}
	getTypeAndKind(d) // type:person kind:struct
	getTypeAndKind(e) // type:book kind:struct
}
