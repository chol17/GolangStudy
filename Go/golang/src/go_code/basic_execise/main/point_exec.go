package main

import "fmt"

func main() {
	// 写一个程序，获取一个int 变量 num 的地址，并显示到终端
	// 将num 的地址赋给指针ptr, 并通过ptr 去修改num 的值

	int_var := 100
	fmt.Printf("int_var 的地址为%v \n", &int_var)
	fmt.Println("int_var=", &int_var)
	//var ptr *int = &int_var
	ptr := &int_var
	*ptr = 200
	fmt.Printf("int_var=%v", int_var)
}
