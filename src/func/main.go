package main

import "fmt"

/*函数是执行特定任务的代码块。*/

func main() {
	fmt.Println(max(2, 4))
	x := 3
	fmt.Println("x =", x)
	x1 := add1(&x)
	fmt.Println("x+1 = ", x1)
	fmt.Println("x = ", x)
}

//形式参数：定义函数时，用于接收外部传入的数据，叫做形式参数，简称形参。
//实际参数：调用函数时，传给形参的实际的数据，叫做实际参数，简称实参。
func max(parnum1 int, parnum2 int) int {
	if parnum1 >= parnum2 {
		return parnum1
	} else {
		return parnum2
	}
}

//引用传递
func add1(a *int) int {
	*a = *a + 1
	return *a
}
