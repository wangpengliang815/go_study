package main

import (
	"fmt"
)

func main() {
	for a := 0; a < 10; a++ {
		if a == 5 {
			//break：跳出循环体。break语句用于在结束其正常执行之前突然终止for循环
			break
		}
		if a == 3 {
			//continue：跳出一次循环。continue语句用于跳过for循环的当前迭代。
			//在continue语句后面的for循环中的所有代码将不会在当前迭代中执行。循环将继续到下一个迭代。
			continue
		}
		fmt.Println(a)
	}

	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
