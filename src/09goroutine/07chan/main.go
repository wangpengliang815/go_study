package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int)
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go send()
	// go receive()
	wg.Wait()
}

func send() {
	defer wg.Done()
	ch <- 10
	close(ch)
}

func receive() {
	defer wg.Done()
	x := <-ch
	fmt.Println(x)
}
