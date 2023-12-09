package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 练习1：从终端中循环输入5个成绩，保存到float64数组中，并输出
	//var score int
	//var scoreArray [5]float64
	//for i := 0; ; i++ {
	//	if i < 5 {
	//		fmt.Printf("请输入第%v个元素\n", i+1)
	//		// fmt.Scanln() 从终端输入，传入的是地址，地址是引用类型变量，因此没有返回值，
	//		// 直接操作的是变量地址值
	//		fmt.Scanln(&score)
	//		fmt.Printf("score 的值为：%v\n", score)
	//		scoreArray[i] = float64(score)
	//	} else {
	//		break
	//	}
	//
	//}
	//fmt.Printf("数组存储的数据为：%v\n", scoreArray)

	// 练习2：创建一个byte 类型的26个元素的数组，分别置放'A'-'Z'
	// 使用for 循环访问所有元素并打印出来。提示：字符数据运算 'A' + 1 -> 'B'
	//byteArr := [26]byte{}
	var byteArr [26]byte
	for i := 0; i < 26; i++ {
		if i == 0 {
			byteArr[i] = 'A'
		} else {
			tmp := byteArr[0] + byte(i)
			byteArr[i] = tmp
		}
		//byteArr[i] = 'A' + byte(i)
	}
	fmt.Printf("遍历完之后的数组为：%c\n", byteArr)

	// 练习3：求出数组中的最大值，并获取其下标
	tmpVar := byteArr[0]
	tmpIndex := 0
	for index, value := range byteArr {
		if tmpVar < value {
			tmpVar = value
			tmpIndex = index
		}
	}
	fmt.Printf("遍历后，数组的最大值为：%c, 其下标为:%v\n", tmpVar, tmpIndex)

	// 练习4：求一个数组中元素的和以及平均值
	sumArr := [...]int{3, 4, 5, 6, 7, 8}
	var sumVar int
	for i := 0; i < len(sumArr); i++ {
		sumVar += sumArr[i]
	}
	avgVar := float64(sumVar) / float64(len(sumArr))
	fmt.Printf("数组的总和为：%v\n", sumVar)
	fmt.Printf("数组的平均值为：%v\n", avgVar)

	// 练习5：随机生成5个数，并将其反转打印
	// 思路：生成5个随机数，rand.Intn() 函数，并将数据存到数组中
	// 遍历数组，通过下标反转赋值
	var randArr [5]int
	var reverseArr [5]int
	for index, _ := range randArr {
		// 为了每次生成的随机数都不一样，因此需要给一个seed
		rand.Seed(time.Now().UnixNano())
		randArr[index] = rand.Intn(100)
		reverseArr[5-index-1] = randArr[index]
	}
	fmt.Printf("randArr 的结果为：%v\n", randArr)
	fmt.Printf("反转打印后的结果为：%v\n", reverseArr)
}
