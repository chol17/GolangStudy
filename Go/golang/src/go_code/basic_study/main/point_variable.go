package main

import "fmt"

func main() {
	// 指针类型: 指针变量存储的是一个地址，这个地址指向的空间存的才是值
	// 指针类型的变量存储的值是一个地址，这个地址指向的是该地址对应的值，使用 *指针变量名 能获取到该指针变量指向的地址对应的值
	// 非指针变量，通过 &变量名 能获取到该变量对应的地址
	// int 类型的变量地址只能赋值给 int类型指针
	// 指针的值类型形式为 *数据类型， eg. *int  *float32
	// 值类型包括：基本数据类型 int float bool string 数组 结构体
	var num = 10
	var point_addr *int = &num
	fmt.Printf("point_addr=%v \n", point_addr)
	fmt.Printf("point_addr的类型为%T \n", point_addr)
	fmt.Printf("point_addr的地址为%v \n", *point_addr)
	addr := &num
	fmt.Printf("addr=%v\n", addr)
	fmt.Printf("addr的地址为%v\n", &addr)
	fmt.Printf("addr的类型为%T\n", addr)

	var num1 float32 = 30
	//var int_ptr *int = &num1  // float32 类型变量的地址不能赋给 int 类型指针
	var float_ptr *float32 = &num1
	fmt.Printf("float_ptr=%v \n", float_ptr)
	fmt.Printf("float_ptr指向的值为%v \n", *float_ptr)
	//var float_ptr2 *float64 = &num1  // 不同类型之间无法赋值

	var int_num1 int = 10
	var int_num2 int = 20
	var ptr_int *int = &int_num1
	*ptr_int = 40
	ptr_int = &int_num2
	*ptr_int = 90
	fmt.Printf("int_num1=%v, int_num2=%v,ptr_int=%v,*ptr_int=%v", int_num1, int_num2, ptr_int, *ptr_int)
}
