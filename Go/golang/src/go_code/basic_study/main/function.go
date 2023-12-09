package main

import (
	"fmt"
	"strconv"
)

// 函数注意事项和细节：
// 1.函数的形参列表可以是多个，返回值列表也可以是多个
// 2.形参列表和返回值列表的数据类型可以是值类型和引用类型
// 3.函数的命名遵守标识符命名规范，首字母不能是数字，首字母大写表示该函数可以被本包文件和其他包文件使用，类似pubilc，首字母小写，只能被本包文件使用，其他包文件不能使用，类似private
// 4.函数中的变量时局部的，函数外不生效
// 5.基本数据类型和数组类型默认都是值传递的，即进行值拷贝，在函数内修改，不会影响到原来的值
// 6.如果希望函数内的变量能修改函数外的变量，可以传入变量的地址 &变量，函数内以指针的方式操作变量，从效果上看类似引用
// 7.GO 函数不支持重载
// 8.在Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
// 9.函数既然是一种数据类型，因此在Go 中，函数可以作为形参，并且调用
// 10.为了简化数据类型的定义，Go支持自定义数据类型
// 基本语法： type 自定义数据类型名 数据类型     // 理解：相当于一个别名
// 11.支持对函数返回值命名
// 12.使用 _ 标识符，忽略返回值
// 13.Go 中支持可变参数

func calculate(num1 int, num2 int, sign byte) int {
	var result int
	//flag := true
	switch {
	case sign == '+':
		result = num1 + num2
	case sign == '-':
		result = num1 - num2
	case sign == '*':
		result = num1 * num2
	default:
		result = 0
		//flag = false
	}
	fmt.Printf("result=%v\n", result)
	return result
}

func multi(varchar string) (int, string) {
	var res int
	switch varchar {
	case "1":
		res, _ = strconv.Atoi(varchar)
	case "2":
		res, _ = strconv.Atoi(varchar)
	default:
		res = 0
	}
	ignore_str := varchar
	// 如果返回多个值，在接收时，希望忽略某个返回值，则使用 _ 符号表示占位忽略
	// 如果返回值只有一个，返回值类型雷彪可以不写()
	return res, ignore_str
}

func pointer(addr *int) int {
	// &变量名：表示寻址，获取变量在内存中的地址值
	// *指针变量：表示取值，获取指针变量指向的地址对应的变量值
	fmt.Printf("传进来的变量值为:%v\n", addr)
	fmt.Printf("*addr值为:%v\n", *addr)
	*addr += 10
	return *addr
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func getMulti(getSum func(int, int) int, m1 int, m2 int) int { // 在函数的形参变量中，设置一个函数类型的形参变量， 格式为 形参变量变量名 func(函数类型形参需要的参数) 函数类型的变量返回值类型
	return getSum(m1, m2)
}

type myFunc func(int, int) int // 给函数类型起一个别名
func getMulti2(getSum myFunc, num1 int, num2 int) int { //将函数类型别名作为形参传到其他函数中
	return getSum(num1, num2)
}

func getSumAndSub(num1 int, num2 int) (sub int, sum int) { // Go 中支持对返回值起别名
	sub = num1 - num2 // 因为已经在函数的返回参数中定义了变量名和类型，因为在函数中可以直接使用返回的变量
	sum = num1 + num2
	return // 返回时也不需要带上返回变量，默认按函数定义里面的返回形参顺序返回结果
}

//Go 中支持可变参数
func multiParams(args ...int) (sum int) {
	// args 是slice 切片，通过args[index] 可以访问到各个值
	// 如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后
	for index, value := range args {
		arg := args[index]
		sum += arg
		fmt.Printf("arg=%v\n", arg)
		fmt.Printf("index=%v, value=%v\n", index, value)
	}
	return
}

func main() {
	// 函数定义的格式：
	// func 函数名(形参列表)(返回值类型列表){
	//		执行语句...
	//      return 返回值列表
	//}
	// 1.形参列表，表示函数的输入，形参名 形参类型，多个传参之间用逗号隔开
	// 2.函数中的语句: 表示为了实现某一功能的代码块
	// 3.函数可以有返回值，也可以没有返回值，单个返回值可以不需要括号括起来，多个返回值需要用括号括起来，并且多个之间用逗号隔开
	var num1 = 3
	var num2 = 4
	var sign byte = '*'
	//var flag bool
	//var result int
	result := calculate(num1, num2, sign)
	str := "1"
	multi_int, _ := multi(str)
	fmt.Printf("multi func 的结果为:%v\n", multi_int)
	fmt.Printf("result的结果为:%v\n", result)
	addr := 10
	addr_res := pointer(&addr)
	fmt.Printf("通过指针引用传递参数后的结果为:%v\n", addr_res)

	func_var := getSum // 将一个函数赋值给一个变量，那么这个变量的数据类型就是函数类型
	fmt.Printf("func_var 的数据类型为%T\n", func_var)
	res := func_var(10, 20)
	fmt.Println("res=", res)

	multi_res := getMulti(func_var, 10, 10)
	fmt.Printf("multi_res的结果为%v\n", multi_res)

	type myInt int        // 给int 类型取别名为 myInt
	var myNum1 myInt = 10 // 定义一个myInt 类型的变量
	var myNum2 int
	fmt.Printf("myNum2的数据类型为%T\n", myNum2)
	fmt.Printf("myNum的值为%v\n", myNum1)
	//myNum2 = myNum1  // myNum2 是int 类型，但是无法将myNum1的值赋给myNum2, 是因为在Go 中，认为起了别名后属于两个不同的数据类型，因为需要显示转换后才能赋值
	myNum2 = int(myNum1)
	fmt.Printf("myNum2的数据类型为%T\n", myNum2)

	getMultiResult := getMulti2(getSum, 100, 20)
	fmt.Println("getMultiResult=", getMultiResult)

	sub, sum := getSumAndSub(13, 17)
	fmt.Printf("sub 的值为%v, sum的值为%v\n", sub, sum)

	multiParams := multiParams(2, 3, 4)
	fmt.Printf("multiParams=%v\n", multiParams)
}
