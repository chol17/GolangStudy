package main

// 数组
// 数组可以存放多个同一类型的数据，数组也是一种数据类型，在Golang 中，数组是值类型
// 数组变量指向的是数组的内存空间，而非是数组空间的地址，因此数组是值类型
// 数组定义 var 数组名[数组长度]数组类型

import "fmt"

func testArr(arr [3]int) {
	arr[1] = 100
}

func testArrPointer(arr *[3]int) {
	// 注意：数组以地址传递的书写方式为： 数组变量名 *[数组长度]数组类型
	// 寻址赋值的方式为：(*数组变量名)[下标] = 值
	(*arr)[0] = 13
	(*arr)[1] = 14
	(*arr)[2] = 15
}

func main() {
	// 整型数组，元素的默认值为0
	var numArray [6]int
	fmt.Printf("数组默认值为：%v\n", numArray)

	// 给数组各个元素赋值
	numArray[1] = 10 // 给第二个元素赋值
	fmt.Printf("赋值后的数组为：%v\n", numArray)

	// 获取数组的地址
	// 数组的地址可以通过 &数组变量名 获取
	arrayPointer := &numArray
	fmt.Printf("数组的地址为：%p\n", arrayPointer) // 通过%p 拿到地址值
	// 且数组的地址就是数组元素中第一个元素的地址值
	fmt.Printf("数组第一个元素的地址值为：%p\n", &numArray[0])
	// 依次类推，数组元素后续的元素地址值就是前一个元素的地址值 + 数组类型占用的字节数
	// int 类型占用8个字节，因此int 类型数组的下一个元素地址是在前一个元素的地址值上加8
	fmt.Printf("数组第二个元素的地址值为：%p\n", &numArray[1])

	// 数组元素访问： 数组名[下标]
	arrayValue := numArray[1]
	fmt.Printf("数组的值为：%v\n", arrayValue)

	// 初始化数组的4种方式
	newArray1 := [3]int{1, 2, 3} //使用类型推到的方式初始化数组
	fmt.Printf("数组newArray 的值为：%v\n", newArray1)

	// 使用省略号不指定数组元素个数的方式初始化数组
	// [...]数组类型{} 是固定语法
	newArray2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("数组newArray2 的值为：%v\n", newArray2)

	// 对不定长数组中指定元素赋值
	newArray3 := [...]int{1: 30, 2: 10, 4: 50}
	fmt.Printf("数组newArray3的值为：%v\n", newArray3)

	// 初始化字符串类型数组
	strArray1 := [...]string{"jack", "tom", "mack"}
	fmt.Printf("字符串类型的数组strArray1 的值为：%v\n", strArray1)

	// 字符串类型数据默认不赋值 默认为空字符串
	strArray2 := [3]string{}
	fmt.Printf("字符串数组默认值为：%v\n", strArray2)

	//数组的遍历
	// 方式1：使用传统的for 循环遍历
	forArray := [...]int{1, 2, 3, 4}
	for i := 0; i < len(forArray); i++ {
		fmt.Printf("数组的第%v 个元素的值为：%v\n", i+1, forArray[i])
	}
	//方式2：for-range 遍历数组, range 后面跟数组变量名
	// 1.第一个返回值index 是数组下标值
	// 2.第二个返回值value 是数组该下标对应的值
	// 3.他们都是在for 循环内部可见的变量
	// 4.遍历数组时，如果不想使用下标index,可以直接把下标index 换做下划线 _
	for index, value := range forArray {
		fmt.Printf("数组下标为%v的数值为%v\n", index, value)
	}
	//------------------------------------数组细节讨论---------------------------------
	// 1.数组是多个相同类型数据的组合，一个数组一旦申明/定义了，其长度是固定的，不能动态变化
	var numArr [3]int
	numArr[0] = 10
	numArr[1] = 11
	numArr[2] = 12
	//numArr[3] = 13  // 定义数组时已经确定了数组长度，因此无法再在赋值的时候赋予超出数组长度的数据,否则报越界的错
	//numArr[1] = "hello" // 数组在定义的时候已经声明了数组的类型为int类型，因此在赋值的时候只能赋予int类型的值
	fmt.Printf("numArr 数组的值为：%v\n", numArr)
	// 2.当第一个数组时，没指定数组长度大小时，它其实是一个切片类型
	var numSlice []int
	fmt.Printf("numSlice的类型为%T，值为：%v\n", numSlice, numSlice)
	// 3.在定义数组时，其类型可以是值类型也可以是引用类型，但是不能混用
	// 4.数组在定义后，如果没有赋值，有默认值(零值)
	// 数值类型数组，其零值为0;字符串类型数组，其零值为"" ;bool类型的数组，其零值为false
	var intArr [3]int
	fmt.Printf("intArr 的默认零值为：%v\n", intArr)
	var strArr [3]string
	fmt.Printf("strArr 的默认零值为：%v\n", strArr)
	var boolArr [3]bool
	fmt.Printf("boolArr 的默认零值为：%v\n", boolArr)
	// 5.数组的下标是从0 开始的，当下标超过数组长度时，就会报越界的错误
	// panic:runtime error: index out of range
	// 6.Golang中，数组属于值类型，在默认情况下属于值传递，因此会进行值拷贝，数组见不会相互影响，如果需要进行地址拷贝，可以通过指针传递
	transArr := [...]int{1, 2, 3}
	fmt.Printf("数组transArr 的地址为：%p\n", &transArr)
	testArr(transArr)
	fmt.Printf("调用test函数后的transArr 的值为：%v\n", transArr) // 调用完毕后，transArr 的值并没有变化

	// 将数组指针传递到函数中，再在函数中通过指针修改数组的值
	// 因此如果涉及到调用函数修改同一个数组变量时，建议采用数组指针传递，
	// 因为函数只需要拷贝数组的指针地址，而不需要拷贝数组变量，地址的数据量小，效率高
	testArrPointer(&transArr)
	fmt.Printf("通过地址传递数组后，transArr 的值为：%v\n", transArr)
	// 7.长度时是数组类型的一部分，在传递函数参数时，需要考虑数组的长度
	//var funcArr[1] int
	//testArr(funcArr) // 函数接收的数组长度需要是4个元素的，但是传过去的数组长度是1，因此会报错
}
