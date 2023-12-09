package main

import (
	"encoding/json"
	"fmt"
)

// 结构体
// 1.将一类事物的特性提取出来，形成一个新的数据类型，就是结构体
// 2.通过这个结构体，我们可以创建多个实例/对象
// 3.结构体是自定义的数据类型，代表一类事物

// 结构体定义
// type 结构体名称 struct{
//
// }
// 其中 type 和 struct 都是关键字

type Person struct {
	// 定义结构体，其属性没有赋值时，统一是用的该类型的默认值
	// 结构体名称首字母大写时，表示可以在其他包里引用该结构体
	// 结构体中的属性名称首字母大写时，表示可以在其他包中引用该属性
	// 结构体中的属性类型可以是基本数据类型、数组、引用类型
	Name  string
	Sex   string
	Age   int
	Score [3]int
	// 当结构体中的字段是map、指针、切片这类引用类型时，默认值是nil, 没有分配内存空间
	pMap    map[string]string
	pointer *float64
	slice   []string
}

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, leftDown Point // 在一个结构体的字段定义的时候，指定的类型为其他结构体类型
}

type Rect2 struct {
	left, right *Point // 该属性存储的是Point结构体的地址
	intVal      *int
}

type TransStruct1 struct {
	name string
	age  int
}

type TransStruct2 struct {
	name string
	age  int
}

type TransStruct3 TransStruct2

type Reflect struct {
	Name string `json:"name"` //  `json:"name"` 就是struct 的tag, 也就是结构体在被序列化成json字符串后的key的名字
	Age  int    `json:"age"`
}

