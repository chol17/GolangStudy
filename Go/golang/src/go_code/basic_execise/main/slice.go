package main

import "fmt"

func fbnFunc(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	// 使用切片完成斐波那契数列
	// 斐波那契数的规律：第一个元素为1，第二个元素为1，从第三个元素开始，每个元素的值为前两个元素的和
	num := 50
	fbnSlice := fbnFunc(num)
	fmt.Printf("前%v个斐波那契数为：%v\n", num, fbnSlice)
}
