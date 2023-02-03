package skip_package

import (
	"fmt"
)

func ExampleHello() {
	fmt.Println("hello")
	// Output: hello1
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}

func ExamplePerm() {
	//for _, value := range Perm(5) {
	//	fmt.Println(value)
	//}
	// Unordered output: 4
	// 2
	// 1
	// 3
	// 0
}
