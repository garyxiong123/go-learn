package _interface

import (
	"fmt"
	"testing"
)

//不带任何方法的 interface, 这种类型的 interface 叫 empty interface
type EmptyVariable interface{}

//空 interface =Object
//interface{} 是一个空的 interface 类型， 空的 interface 没有方法，
//所以可以认为所有的类型都实现了 interface{}。如果定义一个函数参数是 interface{} 类型， 这个函数应该可以接受任何类型作为它的参数。
func Test_EmptyVariable(t *testing.T) {

	variable := EmptyVariable("sss")

	println(variable)
	println(variable)

}

func Test_EmptyVariable_MulityType(t *testing.T) {

	vString := EmptyVariable("sss")
	println(vString)
	vInt := EmptyVariable(123)
	println(vInt)

}

//interface 是一种类型

//interface 变量存储的是实现者的值

func Test_Vaule(t *testing.T) {
	// c := &Cat{Name: "Cat"}
	// fmt.Println("c addr out: ", c)
	// c.SetName("DogCat")
	// fmt.Println(c.GetName())
	c := Cat{}
	var i Animal
	i = &c //把变量赋值给一个 interface
	fmt.Println(i.GetName())

}

//如何判断 interface 变量存储的是哪种类型的值
//当一个 interface 被多个类型实现时， 有时候我们需要区分 interface 的变量究竟存储的是哪种类型的值，
//Go 可以使用 comma, ok 的形式做区分 value, ok := em.(T) : em 是 interface 类型的变量，
//T 代表要断言的类型， value 是 interface 变量存储的值， ok 是 bool 类型标识是否为该断言的类型 T
func Test_interface_type(t *testing.T) {
	c := Cat{Name: "cat"}
	var i Animal
	i = &c //把变量赋值给一个 interface
	if t, ok := i.(*Cat); ok {
		fmt.Println("c implement i:", t)
	}
	//当然我们也可以用 switch 语句
	//value, ok := em.(T)
	//
	//switch t := i.(*Cat) {
	//case *S:
	//	fmt.Println("i store *S", t)
	//case *R:
	//	fmt.Println("i store *R", t)
	//}
}

//选择 interface 的实现者

//Go interface 的底层实现
//interface 底层结构
//Type Assertion(类型断言)
