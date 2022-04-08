// @title 《Go语言编程》- 使用使用time.Sleep阻塞
// @description
// @author wangpengliang
// @date 2022-03-31 17:25:12

package main

import (
	"fmt"
	"time"
)

// 使用time.Sleep阻塞main主线程
func main() {
	go func() {
		fmt.Println("hello")
	}()
	fmt.Println("end...") // 先输出:end... 再输出:hello
	time.Sleep(time.Second)
}
