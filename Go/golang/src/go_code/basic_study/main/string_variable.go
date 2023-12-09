package main

import "fmt"

func main() {
	// Golang 的字符串使用UTF-8 编码表示Unicode 文本，这样Golang 统一使用utf-8 编码，不会存在乱码问题
	// 字符串一旦赋值了，字符串就不能修改了，在Go 中字符串是不可变的
	// 字符串的两种表示形式:1. 双引号，会识别转义字符； 2. 反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	// 字符串拼接 使用 + 号
	var name string = "xiaochaoxiong"
	fmt.Println("name=", name)
	//name[0] = "a" // 这里不能对name 字符串进行修改，因为字符串是不可变的

	string1 := `abc` // 使用反引号给字符串变量赋值,这样编码安全
	string2 := `"xiaochaoxiong 123456"`
	fmt.Println("string1=", string1)
	fmt.Println("string2=", string2)
	var str = string1 + string2
	fmt.Println("str = ", str)

	str += "hahahha123"
	fmt.Println("str = ", str)

	// 多次拼接，并且换行时， 当前行没有拼接完的需要在末尾留下 + 号， 因为golang中 每一行默认会带上分号，但是如果行尾是 + 号，系统会识别为当前操作还未结束，因此可以继续拼接
	var str2 = "a" + "b" + "c" + "d" +
		"e" + "f" + "g" +
		"h"
	fmt.Println("str2=", str2)
}
