package main

import "fmt"

//test
type Phone interface {
	speak()
	read()
}

type IPhone struct {
	name string
}

func (i IPhone) speak() {
	fmt.Print("speak")
}

func (i IPhone) read() {
	fmt.Print("speak")
}

func showMySpeak(myphone *IPhone) {
	myphone.speak()
}

func main() {

	showMySpeak(new(IPhone))


}
