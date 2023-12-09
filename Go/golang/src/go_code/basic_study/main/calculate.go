package main

import "fmt"

func main() {
	// 计算运算符
	var num int = 10
	num1 := num + 10
	fmt.Printf("num1=%v \n", num1)

	num2 := num1 / 10
	num3 := float64(num2) / 3
	fmt.Printf("num2=%v,num3=%v \n", num2, num3)
	fmt.Printf("num3的数据类型为%T \n", num3)

	num3++
	//++num3 Golang 中只有 后++ 和 后-- ，没有前++ 和 前-- 的语法
	fmt.Printf("num3=%v\n", num3)

	num3--
	fmt.Printf("num3=%v\n", num3)

	// 自增 和 自减 只能操作当前变量，无法通过自增 或者 自减的同时赋值给其他变量
	//num4 := num3 --  // 自增、自减只能独立使用
	// 在判断时，也不能用自增、自减 去和其他数据做判断，只能先自增、自减后，再用变量去做判断
	//if num++ > 0 {
	//	fmt.Printf("ok")
	//}
	num++ // 先将条件中需要的变量自增、自减完成后，再放到条件判断中才是正确的写法
	if num >= 0 {
		fmt.Printf("ok")
	} else {
		fmt.Printf("error")
	}
}
