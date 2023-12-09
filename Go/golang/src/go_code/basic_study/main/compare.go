package main

import "fmt"

func main() {
	// 关系运算符，比较运算符
	// == != >  >=  < <=
	// 关系运算符的结果都是bool 类型，要么为true 要么为false
	num1 := 7
	num2 := 8
	flag1 := num1 == num2
	flag2 := num1 >= num2
	flag3 := num1 <= num2
	fmt.Printf("flag1=%v,flag2=%v,flag3=%v \n", flag1, flag2, flag3)
	fmt.Printf("flag1的类型为%T,flag2的类型为%T,flag3的类型为%T \n", flag1, flag2, flag3)
}
