package main

import "fmt"

// 继承可以解决代码复用，让我们的编程更加靠近人类思维
// 当多个结构体存在相同的属性(字段) 和方法时，可以从这些结构体中抽象出结构体，在该结构体中定义这些相同的属性和方法
// 其他的结构体不需要重新定义这些属性和方法，只需要嵌套一个匿名结构体即可
// 也就是说，在Golang中，如果一个struct 嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承的特性
// 使用继承，代码的复用性、扩展性和维护性得到了提高

// 继承的深入讨论
// 1.结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大写或者小写的字段、方法都可以用
// 2.匿名结构体可以简化为子类结构体变量.父类属性、子类结构体变量.父类方法
// 3.当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
// 4.结构体嵌入两个(或多个)匿名结构体，如两个匿名结构体有相同的字段和方法（同时结构体本身没有同名的字段和方法），在访问时，就必须明确指定匿名结构体名字，否则编译报错
// 5.如果一个结构体潜逃了一个有名的结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或者方法时，必须带上结构体的名字
// 6.嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接给各个匿名结构体字段赋值
// 7.如果一个结构体嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承
// 8.如果嵌套的匿名结构体有相同的字段名或者方法名，则在访问时，需要通过匿名结构体名称来区分访问
// 9.为了保证代码的简洁性，建议大家尽量不使用多重继承

type Parent struct {
	Name      string
	Age       int
	man       // 继承小写字母开头的结构体，使其成为匿名结构体
	sameField // 父类结构体中有和当前结构体中相同的属性
}

type man struct {
	weight   float64
	handsome bool
}

type sameField struct {
	Name string
}

type childOne struct {
	Parent // 通过匿名结构体的方法继承Parent 结构体，进而绑定该结构体的方法能够访问父类结构体的属性和方法
}

type childTwo struct {
	Parent
}

type multiStruct struct {
	childOne
	childTwo
}

type hasNameStruct struct {
	m man // 结构体中嵌套了有名字的结构体，访问时需要带上名字访问继承的结构体中的属性或者方法
}

type BasicStruct struct {
	sameField
	int     // 结构体中继承基本数据类型
	i   int // 结构体中继承多个相同类型的基本数据类型，可以通过指定名称加以区分
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

func (p Parent) eatLunch() {
	p.handsome = true // 给继承的匿名结构体中的小写字母开头的变量赋值
	p.weight = 120.0
	p.Name = "zhangsan"
	fmt.Printf("%v的体重为%v,是否帅气:%v\n", p.Name, p.weight, p.handsome)
}

func (p Parent) sameFieldFunc() {
	fmt.Printf("名字为%v\n", p.Name)
}

func (s *sameField) SetSameFieldFunc(name string) {
	s.Name = name
}

func (s *sameField) GetSameFieldFunc() {
	fmt.Printf("Name 的值为%v\n", s.Name)
}

func main() {
	c1 := childOne{}
	c2 := childTwo{}
	c1.Parent.parentFunc() // 通过子类继承父类，调用父类的方法
	c1.parentFunc()        // 使用简化后的访问方式调用父类的方法，可以这么访问的原因是：c1结构体变量在访问的时候，会先检索它自身的方法中是否有该方法，如果有，直接访问；反之会检索它继承的结构体是否有该方法
	c1.childOneFunc()      // 通过子类，调用自己的方法
	c2.parentFunc()
	c2.childTwoFunc()
	c1.Age = 10
	c1.Name = "Tom~"
	fmt.Printf("c1调用父类的属性，年龄值为%v，名字值为%v\n", c1.Age, c1.Name)
	fmt.Printf("c2调用父类的属性，年龄值为%v，名字值为%v\n", c2.Age, c2.Name)
	p := Parent{}
	p.eatLunch()
	fmt.Printf("p.Name 的值为%v\n", p.Name)

	// 如果在父类结构体中存在和子类结构体相同的字段时，此时和子类绑定的方法中用到该相同的字段时，优先使用的是子类自己的该字段
	s := sameField{}
	s.SetSameFieldFunc("lisi")
	s.GetSameFieldFunc()
	p.sameFieldFunc() // 使用子类自己的字段

	m := multiStruct{}
	//fmt.Printf("名字为:%v\n", m.Name)  // 同一个结构体中，继承多个结构体，且继承的结构体中的属性相同，就必须在调用属性的时候执行父类名字，否则就会报ambiguous selector 的错
	fmt.Printf("调用的名字为%v\n", m.childOne.Name)

	h1 := hasNameStruct{}
	fmt.Printf("有名字的结构访问，weight值为%v\n", h1.m.weight) // 访问带有名字的结构体，访问时需要带上结构体名字进行访问组合的结构体的方法和属性

	h2 := hasNameStruct{m: man{
		weight:   120.02,
		handsome: true,
	}, // 注意最后要加一个逗号
	}
	fmt.Printf("hasNameStruct 结构体的初始化值为%v,%v\n", h2.m.weight, h2.m.handsome)

	basic := BasicStruct{}
	basic.int = 10 // 访问结构体中基本数据类型的结构体，可以直接访问
	basic.i = 100
	fmt.Printf("基本数据类型的结构体的值为:%v,%v\n", basic.int, basic.i)
}
