// @title 《Go语言编程》-错误处理
// @description
// @author wangpengliang
// @date 2022-03-25 11:23:23
package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := sum(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sum(a, b int) (n int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("error")
	}
	return a + b, nil
}
