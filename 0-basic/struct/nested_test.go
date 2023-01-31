package _struct

import (
	"fmt"
	"testing"
)

type a struct {
	id   int
	name string
}

type b struct {
	a
	age  int
	name string
}

func Test_Nested(t *testing.T) {
	b1 := b{a{1, "a"}, 12, "bb"}
	fmt.Println(b1) // {{1 a} 12 bb}
	// 内外层字段不同时，可直接通过字段名操作内层字段
	fmt.Println(b1.id) // 1
	// 字段名冲突时，外层字段优先级高于内层字段
	b1.name = "bb2"
	fmt.Println(b1) // {{1 a} 12 bb2}
	// 操作冲突的内层字段时，可通过外层变量.内层变量,字段名 来操作
	b1.a.name = "a2"
	fmt.Println(b1) // {{1 a2} 12 bb2}
}
