package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 输出字符串时，可以用 %q 格式化接收
	// 输出其他类型时，建议使用 %v 格式化接收
	var str string = "true"
	var str_number = "123"
	var bool_var bool
	var int_number int64
	var int_num int
	bool_var, _ = strconv.ParseBool(str)
	fmt.Printf("bool_var=%v \n", bool_var)
	fmt.Printf("bool_var 的数据类型为%T \n", bool_var)

	// strconv.ParseInt 转换后返回的int 是 int64 类型，因此必须使用int64 类型的变量接收
	int_number, _ = strconv.ParseInt(str_number, 10, 64)
	fmt.Printf("int_number=%v\n", int_number)

	int_num = int(int_number)
	fmt.Printf("int_num=%v\n", int_num)

	// strconv.ParseFloat 转换后返回的是float 是float64 类型， 因此必须使用float64 类型的变量接收
	var str_float = "123.11"
	//var float_num float64
	float_num, _ := strconv.ParseFloat(str_float, 64)
	fmt.Printf("float_num=%v \n", float_num)
	fmt.Printf("float_num的数据类型为%T \n", float_num)

	// 字符串转数值类型时，如果没有转成功，则转换后得到的结果为0
	// 字符串转布尔类型时，如果没有转成功，则转换后得到的结果为false
	string_ := "hello"
	//var to_int int64
	to_int, _ := strconv.ParseInt(string_, 10, 64)
	fmt.Printf("to_int=%v \n", to_int)

	to_bool, _ := strconv.ParseBool(string_)
	fmt.Printf("to_bool=%v \n", to_bool)
	fmt.Printf("to_bool的数据类型为%T \n", to_bool)

	// 方式2：字符串类型转int 使用 strconv.Atoi
	to_int2, _ := strconv.Atoi(string_)
	fmt.Printf("to_int2=%v \n", to_int2)
	fmt.Printf("to_int2的数据类型为%T \n", to_int2)
}
