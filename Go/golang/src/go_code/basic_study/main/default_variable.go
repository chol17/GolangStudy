package main

import "fmt"

func main() {
	// int 类型的默认值为0 ，float 类型的默认值为0， 布尔类型的默认值为false, string 类型的默认值为 ""
	var int_num int
	var float_num float32
	var bool_value bool
	var str string
	// 格式化输出的时候， %d 表示输出整数，%f 表示输出浮点数, %v 表示按照原值输出
	fmt.Printf("int类型的默认值为%v, float类型的默认值为%v, 布尔类型的默认值为%v, 字符串类型的默认值为%v", int_num, float_num, bool_value, str)
}
