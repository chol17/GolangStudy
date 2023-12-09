package main

import "fmt"

// 二维数组练习
// 定义一个二维数组，保存3个班，每个班5名同学的成绩
// 求出每个班的平均分和所有班的平均分

func main() {
	scores := [3][5]float64{}
	var sumScore float64
	for index, value := range scores {
		var count float64
		for i, _ := range value {
			fmt.Printf("请输入第%v班第%v个同学的成绩：\n", index+1, i+1)
			var score float64
			fmt.Scanln(&score)
			scores[index][i] = score
			count += float64(score)
			sumScore += float64(score)
		}
		avg_count := count / float64(len(value))
		fmt.Printf("第%v个班的平均成绩为：%v\n", index+1, avg_count)
	}
	sum_avg := sumScore / float64(len(scores)*len(scores[0]))
	fmt.Printf("所有班的平均成绩为：%v\n", sum_avg)
}
