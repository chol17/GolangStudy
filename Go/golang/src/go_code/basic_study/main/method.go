package main

import (
	"encoding/json"
	"fmt"
)

// 方法基本介绍
// 在某些情况下，我们必须声明/定义方法, 比如结构体有字段(属性) 的概念，
// 同理，该结构体还需要有一些行为，那么这时就需要用方法实现结构体的这些行为
// Golang中的方法是作用在指定的数据类型上，即方法是和指定的数据类型绑定，
// 因此，自定义类型都可以有方法，而不仅仅是struct 有方法的概念

type PersonStruct struct {
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int
	Pointer *string `json:"pointer"`
}

func (person1 PersonStruct) testPerson1(name string) {
	fmt.Printf("参数值为:%v\n", name)
	person1.Name = name
	sex := person1.Sex
	person1.Pointer = &sex
	fmt.Printf("结构体中name字段的值为：%v, sex字段的值为：%v\n", name, sex)
	fmt.Printf("结构体中pointer字段的值为：%v\n", person1.Pointer)
}

// 方法的定义(声明)
// func(receiver type)methodName(参数列表)(返回列表){
//		方法体
//		return 返回值
//	}
// 说明：
// 1.参数列表：表示方法的输入
// 2.receiver type: 表示这个方法和type这个类型进行绑定，或者说该方法作用于type 类型
// 3.type 可以是结构体，也可以是其他自定义类型
// 4.receiver type类型的实例化变量
// 5.返回列表：表示方法的返回值，可以是多个返回值，多个的情况用括号括起来
// 6.方法主体：表示方法中的处理逻辑
// 7. return: 返回方法的数据，当有返回列表时，对应才有return 返回值

func (person2 *PersonStruct) testPerson2(name string) {
	fmt.Printf("参数值为:%v\n", name)
	person2.Name = name
	person2.Age++
	sex := (*person2).Sex
	(*person2).Pointer = &sex
	// 以指针传递结构体变量的形式调用方法时，person2变量存储的值就是调用该方法的结构体变量的地址
	fmt.Printf("方法中person2的地址为%p\n", person2)
	fmt.Printf("结构体中name字段的值为：%v, sex字段的值为：%v,age 字段的值为：%v\n", name, sex, person2.Age)
	fmt.Printf("结构体中pointer字段的值为：%v\n", (*person2).Pointer)
}

func (person *PersonStruct) String() string {
	return "[" + person.Name + "]"
}

func main() {

	fmt.Printf("main 方法....\n")
	personName := "Duck~~"
	pointer := &personName
	person1 := PersonStruct{
		Sex:     "男",
		Pointer: pointer,
	}
	person2 := &PersonStruct{
		Name:    "Tom~~",
		Age:     20,
		Sex:     "男",
		Pointer: pointer,
	}
	// person2 变量是通过类型推到定义的，因此从代码来看，它是一个指针变量，变量存储的是结构体的地址
	fmt.Printf("结构体变量person2的地址为%p\n", person2)
	fmt.Printf("结构体变量person1的类型为%T\n", person1)
	fmt.Printf("结构体变量person2的类型为%T\n", person2)
	fmt.Printf("调用结构体方法前的结构体变量person1值为%v\n", person1)
	fmt.Printf("调用结构体方法前的结构体变量person2值为%v\n", person2)
	// 总结：
	// 1.testPerson方法和PersonStruct这个结构体绑定，用来描述这个结构体的行为的方法
	// 2.testPerson方法只能通过PersonStruct这个结构体变量调用，
	// 且在方法里修改结构体值类型的属性是不会对结构体变量的值产生影响
	// 3.如果需要通过调用方法修改结构体变量的字段值，那么可用通过指针的方法声明变量并在调用方法的时候传结构体指针
	// 3.方法中指定的结构体变量名称可以自定义命名，只要类型是绑定的结构体类型即可
	person1.testPerson1(personName)

	// 方法的调用和传参机制
	// 方法的调用和传参机制和函数基本一样，不一样的地方是方法调用时，会将调用方法的变量，当做实参也传递给方法
	// 在main函数中定义一个结构体变量，通过结构体变量调用方法时，内存中会把这个结构体变量进行值拷贝，
	// 方法内部的操作完全在一个新的空间中运行，当方法执行完毕后，操作系统对内存进行垃圾回收
	// 因此在使用结构体变量调用方法时，
	// 不仅仅是传递了方法需要的形参，还传递了方法的实例化对象，这个对象以实参的形式传递给方法
	// 注意：如果方法中的结构体参数类型是结构体指针类型，则方法中的结构体变量属于是引用类型，
	// 在方法中操作结构体变量的字段后，对影响到原本的结构体变量的值
	person2.testPerson2(personName)
	jsonStruct1, err1 := json.Marshal(person1)
	jsonStruct2, err2 := json.Marshal(person2)
	if err1 != nil {
		fmt.Printf("错误信息为%v\n", err1)
	} else {
		fmt.Printf("调用结构体的方法后，序列化结构体的值为%v\n", string(jsonStruct1))
	}
	if err2 != nil {
		fmt.Printf("错误信息为%v\n", err2)
	} else {
		fmt.Printf("调用结构体的方法后，序列化结构体的值为%v\n", string(jsonStruct2))
	}
	// 方法的注意事项和细节讨论
	// 1.结构体类型是值类型，在方法调用时，遵守值类型的传递机制，是值拷贝传递方式
	// 2.如果想在代码中通过方法修改结构体变量的值，可以在调用方法时，传递结构体指针的方式传递结构体变量实现
	// 3.Golang中的方法作用在指定的数据类型上的，即方法是和数据类型绑定的，因此自定义类型都可以有方法
	// 4.方法的访问范围控制的规则，和函数一样。方法的首字母小写，只能在当前包中访问;
	// 方法的首字母大写，可以在本包和其他包中访问
	// 5.如果一个变量实现了String()这个方法，那么fmt.Println默认会调用当前代码中实现的这个变量的String()方法进行输出
	// 而不是调用内置的String() 方法，调用String()方法的原因是Println()源码中会在输出的时候调用String()方法，
	// 但是当前代码的结构体方法中实现了String()方法，因此就调用当前结构体的方法
	//fmt.Printf()
	var person3 PersonStruct
	person3.Name = "Mark~~"
	fmt.Printf("打印出来的名字为%v\n", person3.Name)

	// 函数和方法的区别：
	// 1.调用方法不一样
	// 函数的调用方式：函数名(实参列表)
	// 方法的调用方式：变量名.方法名(实参列表)
	// 2.对于普通函数，实参列表接收的为值类型时，就不能将指针类型的数据直接传递;
	// 当接收的为指针类型时，就不能将值类型数据传递进来
	// 3.对于方法(如struct类型的方法)，调用方法的结构体变量不受类型限制，只要该变量的类型和方法绑定的结构体类一致即可，
	// 如果方法中指定的结构体类型为指针类型，那么传进来的结构体变量在调用方法中被修改后，会直接影响结构体变量本身的字段值
	// 如果方法中指定的结构体类型为值类型，那么传进来的结构体变量在调用方法后，不会影响结构体变量本身的字段的值
}
