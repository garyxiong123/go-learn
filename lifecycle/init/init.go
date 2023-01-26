package main

import (
	"fmt"
	_ "go-learn/lifecycle/init1"
	_ "go-learn/lifecycle/init2"
	"os"
)

//经测试，每个源文件可以包含多个init函数，而且会按先后顺序执行，优先级高于main
func main() {
	println(os.Args[0])

	fmt.Println("main")
}

//func init() {
//	fmt.Println("init1")
//}
//
//func init() {
//	fmt.Println("init2")
//}
//
//func init() {
//	fmt.Println("init3")
//}

//2、导包顺序决定init函数执行顺序：
//
//go中不同包中init函数的执行顺序是根据包的导入关系决定的。
