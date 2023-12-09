package main

import "fmt"

func main() {
	//使用switch 把小写类型的char 型转为大写（键盘输入）
	//只转a,b,c,d,e, 其他的输出other
	var char byte
	fmt.Printf("请输入一个字符...")
	fmt.Scanf("%c", &char) // Scanf 函数 获取键盘输入
	switch char {
	case 'a':
		fmt.Printf("A")
	case 'b':
		fmt.Printf("B")
	case 'c':
		fmt.Printf("C")
	case 'd':
		fmt.Printf("D")
	case 'e':
		fmt.Printf("E")
	default:
		fmt.Printf("Other")
	}

	// 对学生成绩大于60分的，输出"合格", 低于60分的，输出"不合格"
	//注：输入的成绩不能大于100
	var score int32
	fmt.Printf("请输入学生成绩：")
	fmt.Scanln(&score) // scanln函数是获取键盘输入的一行内容; scanf 是获取格式化后的键盘输入内容  scanln(&score) 等价于 scanf("%d",&score)
	switch {
	case score > 100:
		fmt.Printf("输入的成绩不能大于100")
	case score >= 60:
		fmt.Printf("合格")
	case score < 60:
		fmt.Printf("不合格")
	}

	//根据用户指定的月份，打印该月份所属的季节
	//3,4,5 为春季；6,7,8为夏季；9,10,11为秋季；12,1,2 为冬季
	var season byte
	fmt.Printf("请输入月份：")
	fmt.Scanln(&season)
	switch {
	case season >= 3 && season <= 5:
		fmt.Printf("春季")
	case season >= 6 && season <= 8:
		fmt.Printf("夏季")
	case season >= 9 && season <= 11:
		fmt.Printf("秋季")
	case season == 12 || season == 1 || season == 2:
		fmt.Printf("冬季")
	default:
		fmt.Printf("输入的月份有误")
	}
}
