package main

import (
	"fmt"
)

func main() {

	// 条件控制
	// 编写一个程序，可以输入人的年龄，如果年龄大于18岁，则输出"你年龄大于18，要对自己的行为负责"
	var age int
	fmt.Scanln(&age)
	if int(age) >= 18 {
		fmt.Printf("你年龄大于18，要对自己的行为负责\n")
	} else {
		fmt.Printf("你不需要要对自己的行为负责\n")
	}
	// Golang 中，在条件判断时定义变量，并对变量进行判断
	if num := 20; num >= 18 {
		fmt.Printf("你年龄大于20，要对自己的行为负责\n")
	} else {
		fmt.Printf("你不需要要对自己的行为负责\n")
	}

	// 多分支
	num1 := 30
	if num1 > 60 {
		fmt.Printf("你已经是老年人了\n")
	} else if num1 > 30 {
		fmt.Printf("你已经是中年人\n")
	} else if num1 > 18 {
		fmt.Printf("你已经成年\n")
	} else {
		fmt.Printf("你年龄还小\n")
	}

	// 注意语法， if...else if ... else ... 中的条件表达式不能是赋值语句，只能是为能判断结果的条件表达式
	//var num2 = true
	//if num2 = false{   // 条件表达式的位置写赋值语句，会导致编译不通过
	//	fmt.Printf("num2 = %v",num2)
	//}

}
