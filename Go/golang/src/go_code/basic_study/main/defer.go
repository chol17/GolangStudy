package main

import "fmt"

func sumFunc(n1, n2 int) int {
	defer fmt.Printf("执行sumFunc 函数中的defer 语句\n")
	// 该语句会在函数调用之后执行，但是入栈的时候，并没有执行加加操作，因此随同语句一起拷贝到栈中的变量并没有加加，还是原来的值，
	//只是执行顺序变为函数执行后出栈执行,说明入栈时也对变量值进行拷贝
	defer fmt.Printf("n1=%v,n2=%v\n", n1, n2)
	n1++
	n2++
	return n1 + n2
}

// 当程序执行到defer 时，暂时不执行defer 申明的这行代码，而是将defer 申明的内容压入到独立的栈中（defer 栈）
// 当函数执行或者调用完成后，再从defer 栈中按照先入后出的方式出栈，执行defer 语句
// 在defer 将语句放入到栈时，也会将相关的值拷贝同时入栈
// defer 的作用是将函数执行完毕需要执行的操作压入栈中。等函数执行完毕后，依次执行defer 栈的内容，在资源释放时经常使用这种方式！！！
// 说明：
// 1.在golang 编程中，通常的做法是：创建资源后，比如打开文件、获取数据库连接或者锁资源时，可以执行defer file.Close()、defer connect.Close()
// 2.在defer 后，可以继续使用创建的资源
// 3.在函数执行完毕后，系统会依次从defer 栈中，以先进后出的方式，执行语句，关闭资源
func main() {
	defer fmt.Printf("执行main 函数中第一条defer 语句\n")
	defer fmt.Printf("执行main 函数中第二条defer 语句\n")
	res := sumFunc(10, 11)
	fmt.Printf("调用sumFunc的结果为%v\n", res)

}
