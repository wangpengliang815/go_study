package main

import "fmt"

func main() {
	mapTest()
}

// map
func mapTest() {
	// 正常创建map
	ageMap := make(map[string]int, 8)
	ageMap["小王"] = 18
	ageMap["小李"] = 20
	fmt.Println(ageMap)
	fmt.Println(ageMap["小王"])
	fmt.Printf("type of a:%T\n", ageMap)

	// 声明时候填充map
	ageMap2 := map[string]int{
		"小周": 22,
		"小张": 23,
	}
	fmt.Println(ageMap2)
	fmt.Println(ageMap2["小王"])
	fmt.Printf("type of a:%T\n", ageMap2)

	// 判断键是否存在：如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := ageMap2["小张"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}

	// 遍历map
	for k, v := range ageMap2 {
		fmt.Println(k, v)
	}
	// 只想遍历key
	for k := range ageMap2 {
		fmt.Println(k)
	}

	// delete删除某个键值对
	ageMap3 := map[string]int{
		"小周": 21,
		"小张": 22,
	}
	for k, v := range ageMap3 {
		fmt.Println(k, v)
	}
	delete(ageMap3, "小周")
	for k, v := range ageMap3 {
		fmt.Println(k, v)
	}
}

// slice切片
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

	// 切片的赋值拷贝
	w1 := make([]int, 3) //[0 0 0]
	w2 := w1             //将w1直接赋值给w2，w1和w2共用一个底层数组
	w2[0] = 100
	fmt.Println(w1) //[100 0 0]
	fmt.Println(w2) //[100 0 0]

	// 切片遍历
	for i := 0; i < len(w1); i++ {
		fmt.Println(i, w1[i])
	}

	for index, value := range w1 {
		fmt.Println(index, value)
	}

	// append()追加元素到切片中
	w3 := make([]int, 3)           // 创建切片：[0 0 0]
	w3 = append(w3, 1)             // 切片中添加第一个元素 1
	w3 = append(w3, 2, 3, 4, 5, 6) // 继续添加元素 2,3,4,5,6
	w4 := []int{7, 8, 9}           // 创建新的切片
	w3 = append(w3, w4...)         // 将新的切片中的元素都放到w3中,这里...代表将w4中的元素拆分
	fmt.Println(w3)                // 输出：[0 0 0 1 2 3 4 5 6 7 8 9]

	// append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	// copy()
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]

	c1 := []int{1, 2, 3, 4, 5}
	c2 := make([]int, 5, 5)
	copy(c2, c1)    //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(c1) //[1 2 3 4 5]
	fmt.Println(c2) //[1 2 3 4 5]
	c1[0] = 1000
	fmt.Println(c1) //[1000 2 3 4 5]
	fmt.Println(c2) //[1 2 3 4 5]

	// 从切片中删除元素
	c3 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素32
	c3 = append(c3[:2], c3[3:]...) //其实这就是利用append的特性修改了切片内容再返回
	fmt.Println(c3)                //[30 31 33 34 35 36 37]
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
