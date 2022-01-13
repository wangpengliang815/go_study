package main

import "fmt"

func main() {
	var f = Add()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60
}
