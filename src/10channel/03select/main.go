package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var ch1 = make(chan int, 10)
	// var ch2 = make(chan bool, 10)

	go func(chan int) {
		defer wg.Done()
		for v := range ch1 {
			fmt.Println(v)
		}
	}(ch1)

	// go func(chan bool) {
	// 	defer wg.Done()
	// 	for v := range ch2 {
	// 		fmt.Println(v)
	// 	}
	// }(ch2)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		ch1 <- i
	}
	close(ch1)

	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)
	// 	ch2 <- false
	// }
	// close(ch2)

	wg.Wait()
}
