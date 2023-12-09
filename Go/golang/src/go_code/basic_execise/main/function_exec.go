package main

import "fmt"

type mySum func(int, int) int

func sum(n1, n2 int) int {
	return n1 + n2
}

func sum2(n1 int, n2 int) (sub int) {
	sub = n1 + n2
	return
}

func myFunc(mySumVar mySum, num1, num2 int) int {
	return mySumVar(num1, num2)
}

func swap(n1, n2 *int) (int, int) {
	// 请编写一个函数，可以交换n1 和 n2 的值
	n3 := *n2
	n2 = n1
	n1 = &n3
	fmt.Printf("交换后的结果为：n1=%v,n2=%v\n", *n1, *n2)
	return *n1, *n2
}

func main() {
	a := sum
	b := sum2
	fmt.Printf("myFunc_a=%v\n", myFunc(a, 1, 2))
	fmt.Printf("myFunc_b=%v\n", myFunc(b, 1, 2))

	var n1, n2 = 1, 2
	resN1, resN2 := swap(&n1, &n2)
	fmt.Printf("main()调用之后的结果为:n1=%v,n2=%v\n", resN1, resN2)
}
