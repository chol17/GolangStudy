package main

import "fmt"

func main() {

	// 假如还有97天放假，问：97天是多少个星期零多少天
	var total_day = 97
	week := total_day / 7
	day := total_day % 7
	fmt.Printf("97天是%v 个星期，零%v天\n", week, day)

	// 定义一个变量报错华氏温度，花式温度转换摄氏度的公式为：
	// 5/9*(华氏度-100)， 求出华氏温度对应的摄氏温度
	// int 和 float 类型之间不能进行运算
	var temp int = 200
	temp1 := float32(5) / 9 * (float32(temp) - 100)
	fmt.Printf("%v 华氏温度对应的摄氏温度为:%v\n", temp, temp1)

	// 将两个变量的值进行交换，但不允许使用中间变量，打印结果
	var num1 float32 = 10
	var num2 float32 = 20
	fmt.Printf("num1原来的值为%v, num2原来的值为%v\n", num1, num2)
	num1 += num2       // 30
	num2 = num1 - num2 // 10
	num1 = num1 - num2 // 20
	fmt.Printf("num1=%v,num2=%v\n", num1, num2)

	// 求三个数的最大值
	var int_var1 = 12
	var int_var2 = 15
	var int_var3 = 18
	if int_var1 > int_var2 && int_var1 > int_var3 {
		fmt.Printf("max value is %v", int_var1)
	}
}
