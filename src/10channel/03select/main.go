package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var chs [2]chan int
	chs[0] = make(chan int, 10)
	chs[1] = make(chan int, 10)

	go func() {
		defer wg.Done()
		for v := range chs[0] {
			fmt.Println(v)
		}
	}()

	go func() {
		defer wg.Done()
		for v := range chs[1] {
			fmt.Println(v)
		}
	}()

	for i := 0; i < len(chs); i++ {
		chs[i] <- 100
		wg.Add(1)
		fmt.Printf("chs[%v]写入100完成 \n", i)
	}
	wg.Wait()
}
