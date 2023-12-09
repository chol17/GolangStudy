package main

import (
	"fmt"
	"unsafe"
)

// 所谓全局变量，即定义在函数外面的变量； 定义在函数里面的变量为局部变量
// 申明全局变量方式1:
var global_var = "global variable"

// 申明全局变量方式2:
var (
	global_var2 = "zhangsan"
)

func main() {
	// 第一种申明方式: var 变量名 类型
	var i int                     // 申明变量
	fmt.Println("i 的值为=====>", i) // 申明的变量没有赋值，其值为默认值 int 类型的默认值是 0；string类型的默认值为空字符串; 浮点类型的默认值为0
	i = 10                        // 给变量赋值
	fmt.Println("i=", i)          // 输出变量， golang 中 如果定义了变量而不使用会报错

	// 第二种申明方式: var 变量名 = 变量值   【变量类型自动推导，变量类型根据赋值的类型自动确定】
	var num = 10 // 默认为 int 类型
	fmt.Println("num 的值为=====>", num)

	// 第三种申明方式: 变量名 := 变量值
	name := "xiaochaoxiong"
	fmt.Println("num1的值为=====>", name)

	// 一次性申明多个变量方式1
	var num1, num2, num3 int
	fmt.Println("三个变量的值为====>", num1, num2, num3)

	// 一次性申明多个变量方式2
	var name1, name2, name3 = "TOM", "JACK", "MAKE"
	fmt.Println("三个名字为=====>", name1, name2, name3)

	// 一次性申明多个变量方式3
	gender1, gender2, gender3 := 1, 2, 3
	fmt.Println("三个性别分别为====>", gender1, gender2, gender3)

	// 在函数里输出全局变量
	fmt.Println("全局变量1的值为=====>", global_var)
	fmt.Println("全局变量2的值为=====>", global_var2)

	// + 号使用, 不同的变量之间不能使用 + 号 操作
	var int_num = 10.3
	var str_num = 10.1
	fmt.Println("两者相加为====>", int_num+str_num)

	var a int = 8900
	fmt.Println("a=", a)
	var b uint = 1
	fmt.Println("b=", b)

	var c byte = 255
	fmt.Println("c=", c)

	var d byte = 'd'
	fmt.Println("d=", d)

	// 查看变量的数据类型
	fmt.Printf("d 的类型为：%T \r\n", d)

	// 查看变量占用的字节数
	fmt.Printf("变量d 占用的字节数为：%d \r\n", unsafe.Sizeof(d))
}
