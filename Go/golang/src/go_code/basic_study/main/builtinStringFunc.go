package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Golang 中的内置函数--字符串篇
func main() {
	str1 := "hello"
	str2 := "北京"
	// 1.len 统计字符串的长度，按字节统计
	// Golang 中的编码统一为utf-8, 即ascii 的字符(字母和数字)占用一个字节，汉字占用3个字节
	fmt.Printf("str1 的长度为：%v\n", len(str1))
	fmt.Printf("str2 的长度为：%v\n", len(str2))

	// 2.字符串遍历，同时处理有中文的问题，使用切片处理   r := []rune(str)
	toRune := []rune(str2)
	for index, _ := range toRune {
		fmt.Printf("index=%v, value=%c\n", index, toRune[index])
	}
	// 3.使用strconv.Atoi(str) 会把字符串转成数字，如果转不成功，会返回异常，但是代码不会报错
	//str3 := "12"
	str3 := "hello"
	toInt, err := strconv.Atoi(str3) // strconv.Atoi()  会把转换异常情况返回，并且保证代码不报错，可以通过判断
	if err != nil {                  // 对于判断的时候，判断是否为空，需要用 nil, Golang 中nil 相当于 Python 中的None
		fmt.Printf("错误信息为：%v\n", err)
	} else {
		fmt.Printf("转换后的结果为%v, 转换后的类型为：%T\n", toInt, toInt)
	}
	// 4.使用strconv.Itoa(str) 会把数字转成字符串，返回结果就是数字转成字符串的结果
	str4 := 123
	toStr := strconv.Itoa(str4) // 注意：只有一个返回值
	fmt.Printf("转换后的结果为：%v, 类型为：%T\n", toStr, toStr)

	// 5.字符串转byte 切片，应用场景： 将字符串写入一个文件中，就需要将字符串转成byte 数据类型
	var bytes = []byte("hello Go") // 将字符串中每一个字符转成对应的byte整数，并以一个数组类型返回
	fmt.Printf("bytes=%v，转换后的类型为：%T\n", bytes, bytes)

	// 6.将byte数组、切片转成字符串，  str := string([]byte{97,98,99})
	str5 := string([]byte{97, 98, 99}) // 将byte 类型的整数97,98,99 转成它对应的字符，并且拼接成一个字符串
	fmt.Printf("转换后的结果为：%v， 转换后的类型为：%T\n", str5, str5)

	// 7.将十进制整数转成二进制或者八进制字符串   str := strconv.FormatInt(123,2) 返回对应的字符串
	str6 := strconv.FormatInt(123, 2)
	fmt.Printf("转换后的结果为：%v, 转换后的类型为：%T\n", str6, str6)

	// 8.查找子串是否在指定的字符串中， strings.Contains("seafood", "foo"), 如果在，返回true, 反之返回false
	str7 := "seafood"
	is_sub := strings.Contains(str7, "foo") // 查找str7 这个字符串中是否包含 "foo"
	fmt.Printf("查找的结果为：%v, 返回的数据类型为：%T\n", is_sub, is_sub)

	// 9. 统计一个字符串中有几个指定的字符串
	str8 := "seafood"
	num1 := strings.Count(str8, "s")
	fmt.Printf("统计的结果为：%v,返回的数据类型为:%T\n", num1, num1)

	// 10.比较两个字符串是否相等，
	//如果使用 == 比较，会区分大小写； 使用strings.EqualFold() 比较，不会区分大小写
	str9 := "CompareAbc"
	//str10 := "CompareAbc"
	str10 := "compareABc"
	if str9 == str10 {
		fmt.Printf("两者大小写一致，即完全一致")
	} else if strings.EqualFold(str9, str10) {
		fmt.Printf("两者内容一致，但是有大小写区分")
	} else {
		fmt.Printf("两者不一样...")
	}

	// 11.返回子串在字符串中第一次出现的index值，如果没有返回-1
	// 作用：查询子串在字符串中的下标；通过返回值，判断子串是否在字符串中
	str11 := "Student"
	index := strings.Index(str11, "tu") // 字符串的下标从0 开始
	fmt.Printf("第一次出现的下标为：%v\n", index)

	// 12.返回子串在字符串中最后一次出现的index 值
	str12 := "Go Golang"
	lastIndex := strings.LastIndex(str12, "Go")
	fmt.Printf("最后一次出现的下标为：%v\n", lastIndex)

	// 13.将指定的子串替换成另一个子串
	str13 := "Study Golang"
	replaceStr := strings.Replace(str13, "Golang", "Go 语言", -1) // 需要指定替换多少个， n=-1 表示全量替换
	fmt.Printf("替换后的结果为：%v\n", replaceStr)

	// 14.按照指定的某个字符，为分割标识，将一个字符串拆分为字符串数组
	str14 := "hello,world"
	strArr := strings.Split(str14, ",")
	for index, value := range strArr {
		fmt.Printf("index=%v,value=%v\n", index, value)
		fmt.Printf("通过下标获取的值为：%v\n", strArr[index])
	}
	fmt.Printf("str14的值为%v\n", str14) // 操作字符串，并不会影响原字符串的结果，说明操作的是值拷贝，而不是引用的拷贝

	// 15.将字符串的字母进行大小写转换
	str15 := "Nsfocus"
	toLowerStr := strings.ToLower(str15)
	fmt.Printf("字符串转成小写的结果为：%v\n", toLowerStr)
	fmt.Printf("str15的结果为:%v\n", str15)
	toUpperStr := strings.ToUpper(str15) // 操作字符串，并不会影响原字符串的结果，说明操作的是值拷贝，而不是引用的拷贝
	fmt.Printf("字符串转成大写的结果为：%v\n", toUpperStr)

	// 16.将字符串左右两边的空格去掉
	str16 := " nsfocus   "
	fmt.Printf("str16的值为：%v\n", str16)
	stripStr := strings.TrimSpace(str16) // strings.TrimSpace() 用于处理字符串左右的空格
	fmt.Printf("去掉空格后的结果为：%v\n", stripStr)

	// 17.将字符串左右两边指定的字符去掉
	str17 := "!=nsfocus=!"
	fmt.Printf("str17的值为：%v\n", str17)
	trimStr := strings.Trim(str17, "!=")
	fmt.Printf("去掉指定字符后的结果为:%v\n", trimStr)

	// 18.将字符串左边指定字符去掉
	str18 := "!=nsfocus=!"
	fmt.Printf("str18的值为:%v\n", str18)
	trimLeftStr := strings.TrimLeft(str18, "!=")
	fmt.Printf("去掉左边的左边字符后的结果为：%v\n", trimLeftStr)

	// 19.将字符串右边指定字符去掉
	str19 := "!=nsfocus=!"
	fmt.Printf("str19 的值为：%v\n", str19)
	trimRightStr := strings.TrimRight(str19, "=!")
	fmt.Printf("去掉字符串右边字符后的结果为：%v\n", trimRightStr)

	// 20. 判断字符串是否以指定字符串开头、结尾
	str20 := "https:10.1.1.1:8443"
	fmt.Printf("str20 的值为：%v\n", str20)
	isStart := strings.HasPrefix(str20, "h") // 判断字符串以...开头，用 strings.HasPrefix
	isEnd := strings.HasSuffix(str20, "4")   // 判断字符串以...结尾，用strings.HasSuffix
	fmt.Printf("isStart的值为：%v, isEnd的值为：%v\n", isStart, isEnd)
}
