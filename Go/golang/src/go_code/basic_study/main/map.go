package main

import (
	"fmt"
	"sort"
)

// map 的基本概念
// 语法：var 变量名 map[key数据类型]值数据类型
// Golang中，key的数据类型可以是bool、整型、string、指针、channel， 也可以是接口、结构体、数组
// 通常，key是 int 和 string 类型
// 不可以将slice、map 和 function 作为key, 因为这三种类型没法使用 == 做判断
// value的类型通常是整型、浮点型、string、map、struct
// Golang中，map 是无序的数据类型

func modifyMap(mapArgs map[string]string) {
	mapArgs["name"] = "张三~"
	mapArgs["age"] = "27"
}

type Student struct {
	Name string
	Age  int
}

func main() {

	// map 的定义
	// map 类型的数据在声明的时候不会分配内存，初始化需要make分配内存后才能赋值和使用
	// key 是string类型，value 也是string类型
	var stringMap map[string]string
	fmt.Printf("stringMap 的值为：%v\n", stringMap)
	// key 是int 类型，value 也是int 类型
	var intMap map[int]int
	fmt.Printf("intMap的值为：%v\n", intMap)
	// key 是string 类型，value 是map 类型
	var strMap map[string]map[string]int
	fmt.Printf("strMap的值为%v\n", strMap)
	fmt.Printf("strMap的类型为%T\n", strMap)
	// 此时没有给声明的map使用make 分配内存空间，因此无法使用，程序报错panic: assignment to entry in nil map
	//stringMap["No.1"] = "hello"
	stringMap = make(map[string]string, 10) // 后面的10表示给这个map 指定key 的最大个数
	stringMap["No.1"] = "hello"
	fmt.Printf("stringMap的值为：%v\n", stringMap)
	// 如果给同一个key赋值不同的value, 会将原来的value 覆盖掉
	// 也就是说key 在map 中是唯一的，但是不同key 对应的value 可以是一样的
	stringMap["No.1"] = "Hello"
	fmt.Printf("stringMap的值为：%v\n", stringMap)

	// map 的定义方式
	// 方式1：先声明map变量，然后再使用make分配内存空间
	var mapVar map[string]string
	mapVar = make(map[string]string, 10)
	mapVar["city"] = "武汉"
	fmt.Printf("mapVar的值为：%v\n", mapVar)

	// 方式2：直接使用make 初始化一个map
	mapVar2 := make(map[string]int)
	mapVar2["age"] = 27
	fmt.Printf("mapVar2的值为：%v\n", mapVar2)

	// 方式3：直接在声明的时候给map 赋值
	// 注意：花括号中value 后面必须要有逗号
	mapVar3 := map[string]string{
		"name": "tom",
	}
	fmt.Printf("mapVar3的值为：%v\n", mapVar3)

	// 方式4：定义一个value 也是map 的数据
	valueMapVar := make(map[string]map[string]int)
	valueMapVar["002"] = map[string]int{
		"age": 27,
	}
	fmt.Printf("valueMapVar的值为：%v\n", valueMapVar)

	// map 数据类型的增删改查操作
	curdMap := make(map[string]string)
	curdMap["name"] = "Jack" // map 的新增：往map类型的数据中赋值不存在的key
	fmt.Printf("新增后curdMap 的值为：%v\n", curdMap)
	curdMap["name"] = "Duck" // map 的修改：对已存在的key 进行重新赋值
	fmt.Printf("修改后curdMap 的值为：%v\n", curdMap)
	// map 的查询：通过指定key 获取对应的value 值; 如果值依然是一个map 类型，还可以继续从中通过key 获取value
	mapValue := curdMap["name"]
	fmt.Printf("获取到curdMap的值为%v\n", mapValue)
	// map 删除操作使用内置的delete函数，指定要删除的map变量和变量的key
	// 如果指定的key 存在，则从map 中删除掉; 如果不存在，程序也不会报错
	delete(curdMap, "name")
	fmt.Printf("删除key 后的curdMap值为：%v\n", curdMap)
	delete(curdMap, "age") // 删除一个不存在的key,程序也不会报错
	fmt.Printf("curdMap的地址为%p\n", &curdMap)
	// map 类型数据删除说明：
	// Golang 中没有一个专门的方法删除所有的key,只能通过遍历map, 逐个删除key
	// 或者使用make() 重新申明一个新的map, 进而之前的map 没有被引用，会被系统进行垃圾回收
	curdMap = make(map[string]string)
	fmt.Printf("重新声明的curdMap的地址为：%p\n", curdMap)
	curdMap["name"] = "Lucy"
	getValue := curdMap["name"]
	fmt.Printf("获取到的值为：%v\n", getValue)
	// map 的查找，通过返回的查找标识 ，判断该key 是否存在
	getVal, flag := curdMap["name"] // map 类型的数据获取时，会返回key 对应的值和是否获取到的标识
	fmt.Printf("从curlMap中获取到的值为：%v, 是否获取到的标识为：%v\n", getVal, flag)
	if flag {
		fmt.Printf("获取到的值为：%v\n", getVal)
	} else {
		fmt.Printf("没获取到")
	}

	// map 类型数据的遍历
	// 注意：map 的遍历只能用for-range，不能用常规的for 遍历，
	// 因为常规的for 遍历需要知道索引下标值，并对下标值做累加，而map数据类型没有下标，因此无法使用
	forRangeMap := make(map[string]string)
	forRangeMap["001"] = "Tom~"
	forRangeMap["002"] = "Timmy~"
	for key, value := range forRangeMap {
		fmt.Printf("遍历的key为：%v, value 为：%v\n", key, value)
	}

	// 使用for-range 遍历一个结构较复杂的map
	complexMAP := make(map[string]map[string]string)
	complexMAP["001"] = map[string]string{
		"name": "Tom",
		"sex":  "男",
	}
	complexMAP["002"] = map[string]string{
		"name": "Lucy",
		"sex":  "女",
	}
	for _, mapVal := range complexMAP {
		for key, value := range mapVal {
			fmt.Printf("key 为%v, value 为%v\n", key, value)
		}
	}

	// 获取map 的长度，使用内置的len函数
	mapLength := len(complexMAP)
	fmt.Printf("map 的长度为：%v\n", mapLength)

	// map 切片，即切片的数据类型是map,则我们称之为map 切片，这样使用，就可以让切片中map个数动态变化
	// map 切片相当于是切片中的元素是map 类型
	slice := make([]map[string]string, 1)
	fmt.Printf("切片的长度为:%v\n", len(slice))
	slice[0] = map[string]string{
		"name": "李四",
		"sex":  "男",
	}
	slice = append(slice, map[string]string{
		"name": "张三",
		"sex":  "男",
	})
	for index, mapVar := range slice {
		fmt.Printf("切片的下标是:%v\n", index)
		for key, value := range mapVar {
			fmt.Printf("切片中的map的key是%v, value是%v\n", key, value)
		}
	}

	// 数组切片，即数组中元素是map 类型，
	// 因为数组长度在声明的时候就已经确定了，因此就不能像切片那样动态变化
	arrayMap := [3]map[string]string{make(map[string]string), make(map[string]string), {"name": "jack"}}
	arrayMap[0] = map[string]string{
		"name": "Tom~",
	}
	fmt.Printf("arrayMap的值为%v\n", arrayMap)

	// map 类型的数据排序
	// Golang中没有一个专门的方法针对map 的key 进行排序
	// Golang 中的map 是无序的, 注意不是按照顺序存放的，因此每次变量的顺序都可能不一样
	// 因此如果要对map数据进行排序，则需要通过key的顺序进行排序
	// 1.现将map 中的key 存放到切片中
	// 2.对切片排序
	// 3.遍历切片，按照key输出结果
	sortSlice := []int{}
	sortMap := make(map[int]string)
	sortMap[0] = "0"
	sortMap[1] = "1"
	sortMap[2] = "2"
	sortMap[3] = "3"
	sortMap[4] = "4"
	sortMap[5] = "5"
	for key, _ := range sortMap {
		sortSlice = append(sortSlice, key)
	}
	fmt.Printf("排序前sortSlice的值为%v\n", sortSlice)
	sort.Ints(sortSlice)
	fmt.Printf("排序后的sortSlice的值为%v\n", sortSlice)
	for _, val := range sortSlice {
		fmt.Printf("顺序输出map的值为：%v\n", sortMap[val])
	}

	// map 使用细节
	// 1.map是引用类型，遵守引用类型的传递机制，在一个函数接收map，函数内修改map后，会直接修改map的值
	// 2.map的容量达到后，再往map中增加元素，map会自动扩容并不会发生panic报错，也就是说map能动态增长键值对
	// 3.map的value也经常使用struct 类型,更适合管理复杂的数据，比value是map类型更好
	funcMap := map[string]string{
		"name": "张三",
	}
	fmt.Printf("调用函数之前的funcMap值为%v\n", funcMap)
	modifyMap(funcMap)
	fmt.Printf("调用函数之后的funcMap的值为%v\n", funcMap) //调用函数后，map 值发生变化，说明map 是引用类型

	// 结构体map
	structMap := map[string]Student{
		"001": {Name: "张三", Age: 27},
		"002": {Name: "李四", Age: 28},
	}
	fmt.Printf("结构体map的值为%v\n", structMap)
	// 遍历结构体map
	for key, structVal := range structMap {
		fmt.Printf("结构体map的key为：%v, 属性中name为%v, age为%v\n", key, structVal.Name, structVal.Age)
	}
}
