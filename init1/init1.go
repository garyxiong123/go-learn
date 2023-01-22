package init1

import (
	"fmt"
	"runtime"
)

func init() {
	goos := runtime.GOOS
	println(goos)
	fmt.Println("init1---------------")
}
