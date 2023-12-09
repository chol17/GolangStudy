package main

import (
	"fmt"
	_ "fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 23.11
	var char_ byte = 'a'
	var bool_var bool = true
	var str string
	// 转换方式1  str = fmt.Sprint 方法, 将基本数据类型数据转换为字符串类型
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str=%q \n", str)
	fmt.Printf("str 的数据类型为%T \n", str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str=%q \n", str)
	fmt.Printf("str 的数据类型为%T \n", str)

	str = fmt.Sprintf("%t", bool_var)
	fmt.Printf("str=%q \n", str)
	fmt.Printf("bool_var=%v \n", bool_var)

	str = fmt.Sprintf("%c", char_)
	fmt.Printf("%q \n", str)
	fmt.Printf("char_的数据类型为%T \n", char_)
	fmt.Printf("str 的数据类型为%T \n", str)

	// 转换方式2  strconv 函数
	var int_var = 100
	var float_var float32 = 10.33
	//var char_var byte = 's'
	var boolean_var bool = false
	var str_var string

	// strconv.FormatInt 必传传参是2个， 第一个参数为 int64 类型的整型数据， 第二个是指明转成多少进制的数
	str_var = strconv.FormatInt(int64(int_var), 10)
	fmt.Printf("str_var=%q \n", str_var)

	// strconv.FormatFloat 必传参是4个，第一个参数为float64 类型的浮点数，第二个参数是转换格式为'f', 第三个参数为小数保留位数（精度）， 第四个参数为小数类型为float64
	str_var = strconv.FormatFloat(float64(float_var), 'f', 2, 64)
	fmt.Printf("str_var=%q \n", str_var)

	// strconv.FormatBool 传一个参数， 需要转换的bool 类型变量值
	str_var = strconv.FormatBool(boolean_var)
	fmt.Printf("str_var=%q \n", str_var)

	// 转换方式3 strconv 包中有一个函数 Itoa
	var number = 100
	itoa_str := strconv.Itoa(number)
	fmt.Printf("itoa_str=%q \n", itoa_str)

	// 如果待转换的数值不是int64, 调用itoa 时需要用int转换一下
	var number1 int8 = 12
	itoa_str1 := strconv.Itoa(int(number1))
	fmt.Printf("itoa_str1=%q", itoa_str1)
}
