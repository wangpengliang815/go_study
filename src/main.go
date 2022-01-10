package main

import "fmt"

func main() {
	var a calculation = Sum2
	fmt.Printf("type of c:%T\n", a) // type of c:main.calculation
	fmt.Println(a(1, 2))            // 像调用sum2一样调用c

	var b calculation = Sub2
	fmt.Printf("type of c:%T\n", b) // type of c:main.calculation
	fmt.Println(b(1, 2))            // 像调用sub2一样调用c
}
