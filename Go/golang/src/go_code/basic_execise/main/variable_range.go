package main

import "fmt"

// 变量作用域
// 全局变量不能使用类型推导的方式定义
//Name := "Tom"
var Name string = "jack"

var (
	num1 = 10
	num2 = 20
)

func test1() {
	num1 := 5 // 该变量的作用域是test1 函数，在函数内部定义了该变量，操作该变量不会影响全局变量
	fmt.Printf("num1的值为：%v\n", num1)

}

func test2() {
	num1++ // 函数内未定义此变量，直接使用，表示在操作全局变量，只要该函数一经调用，全局变量num1 的值就发生变化了
	fmt.Printf("num1=%v\n", num1)
}

func main() {
	fmt.Printf("num2的值为：%v\n", num2)
	test1()
	test2()
	fmt.Printf("num1的值为：%v\n", num1)
	fmt.Printf("Name的值为：%v\n", Name)
}
