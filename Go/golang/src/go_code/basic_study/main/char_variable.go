package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Golang 中, 字符使用UTF-8 编码， 英文字母占用 1个字节，汉字占用3个字节；
	// 字符的本质是一个整数, 直接输出时，是该字符对应的utf-8 编码的码值， 可以使用%c 输出其unicode 字符
	// 字符类型是可以进行运算的， 相当于一个整数，因为它都有对应的unicode 码

	var c1 byte = 'a'
	// 当直接输出byte 类型的变量时，其实是输出了对应字符的字符码值
	fmt.Println("c的值为=====", c1)

	// 格式化输出，才能输出对应的字符
	fmt.Printf("c1 的字符值为=====%c  \r\n", c1)

	// 字符类型的变量用单引号，且单引号中只有一个字符
	var c2 int = '我'
	fmt.Printf("c2 的值为：%c \r\n", c2)
	fmt.Printf("c2 的数据类型为：%T \r\n", c2)
	fmt.Println("c2 占用字节数为：", unsafe.Sizeof(c2))

	//c4 := 'abc'
	//fmt.Println("c4 的值为=====",c4)

	var int_num = 10
	var char_ = 'a' + int_num
	fmt.Printf("两者相加的类型为 %T \r\n", char_)
	fmt.Println("两者相加直接输出结果为", char_)
	fmt.Printf("两者相加格式化输出的结果为 %c \r\n", char_)

	c3 := "你好呀"
	fmt.Printf("c3 的数据类型为=====%T \r\n", c3)
}
