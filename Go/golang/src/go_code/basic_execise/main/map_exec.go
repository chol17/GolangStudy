package main

import "fmt"

// 使用一个key-value 的map 存放学生的信息
// 存放3个学生的信息，每个学生有name 和 sex 信息

func main() {
	studentMap := make(map[string]map[string]string)
	studentMap["001"] = map[string]string{
		"name": "tom",
		"sex":  "男",
	}
	studentMap["002"] = map[string]string{
		"name": "jack",
		"sex":  "男",
	}
	studentMap["003"] = map[string]string{
		"name": "Anny",
		"sex":  "女",
	}
	fmt.Printf("studentMap的值为：%v\n", studentMap)
	// 根据key 取值
	fmt.Printf("studentMap 中003的值为：%v\n", studentMap["003"])
	fmt.Printf("studentMap 中003的性别为：%v\n", studentMap["003"]["sex"])

}
