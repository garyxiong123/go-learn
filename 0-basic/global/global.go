package global

import "go-learn/db/orm"

var name = "gary"

//init()函数是一个特殊的函数，存在以下特性：
//
//不能被其他函数调用，而是在main函数执行之前，自动被调用
//init函数不能作为参数传入
//不能有传入参数和返回值
func init() {
	println("init")
}

//每个源文件都只能包含一个 init() 函数（此处存在错误）。初始化总是以单线程执行，并且按照包的依赖关系顺序执行。
//
//一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性。

var person = &orm.Person{
	Name: "gary",
}

func newPerson(name string) {
	println(person)

}
