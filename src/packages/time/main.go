package main

import (
	"fmt"
	"time"
)

func main() {
	getTime()
}

// 获取当前时间对象的年月日时分秒
func getTime() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
}

func timeBefore() {
	now := time.Now()
	fmt.Println(now.Before(now.Add(time.Hour)))
}

// 时间比较,equal会比较地点和时区信息
func timeEqual() {
	now := time.Now()
	fmt.Println(now.Equal(now.Add(time.Hour)))
}

// 求两个时间之差
func timeSub() {
	now := time.Now()
	value := now.Sub(now.Add(time.Hour))
	fmt.Println(value)
}

// 获取当前时间+指定时间后的时间
func timeAdd() {
	now := time.Now()
	value := now.Add(1) // 当前时间加1小时后的时间
	fmt.Println(value)
}

// 将时间戳转换为时间格式
func unixConverTime() {
	// 将时间戳转换为时间格式
	timestamp := time.Now().Unix() //时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp)

	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 获取时间戳
func getUnix() {
	// 获取时间戳
	timestamp1 := time.Now().Unix()     //时间戳
	timestamp2 := time.Now().UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

// 获取当前时间
func getTimeNow() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
