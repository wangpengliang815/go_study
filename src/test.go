package main

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func main() {
	a := []int{3, 4, 1, -1}
	fmt.Println(solution(a))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func solution(a []int) int {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] <= 0 {
			a[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		num := abs(a[i])
		if num <= n {
			a[num-1] = -abs(a[num-1])
		}
	}

	for i := 0; i < n; i++ {
		if a[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}
