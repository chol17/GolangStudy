package main

import "fmt"

func main() {
	var num1 int = 4
	var num2 int = 1

	if num1 > 2 {
		if num2 > 2 {
			fmt.Printf("num1+num2=%v \n", num1+num2)
		}
		fmt.Printf("run here\n")
	} else {
		fmt.Printf("num1-num2=%v\n", num1-num2)
	}

	if num1+num2 > 4 {
		fmt.Printf("num1+num2的和为:%v\n", num1+num2)
	}

	if num1+num2 >= 4 && num1-num2 <= 1 {
		fmt.Printf("run in ...\n")
	}

	if !(num1+num2 == 0) {
		fmt.Printf("判断的结果为：%v\n", !(num1+num2 == 0))
	}

	if num1 > num2 || num1 < num2 {
		fmt.Printf("num1 和 num2 不相等\n")
	}

	var num = 10 // 求负值，直接在正数前面加个负号
	fmt.Printf("-m+10的值为：%v", -num+10)

	// 出票系统：根据淡旺季的月份和年龄，打印票价【考虑学生优先做】
	/*
		4~10 月份旺季：
			成人（18~60岁）：60
			儿童（<18岁）：半价
			老人（>60岁）：1/3
		淡季：
			成人：40
			其他：20
	*/
	var age1 int
	var price float64
	var month int
	var pay float64
	price = 200.48
	fmt.Printf("请输入年龄：")
	fmt.Scanln(&age1)
	fmt.Printf("请输入月份：")
	fmt.Scanln(&month)
	if month >= 4 && month <= 10 {
		if age1 < 18 {
			pay = price * 1 / 2
		} else if age1 <= 60 {
			pay = 60
		} else {
			pay = price * 1 / 3
		}
	} else {
		if age1 >= 18 && age1 < 60 {
			pay = 40
		} else {
			pay = 20
		}
	}
	fmt.Printf("最后需要支付的金额为：%v", pay)
}
