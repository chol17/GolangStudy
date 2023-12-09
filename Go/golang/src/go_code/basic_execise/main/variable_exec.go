package main

import "fmt"
import _ "unsafe" // 如果导包没有用到，但是又不想删掉，可以在包名前面加一个下划线 ，表示忽略掉

func main() {
	// 申明了一个变量并且指定了其数据类型， 无法将一个不同类型的值赋值给这个变量， 如果要赋值时，需要对类型进行强转
	var n1 int32 = 12
	var n2 int64
	var n3 int8
	//n2 = n1 + 20  // n1是int32类型 加上一个整数 后还是int32 类型， 但是n2 是int64 类型，无法将int32 类型的值交给int64 类型的变量
	//n3 = n1 + 20  // n1是int32类型 加上一个整数 后还是int32 类型， 但是n2 是int8 类型，无法将int32 类型的值交给int8 类型的变量
	n2 = int64(n1) + 20
	n3 = int8(n1) + 20
	fmt.Println("n2=", n2)
	fmt.Println("n3=", n3)

	n3 = int8(n1) + 127
	fmt.Println("n3=", n3) // int8 的范围是 -128~127 ，此时相加得到的结果会溢出范围，但不会报错
	//n3 = int8(n1) + 128  // int8的范围是 -128~127 ,直接加128 会报错，编译无法通过
	n3 = int8(n1 + 128)
	fmt.Println("n3=", n3)

}
