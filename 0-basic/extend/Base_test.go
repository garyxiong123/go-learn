package extend

import (
	"fmt"
	"math/big"
	"testing"
)

type person struct {
	name   string
	age    int64
	amount big.Int
}

type teacher1 struct {
	person
}

type teacher2 struct {
	*person
}

func (p person) sayHello(name string) {
	println(name)
	println(p.name)
}

type teacher struct {
	person
}

//func (p teacher) sayHello(name string) {
//	println(name)
//	println(p.name)
//}

func Test_1(t *testing.T) {
	p := &person{
		name: "gary",
		age:  33,
	}

	p.sayHello("boy")
	println(p)
}

func Test_2(t *testing.T) {
	p := &teacher{}
	p.sayHello("boy")
	println("p+", p)

	fmt.Println(*p)

	p1 := &teacher1{
		person{
			name: "3",
		},
	}
	println("p1+", p1)
	fmt.Println(*p1)

	p2 := &teacher2{
		&person{
			name: "3",
		},
	}

	println("p2+", p2)
	fmt.Println(*p2)
}
