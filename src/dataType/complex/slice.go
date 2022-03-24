package main

import "fmt"

func main() {
	delSliceTest()
}

// 切片定义及初始化
func createSliceTest() {
	var s1 []int    // 定义存放int类型的切片
	var s2 []string // 定义存放string类型的切片
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
}

// 创建空切片表示一个空集合
func createEmptySliceTest() {
	slice1 := []int{}
	slice2 := make([]int, 0)
	fmt.Println(slice1, slice2)
}

// 由数组得到切片
func createSliceByArrayTest() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	fmt.Println(arr[0:4]) // =>[1 2 3 4] 基于数组得到切片,从0开始到第4个结束（不包含4）.原则：左包含右不包含
	fmt.Println(arr[:4])  // =>[1 2 3 4] 省略第一个参数，默认从0开始
	fmt.Println(arr[3:])  // =>[4 5 6 7 8 9] 省略第二个参数，默认到len(a1)结束
	fmt.Println(arr[:])   // =>[1 2 3 4 5 6 7 8 9] 两个参数都省略，默认从0开始到len(a1-1)结束
}

// 由切片得到切片
func createSliceBySliceTest() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr1 := arr[0:4]
	fmt.Println(arr1)
	fmt.Printf("len(s5):%d cap(s5):%d \n", len(arr1), cap(arr1))
	//由切片得到切片
	arr2 := arr1[2:4]
	fmt.Println(arr2)
	fmt.Printf("len(s5):%d cap(s5):%d \n", len(arr2), cap(arr2))
}

// 使用make创建一个长度和容量都为5的切片
func createSliceByMakeTest() {
	slice1 := make([]string, 5)
	// 使用make创建一个长度5，容量为10的切片
	slice2 := make([]string, 5, 10)
	fmt.Println(slice1, slice2)
	// fmt.Println(slice2[6]) // 虽然创建的切片对应底层数组的大小为 10，但是不能访问索引值 5 以后的元素,其实相当于底层数组长度是10但是切片只覆盖到了0~5
}

// 获取切片长度和容量
func getSliceLenAndCapTest() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[3:] // [4 5 6 7 8 9]
	fmt.Println(s1)
	// 切片的长度是元素的个数,切片的容量是底层数组从切片的第一个元素到最后一个元素,
	fmt.Printf("len(s1):%d cap(s1):%d \n", len(s1), cap(s1))

	s2 := arr[4:8] // [5 6 7 8]
	fmt.Println(s2)
	// 切片的长度是元素的个数,所以len=4,切片的容量是底层数组从切片的第一个元素到最后一个元素,所以这里就是从4到9
	fmt.Printf("len(s2):%d cap(s2):%d \n", len(s2), cap(s2))
}

// 切片比较
func compareSliceTest() {
	var q1 []int // len(q1)=0;cap(q1)=0;q1==nil; 没有被初始化所以q1==nil is true
	fmt.Printf("len(q1):%d cap(q1):%d q1==nil:%t \n", len(q1), cap(q1), q1 == nil)

	q2 := []int{} // len(q2)=0;cap(q2)=0;q2!=nil; 这里是定义了元素为空的数组,所以q2==nil is false
	fmt.Printf("len(q2):%d cap(q2):%d q2==nil:%t \n", len(q2), cap(q2), q2 == nil)

	q3 := make([]int, 0) // len(q3)=0;cap(q3)=0; q3!=nil; 这里使用了make所以十分分配内存的只不过cap和len都为0而已,所以q3==nil is false
	fmt.Printf("len(q3):%d cap(q3):%d q3==nil:%t \n", len(q3), cap(q3), q3 == nil)
}

// 切片共享一个底层数组
func shareArraySliceTest() {
	w1 := make([]int, 3) // [0 0 0]
	w2 := w1             // 将w1直接赋值给w2，w1和w2共用一个底层数组
	w2[0] = 100
	fmt.Println(w1) // [100 0 0]
	fmt.Println(w2) // [100 0 0]
}

// 切片遍历
func foreachSliceTest() {
	slice := make([]int, 3) // [0 0 0]
	for i := 0; i < len(slice); i++ {
		fmt.Println(i, slice[i])
	}
	for index, value := range slice {
		fmt.Println(index, value)
	}
}

// append()追加元素到切片中,如果使用append()切片可以不被初始化,会自动扩容并添加元素
func appendSliceTest() {
	slice := make([]int, 3)              // 创建切片：[0 0 0]
	slice = append(slice, 1)             // 切片中添加第一个元素 1
	slice = append(slice, 2, 3, 4, 5, 6) // 继续添加元素 2,3,4,5,6
	slice2 := []int{7, 8, 9}             // 创建新的切片
	slice = append(slice, slice2...)     // 将新的切片中的元素都放到w3中,这里...代表将slice2中的元素拆分
	fmt.Println(slice)                   // 输出：[0 0 0 1 2 3 4 5 6 7 8 9]
}

// append()扩容策略
func appendDilatationSliceTest() {
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

//  copy()不会自动扩容
func copySliceTest() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(b) // [1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) // [1000 2 3 4 5]
	fmt.Println(b) // [1000 2 3 4 5]

	c := []int{1, 2, 3, 4, 5}
	d := make([]int, 5)
	copy(d, c)     // 使用copy()函数将切片c1中的元素复制到切片c2
	fmt.Println(c) // [1 2 3 4 5]
	fmt.Println(d) // [1 2 3 4 5]
	c[0] = 1000
	fmt.Println(c) // [1000 2 3 4 5]
	fmt.Println(d) // [1 2 3 4 5]
}

// 从切片中删除元素
func delSliceTest() {
	// 从切片中删除元素
	c3 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素32
	c3 = append(c3[:2], c3[3:]...) // 其实这就是利用append的特性修改了切片内容再返回
	fmt.Println(c3)                // [30 31 33 34 35 36 37]
}
