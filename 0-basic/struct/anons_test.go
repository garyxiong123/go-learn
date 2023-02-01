package _struct

import (
	"fmt"
	"testing"
)

type struct2 struct {
	id int
	string
}

//匿名字段，即在定义结构体时，只定义字段的类型，不定义字段的名称；
//
//　　匿名字段的字段名默认为 字段类型的名称，所以一个结构体中同一类型的匿名字段最多只能存在一个，不同类型的匿名字段可存在多个；

func Test_NO_Name(t *testing.T) {
	st1 := struct2{1, "12"}
	// 输出：1 12
	fmt.Println(st1.id, st1.string)
	st1.string = "123"
	// 输出：1 123
	fmt.Println(st1.id, st1.string)

}
