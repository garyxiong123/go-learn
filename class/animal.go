package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Printf("%v is eating", a.Name)
	fmt.Println()
}

type Animal1 struct {
	Name string
}

func (a *Animal1) Eat1() {
	fmt.Printf("%v is eating", a.Name)
	fmt.Println()
}

type Hand struct {
}

func (hand Hand) getHand() {
	fmt.Println(hand)
}

// Cat 继承
type Cat struct {
	*Animal
	*Animal1
	hand Hand
}

func main() {

	cat_ref := &Cat{Animal: &Animal{Name: "Cat"}, hand: Hand{}}
	cat_ref.Eat()

	cat_ref.hand.getHand()

	fmt.Println(cat_ref)

	cat := Cat{Animal: &Animal{Name: "Cat"}}
	cat.Eat()
	fmt.Println(cat)
	//collection.Sum(111)

}
