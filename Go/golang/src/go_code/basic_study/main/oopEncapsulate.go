package main

import "fmt"

// 如何理解封装
// 封装（encapsulation）就是把抽象出的字段和对字段的操作封装在一起，数据被保护在内部，程序的其他包只能通过授权的操作（方法），才能对字段进行操作
// 封装的理解和好处
// 1.隐藏实现细节
// 2.可以对数据进行验证，保证安全合理
// 如何体现封装
// 1.对结构体中的属性进行封装
// 2.通过方法，包 实现封装
// 封装的实现步骤
// 1.将结构体、字段(属性) 的首字母小写（不能导出了，其他包不能使用，类似private）
// 2.给结构体所在的包提供一个工厂模式的函数，首字母大写，类似一个构造函数
// 3.提供一个首字母大写的方法（类似其他语言的public）,用于对属性判断并赋值
// 4.提供一个首字母大写的方法（类似其他语言的public）,用于获取属性的值
// 特别说明：在Golang中，并没有特别强调封装，这点不像java，所以不用总是用java 的语法特性来看待Golang，Golang本身对面向对象的特性做了简化的

type personLow struct {
	Name   string
	age    int
	salary float64
}

// 使用工厂模式创建一个外部可以访问的函数，该函数返回私有结构体的地址
func getName(name string) *personLow {
	return &personLow{
		Name: name,
	}
}

func (p *personLow) setAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("输入的年龄有误")
	}
}

func (p personLow) getAge() (age int) {
	return p.age
}

func (p *personLow) setSalary(salary float64) {
	if salary >= 3000 && salary <= 30000 {
		p.salary = salary
	} else {
		fmt.Println("输入的薪资有误")
	}
}

func (p *personLow) getSalary() (sal float64) {
	return p.salary
}

func main() {
	name := "zhangsan"
	personStruct := getName(name)
	personStruct.setAge(100)
	fmt.Printf("%v的年龄为%v\n", name, personStruct.getAge())
	personStruct.setSalary(5000)
	fmt.Printf("%v的薪资为%v\n", name, personStruct.getSalary())

}
