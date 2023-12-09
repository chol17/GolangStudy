package main

import (
	"fmt"
	util "golang/src/go_code/basic_study/utils"
)

//init 函数的注意事项和细节
// 1.如果一个文件同时包含全局变量定义，init函数和 main 函数，则执行的流程是：调用包中的init函数--> 全局变量定义--> init函数--> main函数
// 2.init 函数最主要的作用，就是完成一些初始化的工作

var age = test_age() // 全局变量会在init 函数调用之前
var num = util.Number

func test_age() int {
	fmt.Printf("test function\n")
	return 18
}

//每一个源文件都可以包含一个init 函数，该函数会在main 函数执行前，被Go运行框架调用，也就是说init 会在main 函数前被调用
func init() {
	fmt.Printf("init function\n")
}

func main() {
	fmt.Printf("main function\n")
	fmt.Printf("age=%v\n", age)
	fmt.Printf("util.num=%v\n", num)
}
