package main

import (
	"fmt"
)

func main() {
	var sum int
	make := make(map[string]int)
	make["兴业"] = 19000
	make["工商"] = 13000
	make["琪琪"] = 18000
	make["鸡孩"] = 25000
	make["多多"] = 20000
	make["路云"] = 20000
	make["家里"] = 30000
	for _, v := range make {
		sum += v
	}
	fmt.Printf("总数：%d \n", sum)

	monthMoney := 20760

	var sum3 int
	var sum9 int
	for i := 1; i < 13; i++ {
		if i <= 3 {
			sum3 += 16680
			fmt.Printf("month=%d money=%d \n", i, 16680)
		} else {
			sum9 += monthMoney + 2283
			fmt.Printf("month=%d money=%d \n", i, monthMoney+2283)
		}
	}
	fmt.Println(sum3, sum9, sum3+sum9, 29000*12, 29000*12-(sum3+sum9))
}
