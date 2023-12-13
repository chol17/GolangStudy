package main

import "fmt"

// 接口基本介绍：
// interface 类型可以定义一组方法，但是这些不需要实现，并且interface 不能包含任何变量。到某个自定义类型（比如结构体）要使用的时候，再根据具体情况把这些方法写出来
// 基本语法
// type 接口名 interface{
//		method1(参数列表)返回值列表
//		method2(参数列表)返回值列表
// }
// 小结说明:
// 1.接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚，低耦合的思想
// 2.Golang中的接口，不需要显示的实现。只要一个变量含有接口类型中的所有方法，那么这个变量就实现了这个接口，
// 因此，Golang中没有implement 这样的关键字
// Golang 是基于方法来运作的，接口绑定方法，从而不需要显示的去实现接口

// 接口注意事项
// 1.接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的实例
// 2.接口中所有的方法都没有方法体，即都是没有实现的方法，只有某个结构体绑定的方法包含了接口中申明的方法时，就可以通过接口调用
// 3.在Golang中，一个结构体绑定的方法需要包含某个接口的所有方法，我们就说这个结构体实例实现了该接口
// 4.只要是自定义数据类型，都可以实现接口，不仅仅是结构体类型

type Interface interface {
	stop()
	start()
}

type PersonOne struct {
	name string
}

type PersonTwo struct {
	name string
}

type PersonThree struct {
	name string
}

func (p PersonOne) start() {
	p.name = "zhangsan"
	fmt.Printf("%v开始...\n", p.name)
}

func (p PersonOne) stop() {
	p.name = "zhangsan"
	fmt.Printf("%v停止...\n", p.name)
}

func (p PersonTwo) start() {
	p.name = "lisi"
	fmt.Printf("%v开始...\n", p.name)
}

func (p PersonTwo) stop() {
	p.name = "lisi"
	fmt.Printf("%v停止...\n", p.name)
}

func (p PersonThree) start() {
	p.name = "lisi"
	fmt.Printf("%v开始...\n", p.name)
}

func (p PersonThree) stop() {
	p.name = "lisi"
	fmt.Printf("%v停止...\n", p.name)
}

func (p PersonThree) restart() {
	p.name = "lisi"
	fmt.Printf("%v重新开始...\n", p.name)
}

type Action struct {
}

func (a Action) eat(inter Interface) {
	inter.start()
	inter.stop()
}

// 自定义数据类型，绑定的方法包含了接口的全部方法，就可以使用接口调用
type customType string

func (c customType) start() {
	fmt.Printf("start...\n")
}

func (c customType) stop() {
	fmt.Printf("stop...\n")
}

func main() {
	a := Action{}
	p1 := PersonOne{}
	p2 := PersonTwo{}
	a.eat(p1)
	a.eat(p2)

	// 可以理解为接口中只是负责申明方法，具体方法的具体实现是通过结构体绑定的方法去实现，
	// 只有当结构体绑定的方法包含了接口申明的方法时，才可以使用接口类型变量去调用结构体中的方法
	var interfac1 Interface
	interfac1 = PersonThree{} // 因为该结构体绑定的方法包含了接口申明的全部方法，因此可以用该结构体实例赋值给结构体变量
	interfac1.start()         // 进而可以使用接口变量调用结构体绑定的方法
	//interfac1.restart() // 而且只能调用接口中申明的方法

	var interfac2 Interface
	var custom customType
	interfac2 = custom
	interfac2.stop()
	interfac2.start()

}
