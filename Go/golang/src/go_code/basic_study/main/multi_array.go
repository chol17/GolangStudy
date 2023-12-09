package main

import "fmt"

// 二维数组

func main() {
	// 二维数组的定义 数组变量名 := [m][n]数组类型
	// 其中 m 表示这个二维数组一共有m 个元素，每个元素都是一个数组
	// n 表示这个二维数组中各个一维数组有n 个元素
	multiArr := [5][2]int{} // [[0,0],[0,0],[0,0],[0,0],[0,0]]
	fmt.Printf("二维数组的值为：%v\n", multiArr)

	// 给二维数组赋值
	// 1. 如果直接给二维数组中的元素赋值，需要传递数组
	multiArr[0] = [2]int{3, 4}
	fmt.Printf("赋值后的二维数组值为：%v\n", multiArr)
	// 2. 如果给二维数组中的子数组赋值，需要通过下标找到子数组的元素
	multiArr[1][0] = 5
	fmt.Printf("修改二维数组的子数组元素后的值为：%v\n", multiArr)

	// 遍历二维数组
	for _, arr := range multiArr {
		for _, value := range arr {
			fmt.Printf("一维数组的各个元素值为：%v\n", value)
		}
	}

	// 二维数组的内存解析
	// 二维数组在内存中存储的实际是指向一维数组的指针，指针地址依然可以通过类型和一维数组的个数推到出来
	multiArray := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	// 二维数组的地址就是数组中首元素(首一维数组)的第一个元素的地址值
	fmt.Printf("二维数组的地址为：%p\n", &multiArray)
	fmt.Printf("二维数组中第一个一维数组的第一个元素的地址为：%p\n", &multiArray[0][0])

	// 二维数组的定义和一维数组一样，需要指定数组长度和类型
	var multiArr1 [3][2]int
	fmt.Printf("multiArr1的值为：%v\n", multiArr1)
	multiArr2 := [3][2]int{{1}}
	fmt.Printf("multiArr2的值为%v\n", multiArr2)
	multiArr3 := [...][3]int{{1, 2, 3}, {4, 5}}
	fmt.Printf("multiArr3的值为%v\n", multiArr3)
}
