package main

import (
	"fmt"
)

type books struct {
	titile string
	author string
}

func (b books) displayBooks() {
	fmt.Println(b.author, b.titile)
}

func displayBooks(b books) {
	fmt.Println(b.author, b.titile)
}

func main() {
	b := books{"titile434", "zhangsan"}
	b.displayBooks()
	displayBooks(b)
}
