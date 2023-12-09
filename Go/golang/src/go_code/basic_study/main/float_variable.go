package main

import (
	"fmt"
)

func main() {
	var price float32 = 67.11
	fmt.Println("price 的值为=====", price)

	var num1 float32 = 3.14159253211
	var num2 float64 = 3.14159265311
	num3 := 1.34
	fmt.Println("num1的值为====", num1)
	fmt.Printf("num1的数据类型为 %T====\r\n", num1)
	fmt.Println("num2的值为====", num2)
	fmt.Printf("num2的数据类型为 %T====\r\n", num2)
	fmt.Printf("浮点数据默认类型为 %T====\r\n", num3)
}
