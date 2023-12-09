package main

import (
	"fmt"
	"golang/src/go_code/basic_study/utils"
)

// Golang中的结构体没有构造函数，通常可以使用工厂模式来解决这个问题
// 我们知道， 如果从其他包中导入的变量、结构体时，只有当其他包中的变量、结构体的名称首字母大写时才能正常导入,
// 如果其他包中定义的变量、结构体名称的首字母不是大写，就无法正常导入，因此需要使用工厂模式
// 所谓工厂模式：
// 1.在其他包中引用当前包中首字母小写的结构体时，需要使用工厂模式
// 2.对于结构体名称是小写的情况，在当前包中定义一个函数，
// 通过函数返回值将结构体变量返回，其他包就可以通过调用该函数间接访问到该结构体
// 3.对于结构体的字段名称是小写的情况，可以在当前包中定义一个方法，该方法绑定的结构体类型是指针类型，
// 通过方法返回结构体中首字母是小写的字段，在其他包中通过结构体变量调用该方法，从而获取到结构体中首字母小写的字段值

func main() {
	var fs utils.FactoryStruct
	fs.Name = "张三~~"
	fmt.Printf("导入的结构体变量值为%v\n", fs)

	// 调用工厂模式创建的函数，通过函数的返回值调用其他包中小写字母开头的结构体字段
	nsf := utils.NewFactory()
	// 因为返回的是结构体指针变量，对于结构体变量的操作和结构体指针变量的操作一致，直接用结构体变量.字段名的方式取值
	nsf.Name = "Mack~"
	nsf.Sex = "男"
	// 因为工厂模式的函数返回的是结构体指针变量，因此需要用指针取值的方式得到变量的值
	fmt.Printf("调用工厂函数返回的值为%v\n", *nsf)

	//var age = 20
	agePointer := nsf.NewField()
	//*agePointer = age
	fmt.Printf("结构体中的值为%v\n", agePointer)
	fmt.Printf("结构体变量的值为%v\n", *nsf)
}
