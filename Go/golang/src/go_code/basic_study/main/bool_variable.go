package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 布尔类型也叫bool 类型， 其数据只允许取值为 true 和 false
	// bool 类型占 1 个字节
	// 布尔类型适用于逻辑运算 一般用于程序的流程控制
	// 不可以用 0 或者 非0 的整数替代 false 和 true, 这一点和
	var bool_variable bool = false
	bool1 := true
	fmt.Println("bool_variable=", bool_variable)
	fmt.Println("bool1=", bool1)
	fmt.Println("布尔类型数据占用空间为:", unsafe.Sizeof(bool1))
	fmt.Printf("bool1的数据类型为:%T", bool1)

}
