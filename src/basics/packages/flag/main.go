package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	flagTypeVar()
	// 返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	// 返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	// 返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

// 最简单的获取命令行参数
func getOsArgs() {
	// os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
	/*
		PS D:\MyProjects\go_study\src\commonPackages\flag> go run .\main.go "a" "sa" "wpl1212"
		args[0]=C:\Users\Administrator\AppData\Local\Temp\go-build2994373716\b001\exe\main.exe
		args[1]=a
		args[2]=sa
		args[3]=wpl1212
	*/
}

// 使用flagType定义参数
func flagType() {
	// 定义命令行参数；均为对应类型的指针
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("m", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")

	// 解析命令行参数
	flag.Parse()
	fmt.Println(*name, *age, *married, *delay) // 张三 18 false 0s
}

// 使用flagTypeVar定义参数
func flagTypeVar() {
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "m", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	// 解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay) // 张三 18 false 0s
}
