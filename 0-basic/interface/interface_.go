package _interface

import "fmt"

type Animal interface {
	SetName(string)
	GetName() string
}

type Cat struct {
	Name string
}

func (c Cat) SetName(name string) {
	fmt.Println("c addr in: ", c)
	c.Name = name
	fmt.Println(c.GetName())
}

func (c Cat) GetName() string {
	return c.Name
}
