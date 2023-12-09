package main

import "fmt"

type Person struct {
	Name string `json:"name"`
}

func (person Person) testPerson() {
	person.Name = "张三~~"
	fmt.Printf("%v 是一个好人~\n", person.Name)
}

func (Person Person) calculate(number int) (sumData int) {
	for i := 0; i <= number; i++ {
		fmt.Printf("i的值为%v\n", i)
		sumData += i
	}
	fmt.Printf("累加的结果为%v\n", sumData)
	return sumData
}

func main() {
	var person Person   // 定义结构体变量
	person.testPerson() // 通过结构体变量调用该结构体指定的方法
	sumData := person.calculate(100)
	fmt.Printf("求和的结果为%v\n", sumData)
}
