package main

import "fmt"

func main() {
	// Golang 中没有三元运算符
	// 逻辑运算符： &&  ||  !
	//  && 逻辑与，如果两边的操作数都是true, 则为true,否则为false
	//  || 逻辑或，如果两边的操作数有一个是true, 则为true, 否则为 false
	//  !  逻辑非，如果条件为true,则逻辑运算为 false, 反之为 true
	num1 := 40
	num2 := 50
	if num1 >= 20 && num2 <= 10 {
		fmt.Printf("ok\n")
	} else {
		fmt.Printf("No\n")
	}

	if num1 >= 20 || num2 <= 10 {
		fmt.Printf("ok\n")
	} else {
		fmt.Printf("No\n")
	}

	if !(num1 > 20) {
		fmt.Printf("run here...\n")
	} else {
		fmt.Printf("no run...\n")
	}

	// 没有以下语法
	//if num1 & 10{
	//	fmt.Printf("逻辑与执行")
	//}
	// && 又叫做短路与， 如果第一个判断为true，还需要执行第二个判断才能最终判断出条件的结果
	// 但是如果第一个判断为false , 则 && 符号后面的判断不再执行
	if num1 >= 10 && test() { // 会输出 test...
		fmt.Printf("进入流程\n")
	} else {
		fmt.Printf("没有进入流程\n")
	}

	if num1 <= 10 && test() { // 不会输出 test...
		fmt.Printf("进入流程\n")
	} else {
		fmt.Printf("没有进入流程\n")
	}

	//  || 又叫短路或， 如果第一个判断为 true, 则 || 后面的判断不再执行
	// 如果第一个条件为false , 则后面的判断还会继续执行
	if num1 > 10 || test() { // && 前面为true, 后面不再执行
		fmt.Printf("run in ...\n")
	} else {
		fmt.Printf("not run in ...\n")
	}

	if num1 <= 10 || test() { // && 前面为 false,则后面还会继续执行
		fmt.Printf("running...\n")
	} else {
		fmt.Printf("not running...\n")
	}
}

// 申明一个函数
func test() bool {
	fmt.Printf("test...\n")
	return true
}
