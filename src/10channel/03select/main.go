package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	t1 := make(chan int)
	t2 := make(chan int)
	t3 := make(chan int)

	go f1(t1)
	go f2(t2)
	go f3(t3)

	for {
		select {
		case r1 := <-t1:
			fmt.Printf("f1 invoke %d \n", r1)
		case r2 := <-t2:
			fmt.Printf("f2 invoke %d \n", r2)
		case r3 := <-t3:
			fmt.Printf("f3 invoke %d \n", r3)
		}
	}
}

func f1(ch chan int) *chan int {
	ch <- 10
	return &ch
}

func f2(ch chan int) *chan int {
	ch <- 20
	return &ch
}

func f3(ch chan int) *chan int {
	ch <- 30
	return &ch
}
