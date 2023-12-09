package main

import "fmt"

func test_args(args int32) int32 {
	args += 1
	return args
}

func main() {
	// switch 语句用于不同条件执行不同动作，每一个case 分支都是唯一的，从上到下逐一测试，直到匹配到为止
	// 匹配项后面也不需要再加 break
	// switch 的执行流程是：先执行表达式，得到值，然后和case的表达式进行比较，如果相等，就匹配到，然后执行对应的case 的语句块，然后退出switch 控制
	// 如果switch 的表达式的是没有和任何的case 的表达式匹配成功，则执行default 的语句块，执行后退出switch 控制
	// Golang 的case 后的表达式可以有多个，使用逗号间隔

	//请编写一个程序，可以接收一个字符，比如：a,b,c,d,e,f,g  a表示星期一，b表示星期二... 根据用户的输入显示相应的信息，要求使用switch 语句完成
	fmt.Printf("请输入一个字符：a,b,c,d,e,f,g\n")
	var byte_ byte
	fmt.Scanf("%c", &byte_) // Scanf 会根据格式化字符串解析输入的参数， Scanln 从标准输入读取文本内容
	switch byte_ {
	case 'a':
		fmt.Printf("今天星期一")
	case 'b':
		fmt.Printf("今天星期二")
	case 'c':
		fmt.Printf("今天星期三")
	case 'd':
		fmt.Printf("今天星期四")
	case 'e':
		fmt.Printf("今天星期五")
	default:
		fmt.Printf("输入有误")
	}

	var num1 int32 = 10
	var num2 int8 = 19
	switch num1 {
	//case num2:   // switch 表达式得到的结果必须和case 表达式的结果类型一致，否则会报错
	//	fmt.Printf("num1和num2 相等，均为%v", num2)
	case int32(num2):
		fmt.Printf("num1和num2 相等，均为%v", num2)
	case test_args(int32(num2)):
		fmt.Printf("test 函数的返回结果为%v", test_args(int32(num2)))
	case 10, 19:
		fmt.Printf("和多个值进行匹配，匹配到的值为：%v", num1)
	//case 10,10:  // 注意：如果case 的表达式是常量，那么case 与case 之间的常量不能重复，且当前case 中的常量不能重复
	//	fmt.Printf("重复")
	default: // default 可以有 也可以没有
		fmt.Printf("执行默认值")
	}

	// switch 后也可以不带表达式，类似于 if...else 分支来使用
	// 注意：这样写的时候，case 表达式返回结果必须是bool 类型的结果
	var num3 int64 = 100
	switch {
	case num3 > 10:
		fmt.Printf("num3=%v", num3)
	case num3 > 20:
		fmt.Printf("num3=%v", num3)
	case num3 == 40:
		fmt.Printf("num3=%v", num3)
	default:
		fmt.Printf("判断不对哦...")
	}

	switch score := 100; { // 在switch 中定义一个变量，注意后面要用分号隔开
	case score >= 60:
		fmt.Printf("及格")
	case score >= 80:
		fmt.Printf("良好")
	case score >= 90:
		fmt.Printf("优秀")
	default:
		fmt.Printf("成绩不合格")
	}

	// switch 穿透 fallthrought, 默认只穿透一层， 如果要穿透多层，需要在每一层都加上fallthrought
	switch grade := 80; {
	case grade >= 60:
		fmt.Printf("成绩合格")
		fallthrough
	case grade >= 80:
		fmt.Printf("成绩良好")
		fallthrough
	case grade >= 90:
		fmt.Printf("成绩优秀")
	default:
		fmt.Printf("成绩未合格")
	}

	var num4 interface{} // interface 是接口，后面学习
	var num5 = 10.6
	num4 = num5
	switch i := num4.(type) { // 接口变量.(type) 可以获取到接口变量的数据类型
	case int32:
		fmt.Printf("%v的类型为int32", i)
	case float64:
		fmt.Printf("%v类型为float64", i)
	default:
		fmt.Printf("%v类型有误", i)
	}

}
