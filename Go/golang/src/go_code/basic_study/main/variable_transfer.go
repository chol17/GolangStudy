package main

import "fmt"

func main() {

	// 类型范围小 转 类型范围大， 可正常转换数据
	var num1 int32 = 300
	var num2 int64 = int64(num1)
	var num3 float32 = float32(num1) // 类型转换时，被转换的数据类型不会发生变化，只是将数值转换成指定类型后重新赋值给新的变量
	fmt.Printf("num2=%v , num3=%v \r\n", num2, num3)
	fmt.Printf("num1 的数据类型为%T \r\n", num1)

	// 类型范围大 转 类型范围小， 可能存在数据范围溢出问题,但是不会报错
	var num4 int8 = int8(num1)
	fmt.Println("num4=", num4)
	fmt.Printf("num4=%v  \r\n", num4)
	fmt.Printf("num4的数据类型为%T \r\n", num4)
}
