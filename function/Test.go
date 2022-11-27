package main

import (
	"fmt"
)

func main() {

	sum := Add_public(3, 4)
	sum1 := add_private(3, 4)

	fmt.Errorf(string(sum))
	fmt.Errorf(string(sum1))

}
