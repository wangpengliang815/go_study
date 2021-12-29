package main

import "fmt"

func main() {
	sliceTest()
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

	// 使用make创建一个长度和容量都为5的切片
	slice1 := make([]string, 5)
	//使用make创建一个长度5，容量为10的切片
	slice2 := make([]string, 5, 10)
	fmt.Println(slice1, slice2)
	// fmt.Println(slice2[6]) // 虽然创建的切片对应底层数组的大小为 10，但是不能访问索引值 5 以后的元素,其实相当于底层数组长度是10但是切片只覆盖到了0~5

	// 切片比较
	var q1 []int //len(q1)=0;cap(q1)=0;q1==nil
	fmt.Printf("len(q1):%d cap(q1):%d q1==nil:%t \n", len(q1), cap(q1), q1 == nil)
	q2 := []int{} //len(q2)=0;cap(q2)=0;q2!=nil
	fmt.Printf("len(q2):%d cap(q2):%d q2==nil:%t \n", len(q2), cap(q2), q2 == nil)
	q3 := make([]int, 0) //len(q3)=0;cap(q3)=0;q3!=nil
	fmt.Printf("len(q3):%d cap(q3):%d q3==nil:%t \n", len(q3), cap(q3), q3 == nil)

	// TODO:待补充
}
