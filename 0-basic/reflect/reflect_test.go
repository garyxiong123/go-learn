package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reflect(t *testing.T) {

	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
