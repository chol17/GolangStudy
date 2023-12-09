package main

import (
	"fmt"
	"time"
)

//Golang 中的内置函数--时间篇

func main() {
	// 1.获取时间类型结构体
	//golang 中time.Time类型表示时间,这是一个结构体
	timeNow := time.Now()
	fmt.Printf("timeNow 的值为：%v, 类型为：%T\n", timeNow, timeNow) //time.Now() 得到的是时间类型结构体 time.Time， 因此可以使用结构体中的方法获取时间其他信息

	// 2.调用时间类型结构体中的方法，获取年月日时分秒
	year := timeNow.Year()
	fmt.Printf("year 的类型为%T\n", year)
	_month := timeNow.Month()
	fmt.Printf("_month 的类型为%T\n", _month)
	month := int(_month)
	day := timeNow.Day()
	hour := timeNow.Hour()
	minute := timeNow.Minute()
	second := timeNow.Second()

	// 3.格式化日期时间
	// 方式一：
	fmt.Printf("当前时间为：%v-%v-%v %v:%v:%v\n", year, month, day, hour, minute, second)      // 使用fmt.Printf 仅仅是格式化输出内容，没有返回值
	dateStr1 := fmt.Sprintf("%v-%v-%v %v:%v:%v", year, month, day, hour, minute, second) // 使用fmt.Sprintf 可以格式化输出一个变量并且有返回值
	fmt.Printf("dateStr1:%v\n", dateStr1)
	//方式二：使用时间变量结构体中的Format("2006-01-02 15:04:05")   固定写法，2006-01-02 15:04:05 是固定值，为了纪念Golang 语言想法的诞生
	dateStr2 := fmt.Sprintf(timeNow.Format("2006-01-02 15:04:05")) // 里面的年月日时分秒数字不可改变，其他格式可以根据实际情况调整
	fmt.Printf("dateStr2的值为：%v\n", dateStr2)
	// 通过Format() 分别获取年月日时分秒
	month_ := fmt.Sprintf(timeNow.Format("01"))
	fmt.Printf("月份为：%v\n", month_)

	// 4.时间常量，常量的作用是：在程序中可用于获取指定时间单位的时间
	// 时间常量有Hour、Minute、Second、Millisecond、Nanosecond
	// 转换关系：1Second = 1000Millisecond, 1Millisecond = 1000Nanosecond
	fmt.Printf("time.Second = %v\n", time.Second) // time.Second 的值为1s
	// 时间常量一般和时间等待/睡眠 一起使用， 后面只能乘整型变量，不可乘小数
	time.Sleep(time.Second * 2)

	// 5.获取当前unix时间戳 和 unixnano 时间戳 （作用是可以获取随机数字）
	unixTime := timeNow.Unix()
	fmt.Printf("unix 时间戳是：%v\n", unixTime)
	unixNanoTime := timeNow.UnixNano()
	fmt.Printf("unixnano 时间戳是：%v\n", unixNanoTime)
}
