package tclass

import (
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {

	cat_ref := &Cat{Animal: &Animal{Name: "Cat"}, hand: Hand{}}
	cat_ref.Eat()

	cat_ref.hand.getHand()

	fmt.Println(cat_ref)

	cat := Cat{Animal: &Animal{Name: "Cat"}}
	cat.Eat()
	fmt.Println(cat)
	//collection.Sum(111)

}
