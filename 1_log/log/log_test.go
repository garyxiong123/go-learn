package main

import (
	"fmt"
	"log"
	"testing"
)

func Test_1(t *testing.T) {
	log.Println("----normal")
	log.Fatalln("----fatal")

	log.Println("----normal")
	log.Println("----fatal")

	fmt.Println("sss")

	log.Panic("----fatal")

}