func main() {
	var person Person // 实例化一个结构体
	fmt.Printf("结构体变量的地址为%p\n", &person)
	fmt.Printf("实例化person的值为%v\n", person)
	person.Name = "张三~"
	person.Sex = "男"
	person.Age = 27
	fmt.Printf("给对象赋值后，person的值为%v\n", person)

	// 结构体注意事项：
	// 1.字段/属性声明方式同变量 字段名  字段类型
	// 2.字段的类型可以为：基本数据类型、数组、引用类型
	// 3.在创建一个结构体时变量后，如果没有给字段赋值，字段都对应一个默认值(零值)，
	// 布尔类型默认值是false;数值类型默认值是0;字符串类型默认值是""
	// 数组类型的默认值和数组的类型有关，[3]int --->默认值就是 [0,0,0]
	// 指针、slice和map类型的零值是nil, 即还没有分配内存空间
	// 4.不同结构体变量字段是独立的，互补影响，修改一个结构体变量字段的值，不会影响另一个结构体变量该字段的值
	// 因为结构体是值类型，不同结构体变量之间使用的是值拷贝
	// 5. 结构体中引用类型的字段默认值是 nil, 如果要使用，需要用到make 分配内存空间后再赋值
	if person.pMap == nil {
		fmt.Printf("map类型的属性默认值为%v\n", person.pMap)
	}
	if person.pointer == nil {
		fmt.Printf("指针类型的属性默认值为:%v\n", person.pointer)
	}
	if person.slice == nil {
		fmt.Printf("切片类型的属性默认值为%v\n", person.slice)
	}
	// 直接给结构体中引用类型的字段赋值会报错
	// person.slice[0] = "张三~~"
	// 需要使用make 分配内存空间后再赋值才行
	person.slice = make([]string, 1)
	person.slice[0] = "李四~~"
	person.slice = append(person.slice, "张三~~")
	fmt.Printf("切片的值为：%v\n", person.slice)

	person.pMap = map[string]string{
		"name": "张三@@",
		"sex":  "男",
	}
	fmt.Printf("map的值为：%v\n", person.pMap)

	intVar := 33.3
	person.pointer = &intVar
	fmt.Printf("指针类型的值为：%v, 寻址后的结果为%v", person.pointer, *person.pointer)

	// 结构体定义的几种方式
	// 1.如同定义变量一样定义结构体变量
	var per1 Person
	fmt.Printf("结构体变量的值为%v\n", per1)
	// 2.在创建结构体变量时就赋予字段值
	per2 := Person{Name: "Duck", Sex: "男"} // 定义结构体变量同时赋值时，不一定需要把结构体中的全部变量都赋值
	fmt.Printf("结构体变量的值为:%v\n", per2)

	// 说明：方式3和方式4 定义的结构体变量返回的是结构体指针
	// 结构体指针访问字段的标准方式是：{*结构体变量}.字段名，
	// 但是Golang中做了简化，也支持结构体指针变量.字段名 的方式对字段赋值，
	// 其编译器底层会对指针变量.字段名的书写方式转换为(*结构体变量).字段名
	// 3.通过指针返回结构体变量
	// 在Golang中，通过new关键字，创建的对象是一个指向结构体的指针类型变量
	var per3 = new(Person)
	// 因为per3 是一个指针变量，因此需要通过指针取值的方式重新对字段赋值
	(*per3).Age = 20
	// Golang 中会对per3.Name 加上取值变成 (*per3).Name
	per3.Name = "张三**" // 在Golang中，设计者为了方便程序员编写，底层会对per3.Name 进行处理，底层会用取值的方式操作变量
	fmt.Printf("结构体变量per3的Name字段的值为：%v\n", (*per3).Name)
	fmt.Printf("结构体变量per3的Age字段的值为：%v\n", (*per3).Age)
	// 4.通过指针
	// per4 结构体变量时一个指针变量，因此在赋值的时候需要用指针取值的方式获取，
	// 但是Golang设计者为了简化代码，底层会对结构体变量做取值转换
	var per4 *Person = &Person{
		Name: "wang wu~~",
	}
	fmt.Printf("结构体变量per4的name 字段值为：%v\n", per4.Name)
	per4.Name = "王五&&"
	fmt.Printf("结构体变量per4的name 字段值为：%v\n", per4.Name)

	// Golang中结构体中的多级指针
	per5 := Person{
		Name: "Jack~~",
	}
	per6 := &per5
	// per5结构体变量是一个指针变量，指针地址指向的是结构体，但是per5 也有它自己的地址
	// per6变量的存储的值就是结构体变量per5自己的地址值，通过指针引用，可以访问到per5结构体字段的值
	fmt.Printf("结构体变量per5的地址为%p\n", &per5)
	fmt.Printf("结构体变量per6的值为%p\n", per6)
	// 获取到结构体所有字段的地址发现，结构体变量的地址其实是指向结构体中第一个字段的地址
	fmt.Printf("name的地址为：%p,sex的地址为%p，age的地址为%p\n", &per5.Name, &per5.Sex, &per5.Age)

	// 注意，使用结构体变量取值的方式有结构体变量.字段名 或者 (*结构体变量).字段名
	// 而不能使用 *结构体变量.字段名，因为在运算符的优先级中，. 的优先级高于 *
	//name := *per6.Name // 错误的写法
	name := per6.Name
	fmt.Printf("name的值为：%v\n", name)
	name1 := (*per6).Name
	fmt.Printf("name1的值为：%v\n", name1)

	// 结构体使用注意细节
	// 1.当结构体的字段类型中，存在类型是另外一个结构体类型，那么该字段就不单单是一个字段，而是一个结构体变量，
	// 通过该字段可以访问到它指向的结构体中的全部属性
	// 2.当结构体中的类型是指针类型时，其本身的地址不一定是连续的，但他们指向的地址是连续的
	// 结构体中的类型是非指针类型时，其地址一定是连续的;
	r1 := Rect{leftUp: Point{
		x: 10,
		y: 20,
	}}
	x := r1.leftUp.x
	fmt.Printf("x 的值为%v\n", x)
	var r2 Rect
	r2.leftDown.x = 33
	r2.leftUp.y = 100
	fmt.Printf("r2 结构体的值为%v\n", r2)

	var r3 Rect2
	intVal := 11
	r3.intVal = &(intVal)
	fmt.Printf("int指针字段的值为：%v\n", *r3.intVal)
	// 当结构体中的字段类型是指针类型时，在给该字段赋值前，需要给一个默认地址，否则在调用指针变量时会出现空指针的情况
	r3.left = &Point{}
	r3.right = &Point{}
	r3.left.x = 55
	r3.left.y = 56
	r3.right.x = 66
	r3.right.y = 67
	fmt.Printf("结构体中指针字段指向的地址为：left.x:%p,\nleft.y:%p,\nright.x:%p,\nright.y:%p\n\n",
		&r3.left.x, &r3.left.y, &r3.right.x, &r3.right.y)
	fmt.Printf("结构体中指针字段的值为：left:%p,right:%p\n", r3.left, r3.right)
	fmt.Printf("结构体中指针字段本身的地址为：left:%p,right:%p", &r3.left, &r3.right)

	// 3.结构体是用户单独定义的类型，和其他类型相互转换时，需要有安全相同的字段(包含名称、类型、个数,结构体中的字段顺序都完全相同)
	var struct1 TransStruct1
	var struct2 TransStruct2
	fmt.Printf("结构体struct1的值为：%v,\n结构体struct2的值为%v\n", struct1, struct2)
	// 将结构体进行转换
	struct1 = TransStruct1(struct2)
	fmt.Printf("转换后的结构体struct1的值为：%v\n", struct1)

	// 4.结构体进行type重新定义(相当于取别名), Golang 中会认为是一种新的数据类型,但是相互间可以强转
	// 【在函数的里面有介绍过type 起别名的概念】
	//var struct3 TransStruct3
	struct3 := TransStruct3(struct1)
	fmt.Printf("转换后的struct3 的值为%v\n", struct3)

	// 5.结构体中每个字段，可以写上一个tag,该tag 可以通过反射机制获取，常见的使用场景就是序列化和反序列化
	// 指定tag 的好处：
	// ① 方便前后端字段的大小写书写统一;
	// ② 在不用修改结构体字段的首字母大小写的前提下，让序列化结构体对象的名称可以满足实际需要
	// 因为一旦修改结构体字段的大小写后，该字段的作用域必定会发生变化，
	// 如果换成小写后，使用json 包就无法访问到当前包中的结构体字段了，进而序列化后的json 字符串中就没有该字段了
	reflect := Reflect{
		Name: "红孩儿",
		Age:  120,
	}
	// Golang 中使用json.Marshal 序列化struct 结构体对象，
	// 序列化后的json字符串中的key的名称就是按照结构体声明的时候，tag指定的字段名称
	// json.Marshal底层会根据反射机制，将结构体对象的字段名称转化成tag指定的名称
	// json.Marshal 序列化后返回序列化后的json字符串和报错信息，
	// 可以通过判断报错信息是否为空，进而得知序列化是否成功
	jsonStr, err := json.Marshal(reflect)
	if err != nil {
		fmt.Printf("错误信息为:%v\n", err)
	} else {
		fmt.Printf("序列化后的数据为：%v\n", string(jsonStr))
	}
	fmt.Printf("序列化后的变量类型为：%T\n", string(jsonStr))
}
