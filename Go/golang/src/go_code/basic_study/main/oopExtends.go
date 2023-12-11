package main

import "fmt"

// 继承可以解决代码复用，让我们的编程更加靠近人类思维
// 当多个结构体存在相同的属性(字段) 和方法时，可以从这些结构体中抽象出结构体，在该结构体中定义这些相同的属性和方法
// 其他的结构体不需要重新定义这些属性和方法，只需要嵌套一个匿名结构体即可
// 也就是说，在Golang中，如果一个struct 嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承的特性
// 使用继承，代码的复用性、扩展性和维护性得到了提高

type Parent struct {
	Name string
	Age  int
}

type childOne struct {
	Parent // 通过匿名结构体的方法继承Parent 结构体，进而绑定该结构体的方法能够访问父类结构体的属性和方法
}

type childTwo struct {
	Parent
}

func (p *Parent) parentFunc() {
	fmt.Println("调用父类方法")
}

func (c1 *childOne) childOneFunc() {
	fmt.Printf("第一个子类的方法")
}

func (c2 *childTwo) childTwoFunc() {
	fmt.Printf("第二个子类的方法")
}

func main() {
	c1 := childOne{}
	c2 := childTwo{}
	c1.Parent.parentFunc() // 通过子类继承父类，调用父类的方法
	c1.childOneFunc()      // 通过子类，调用自己的方法
	c2.parentFunc()
	c2.childTwoFunc()
	c1.Age = 10
	c1.Name = "Tom~"
	fmt.Printf("c1调用父类的属性，年龄值为%v，名字值为%v\n", c1.Age, c1.Name)
	fmt.Printf("c2调用父类的属性，年龄值为%v，名字值为%v\n", c2.Age, c2.Name)
}
