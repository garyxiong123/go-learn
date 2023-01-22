package main

import (
	"flag"
	"fmt"
)

//--name="sss" --age=33  或者 -name="sss"
//flag 包支持的命令行参数类型有 bool、int、int64、uint、uint64、float、float64、string、duration，如下表所示：
func main() {

	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	flag.Parse()

	fmt.Println("name:", *name)

	flag.Args()  //返回命令行参数后的其他参数，以 []string 类型
	flag.NArg()  //返回命令行参数后的其他参数个数
	flag.NFlag() //返回使用的命令行参 数个数

	println(name)

	println(age)
}
