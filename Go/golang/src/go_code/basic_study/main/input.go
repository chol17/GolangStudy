package main

import "fmt"

func main() {
	// 控制台输入
	var num int
	fmt.Printf("num=%v", num)
	fmt.Printf("请输入数字：")
	fmt.Scanln(&num) // 读取命令行中输入的一行的内容
	fmt.Printf("num=%v", num)

}
