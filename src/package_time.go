package main

import (
	"fmt"
	"time"
)

func TimeTest() {
	// func Now() Time
	fmt.Println(time.Now())

	// func Parse(layout, value string) (Time, error)
	time1, _ := time.Parse("2016-01-02 15:04:05", "2018-04-23 12:24:51")
	fmt.Println(time1)

	// func ParseInLocation(layout, value string, loc *Location) (Time, error) (layout已带时区时可直接用Parse)
	time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local)

	// func Unix(sec int64, nsec int64) Time
	time.Unix(1e9, 0)

	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	time.Date(2018, 1, 2, 15, 30, 10, 0, time.Local)

	// func (t Time) In(loc *Location) Time 当前时间对应指定时区的时间
	loc, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(time.Now().In(loc))

	// func (t Time) Local() Time
	fmt.Println(time.Local)
}

func TimeBefore() {
	now := time.Now()
	fmt.Println(now.Before(now.Add(time.Hour)))
}

// 时间比较,equal会比较地点和时区信息
func TimeEqual() {
	now := time.Now()
	fmt.Println(now.Equal(now.Add(time.Hour)))
}

// 求两个时间之差
func TimeSub() {
	now := time.Now()
	value := now.Sub(now.Add(time.Hour))
	fmt.Println(value)
}

// 获取当前时间+指定时间后的时间
func TimeAdd() {
	now := time.Now()
	value := now.Add(1) // 当前时间加1小时后的时间
	fmt.Println(value)
}

// 将时间戳转换为时间格式
func UnixConverTime() {
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
func GetUnix() {
	// 获取时间戳
	timestamp1 := time.Now().Unix()     //时间戳
	timestamp2 := time.Now().UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

// 获取当前时间
func GetTimeNow() {
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
