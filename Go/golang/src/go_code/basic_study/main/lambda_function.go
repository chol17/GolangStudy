package main

import "fmt"

//匿名函数 && 闭包

//申明全局变量，并且变量的为匿名函数变量

var (
	VarFunc = func(n1, n2 int) int {
		return n1 * n2
	}
)

// upAddFunc就是一个闭包
// 对闭包的说明：
// 1.upAddFunc是一个函数，返回的数据类型时func(int) int
// 2.返回的是一个匿名函数，但是这个匿名函数引用到函数外的wrap 变量，因此这个匿名函数就和wrap 变量形成一个整体，构成闭包
// 3.当我们把这个闭包函数赋值给一个变量时，反复调用，就能实现闭包里的功能
// 4.我们要搞清楚闭包的关键，就是要分析出返回的函数使用到了哪些变量，因为函数和它引用到的变量构成闭包
// 5.清楚了构成闭包的引用变量，那么该变量只会初始化一次，后续再反复调用时，只会在原来的基础上做操作，闭包最里面返回的是构成闭包的变量
func upAddFunc() func(int) int {
	wrap := 10
	return func(n int) int {
		wrap += n
		return wrap
	}
}

func main() {
	// 在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	// func(形参变量)(匿名函数返回变量){匿名函数逻辑}(参数) ：匿名函数的func 后面没有函数名，函数后面加上括号直接调用匿名函数
	result1 := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Printf("匿名函数执行的结果为:%v\n", result1)

	// 将匿名函数赋给一个变量（函数变量）, 再通过该变量来调用匿名函数
	funcVar := func(num1, num2 int) (sum int) {
		sum = num1 + num2
		return
	}
	result2 := funcVar(3, 7)
	fmt.Printf("通过变量调用得到的结果为:%v\n", result2)

	result3 := VarFunc(2, 9)
	fmt.Printf("在main 函数中调用全局匿名函数结果为:%v\n", result3)

	ua := upAddFunc()
	add1 := ua(1)
	fmt.Printf("调用一次的结果为：%v\n", add1)
	add2 := ua(2)
	fmt.Printf("调用一次的结果为：%v\n", add2)
}
