package debug

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"testing"
)

func Test_Debug(t *testing.T) {

	println("gary-debug", debug.Stack())
	println("gary-PrintStack", debug.PrintStack)
}

func Test_Debug_Stack_Trace(t *testing.T) {
	defer HandleException()
	num := 0
	fmt.Println(1 / num)
}

func Test_Debug_Stack_Trace_Stack(t *testing.T) {
	defer HandleException1()
	num := 0
	fmt.Println(1 / num)
}

func HandleException1() {
	errs := recover()
	if errs == nil {
		return
	}
	fmt.Println(string(debug.Stack()))
	fmt.Println(errs)

}

//test in case of double trigger
func Test_Debug_Stack_Trace_Job(t *testing.T) {

}

func HandleException() {
	errs := recover()
	if errs == nil {
		return
	}
	var stackBuf [1024]byte
	stackBufLen := runtime.Stack(stackBuf[:], false)
	fmt.Printf("==> %s\n", string(stackBuf[:stackBufLen]))
}
