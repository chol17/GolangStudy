package main

import (
	"fmt"
)

func main() {
	// for 循环基本语法
	// for 循环变量初始化;循环条件;循环变量迭代{
	//    循环操作（逻辑代码）
	//}
	// for 循环的使用注意事项和细节：循环条件是返回一个布尔值的表达式;
	// Golang 中没有while 和 do...while 语法

	//for 循环编写方式1
	for i := 0; i <= 3; i++ {
		fmt.Printf("这是个案例\n")
	}

	// for 循环方式编写2
	num := 0
	for num <= 3 {
		fmt.Printf("编写方式2\n")
		num++
	}

	num1 := 0
	for {
		fmt.Printf("死循环...\n")
		num1++
		if num1 >= 5 {
			break
		}
	}

	// for-range 遍历
	str := "hello_world"
	// 传统的遍历方式
	for i := 0; i < len(str); i++ {
		fmt.Printf("遍历的字符为：%c\n", str[i])
	}

	// 使用 for-range 方式遍历
	for index, value := range str {
		fmt.Printf("index:%d, value:%c\n", index, value)
	}

	// 在遍历字符串时，如果字符串中有中文，直接遍历会出问题，因为遍历是按照字节遍历，而一个中文由3个字节组成，所以直接遍历得不到预期结果
	// 此时可以将字符串转为切片，再遍历  []rune
	str1 := "中国你好"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("遍历的字符为：%c\n", str1[i]) // 传统变量带有中文的字符串，会导致结果乱码问题

	}

	str2 := "我爱中国"
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("变量的字符为：%c\n", str3[i]) // 转为切片后就能正常遍历
	}

	str4 := "我爱中国"
	//str2 := []rune(str1)
	for index, value := range str4 {
		fmt.Printf("indes:%d, value:%c\n", index, value) // 如果使用for-range 遍历带有中文的字符串，会自动跳过,可以正常遍历到结果，注意索引并不是连续的
	}

	str5 := "中国您好"
	str6 := []rune(str5)
	for index, value := range str6 {
		fmt.Printf("index:%d,value:%c\n", index, value) // 转为切片就能保证遍历出来的下标是连续的
	}

label1:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i的值为：%v\n", i)
		for j := 0; j <= i; j++ {
			fmt.Printf("j的值为:%v\n", j)
			if j == 1 {
				break label1 // 多层for 循环嵌套时，break 默认只跳出离它最近的那层for循环，但是其他for 循环还是可以正常执行，
				// 可以在 for 循环上面加一个标签， break 标签名，可以跳出指定for 循环， 注意：指定了标签就一定要用到，因为golang 语法要求如此
			}
		}
	}
	fmt.Printf("====================================================================")
label2:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i的值为：%v\n", i)
		for j := 0; j <= i; j++ {
			if j == 1 {
				continue label2 // 多层for 循环嵌套时，continue 默认只跳出离它最近的那层for循环，但是其他for 循环还是可以正常执行，
				// 可以在 for 循环上面加一个标签， continue 标签名，可以跳出指定for 循环， 注意：指定了标签就一定要用到，因为golang 语法要求如此
			}
			fmt.Printf("j的值为:%v\n", j)
		}
	}

	//goto label2 // goto 结合标签使用或者if 判断使用，可以直接跳到标签位置继续执行代码，一般不建议使用goto
	fmt.Println("---------------------------------------------------------------------")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i的值为：%v\n", i)
		for j := 0; j <= i; j++ {
			if j == 1 {
				return // 当执行到return 时，直接跳出所有的for 循环，一般return 在方法或者函数中使用
				// return 如果在普通函数中，则表示跳出该函数，即不再执行函数中return 后面的代码，也可以理解为终止函数
				// 如果return 是在main 函数中，则表示终止main 函数，也就是说终止程序
			}
			fmt.Printf("j的值为:%v\n", j)
		}
	}
}
