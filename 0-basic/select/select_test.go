package _select

import (
	"fmt"
	"testing"
)

func TestSelect_basic(t *testing.T) {
	println("start")
	select {}

	//fatal error: all goroutines are asleep - deadlock!
	println("end")
}

func TestSelect_basic_AddThread(t *testing.T) {

}

//In this example, the select statement is waiting for either channel ch1 or ch2 to be ready for communication. When one of the channels is ready, the corresponding case statement is executed and the value received from the channel is printed.
//
//You can also have a default case in the select statement which is executed if none of the channels are ready for communication.
//
//It is useful when you have multiple channels and you want to listen from multiple channels simultaneously or if you want to listen from one channel with a timeout.
//
//Please note that this is just an example and it is not an exhaustive usage of select statement, it can be used in more complex scenarios as well.
//
//select
func TestSelect(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- "hello"
	}()

	select {
	case i := <-ch1:
		fmt.Println("received", i, "from ch1")
	case s := <-ch2:
		fmt.Println("received", s, "from ch2")
	}
}
