package main

import "fmt"

// 切片的基本介绍：
// 1.切片的英文是slice
// 2.切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制
// 3.切片的使用和数组类似，遍历切片，访问切片的元素和求切片的长度的方式和数组一样
// 4.切片的长度是可以变化的，因此切片是一个可用动态变化的数组
// 5.切片定义的基本语法：
// var 变量名 [] 类型
// 6. 切片内存空间由三部分组成：引用地址值、切片长度、切片容量

func testSlice(slice []int) {
	slice[0] = 100
}

func main() {
	// 定义一个数组，切片直接通过引用数组定义
	intArr := [...]int{12, 13, 14}
	fmt.Printf("数组的值为：%v\n", intArr)
	// 切片声明的方式1：直接通过引用数组指定元素
	// 数组变量名[start:end] 表示的意思是：
	// 获取数组下标为start的元素到下标为end(但不包含下标为end元素)之间的全部数据
	intSlice := intArr[0:3]
	fmt.Printf("切片的值为：%v\n", intSlice)
	// 因为切片引用到数组该下标的元素，因此修改数组该下标的元素值，对应切片的值也会变化
	intArr[2] = 40
	fmt.Printf("切片的值为：%v\n", intSlice)
	// 同理，修改切片的元素值，其引用的数组对应的值也会发生变化
	intSlice[0] = 11
	fmt.Printf("数组的值为：%v\n", intArr)
	// 获取切片的长度
	fmt.Printf("切片的长度为:%v\n", len(intSlice))
	// 获取切片的容量，所谓容量指的是切片能存储多少个元素，一般容量是切片当前长度的两倍
	fmt.Printf("切片的容量为:%d\n", cap(intSlice))

	// 获取切片的第一个元素以及引用数组的该值；并分别获取对应的地址值
	fmt.Printf("切片的第一个元素值为:%v,引用到的数组的值为：%v\n", intSlice[0], intArr[0])
	fmt.Printf("切片的第一个元素的地址值为：%p, 引用数组的该值对应的地址为:%p\n", &intSlice[0], &intArr[0])
	// 从上述的结果可以看出，切片在内存中的分布是：引用地址+切片长度+切片容量，相当于一个结构体
	// slice 从底层来说，其实就是一个结构体 (struct结构体)
	// type slice struct{
	//			ptr *[3]int
	//          len int
	//			cap int
	//	}

	// 切片声明方式2：使用make 创建切片
	// 基本语法：var 切片名 [] 切片类型 = make([]切片类型, 切片长度, 切片容量)
	sliceVar := make([]int, 3, 8)
	sliceVar[0] = 1
	sliceVar[1] = 2
	sliceVar[2] = 9
	sliceVar = append(sliceVar, 10)
	fmt.Printf("切片的值为：%v\n", sliceVar)
	fmt.Printf("切片的长度为：%v\n", len(sliceVar))
	fmt.Printf("切片的容量为：%v\n", cap(sliceVar))

	// 通过make方式创建切片可以指定切片的大小和容量
	// 如果没有给切片的各个元素赋值，那么就会使用默认值
	// int 类型的切片默认值为0，string 类型的切片默认值为"",bool 类型的切片默认值为false
	// 通过make 方式创建的切片，对应切片引用的数组由make 底层维护，对外不可见，只能通过slice 赋值取值的方式访问

	//切片声明方式3：定义一个切片，直接指定具体数组
	sliceTmp := []int{12, 13, 14}
	fmt.Printf("sliceTmp 切片的值为：%v\n", sliceTmp)

	// 方式1和方式2 定义切片的不同
	// 1.方式1 是直接引用数组，这个数组是事先存在的，可通过修改数组进而改变切片
	// 2.方式2 是通过make来创建切片，make也会创建一个数组，是由切片在底层维护，数组对外不可见，只能通过修改切片访问里面的数组

	// 切片的遍历
	// 1.常规for 遍历
	forSlice := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(forSlice); i++ {
		fmt.Printf("常规遍历切片的元素为forSlice[%v] = %v\n", i, forSlice[i])
	}

	// 2.使用for-range 遍历
	for index, value := range forSlice {
		fmt.Printf("for-range 遍历切片的元素为：forSlice[%v] = %v\n", index, value)
	}

	// 切片注意事项：
	// 1.切片初始化时，var slice = arr[startIndex:endIndex]
	// 说明：切片的取值是从arr 数组下标为startIndex开始(包含下标为startIndex的元素) 到 下标为endIndex的元素(不包含endIndex的元素)
	// 2.切片初始化时，仍然不能越界，范围在0~len(arr) 之间，但是可以动态增长
	// 切片获取数组元素的几种简写方式
	// var slice = arr[0:endIndex] ======> var slice = arr[:endIndex]
	// var slice = arr[startIndex:len(arr)]=====> var slice = arr[startIndex:]
	// var slice = arr[0:len(arr)] =======> var slice = arr[:]
	// 3.cap 是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
	// 4.切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组或者make一个空间供切片使用
	// 5.切片可以再切片
	forSlice1 := forSlice[1:2]
	fmt.Printf("切片再切片的数值为：%v\n", forSlice1)

	// 使用append内置函数，可以对切片进行动态追加
	// 注意：追加的数据类型必须和切片定义的类型一致
	// 切片append 操作的底层原理分析：
	// 1.切片append操作的本质就是对数组扩容
	// 2.Golang 底层会创建一个新的数组newArr(安装扩容后的大小)
	// 3.将slice原来包含的元素拷贝到新的数组newArr
	// 4.slice 重新引用到newArr, 这个newArr 在底层维护，对外不可见
	appendSlice := []int{}
	// 使用append 追加后，appendSlice 的地址没有变化,只是改变了内部维护的数组地址
	// append 操作后，之前内部引用的数组被作为垃圾回收
	fmt.Printf("appendSlice追加前的地址为：%p\n", &appendSlice)
	appendSlice = append(appendSlice, 1) // 追加后有一个返回值接收追加后的结果
	appendSlice[0] = 1
	fmt.Printf("appendSlice追加后的地址为：%p\n", &appendSlice)
	fmt.Printf("appendSlice 追加后的值为：%v\n", appendSlice)
	// 给切片追加切片, 注意append 里面的切片 后面必须带上三个点，固定语法
	appendSlice = append(appendSlice, appendSlice...)
	fmt.Printf("追加切片后的结果为：%v\n", appendSlice)

	// 切片的拷贝操作
	// 切片使用copy 内置函数完成拷贝，注意：只能是切片与切片间才能使用copy
	// 参与拷贝的两个切片的内存空间是独立的，copy 完后，修改其中一个切片的数据，不会影响另外一个切片
	// copy 不会受制于切片长度的影响
	copySlice1 := []int{1, 2, 3, 4, 5}
	copySlice2 := make([]int, 10)
	copy(copySlice2, copySlice1) // 将后面的切片拷贝到前面的切片中
	fmt.Printf("copySlice1切片的值为：%v\n", copySlice1)
	fmt.Printf("copySlice2切片的值为：%v\n", copySlice2)
	copySlice1[1] = 100 // 修改其中参与拷贝的切片，并不会影响完成copy 后的另外一个切片的结果
	fmt.Printf("copySlice2切片的值为：%v\n", copySlice2)
	// 将长度大的切片往长度短的切片中拷贝，长度大的会讲数据覆盖到长度短的切片，不会报错
	copy(copySlice1, copySlice2)
	fmt.Printf("copySlice1的值为：%v\n", copySlice1)

	// 切片是引用类型，所以在传递的过程中，遵守引用传递机制
	funcSlice := []int{1, 2, 3}
	fmt.Printf("调用函数前的切片值为：%v\n", funcSlice)
	testSlice(funcSlice)
	fmt.Printf("调用函数后的切片值为：%v\n", funcSlice)

	// string 和 slice
	// 1.string 底层是一个byte 数组，因此string也可以进行切片处理
	// 2.string类型数据是不可变的，因此不能通过下标的方式修改字符串里面的数据
	// 因此如果要修改字符串内容，可以将字符串转成 byte切片或者rune切片，然后重写string
	stringVar := "hello world"
	stringSlice := stringVar[:]
	fmt.Printf("切片处理后的结果为:%v\n", stringSlice)

	// stringVar[0] = "H" // string类型数据是不可变的，因此不能通过下标的方式修改字符串里面的数据
	byteSlice := []byte(stringSlice)
	fmt.Printf("byteSlice 切片的值为：%c\n", byteSlice)
	byteSlice[0] = 'H'
	stringVar = string(byteSlice)
	fmt.Printf("byteSlice 切片的值为：%c\n", byteSlice)
	fmt.Printf("转换后的字符串为：%v\n", stringVar)
	// 注意：使用[]byte 切片只能处理包含英文和数字的字符串，无法处理包含中文的字符串
	// 因为byte 是按照字节处理，但是一个中文占用3个字节，会发生错误
	// 解决办法： 使用[]rune 切片即可，因为[]rune 切片是按照字符来处理的，兼容汉字和英文
	// []rune 其实也是一种byte 类型的切片，因此哪怕是处理中文切片，赋值的时候需要用单引号
	chineseWordString := "中国你好"
	runeSlice := []rune(chineseWordString)
	runeSlice[2] = '您'
	chineseWordString = string(runeSlice)
	fmt.Printf("转换后的结果为：%v\n", chineseWordString)

	englishWordString := "Hello world"
	englishSlice := []rune(englishWordString)
	englishSlice[6] = 'W'
	englishWordString = string(englishSlice)
	fmt.Printf("转换后的结果为：%v\n", englishWordString)
}
