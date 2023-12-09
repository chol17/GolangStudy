// 打包,包名叫做utils, 虽然这个包名叫utils,但和它在utils 目录下没有半毛钱关系，只是为了方便后续阅读，故包名和其目录名保持一致
// 打的包名叫utils, 后续调用就使用utils.函数名 这样调用，如果包名叫abc,后续调用就使用abc.函数名

package utils

import "fmt"

var Number = 100 // 如果在包文件中定一个能被其他位置调用的变量，这个变量的首字母也必须大写

type FactoryStruct struct { // 当结构体名称的首字母大写时，可以被其他包直接引用
	Name string
	Age  int
}

// 当结构体名称的首字母小写时，其他包无法直接引用该结构体，
// 可以在当前包中定义一个函数，然后函数返回结构体变量的指针，在其他包中调用该方法即可实现工厂模式
type newFactoryStruct struct {
	Name string
	Sex  string
	age  int // 当结构体中的字段首字母小写时，在其他包中也无法正常调用
}

// 在当前包中定义一个方法，绑定结构体，通过方法调用结构体小写的字段，
// 并将小写字段的值或者地址返回，进而在其他包调用该方法访问到结构体中小写字母的字段

func (nfs *newFactoryStruct) NewField() int {
	nfs.age = 20
	return nfs.age
}

func NewFactory() *newFactoryStruct {
	var nfs newFactoryStruct
	return &nfs
}

func init() {
	fmt.Printf("utils package")
}
func Calculate(num int) int {
	// 在包文件中，定义的函数首字母必须大写，相当于java 中的public 申明，表示该函数可导出；如果是首字母小写，表示将函数给私有化申明，就无法从其他代码中调用
	var result int
	result = num * num
	return result
}
