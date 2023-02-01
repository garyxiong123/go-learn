package panic

import (
	"errors"
	"fmt"
	"runtime/debug"
	"testing"
)

func funcA() error {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v", p)
			debug.PrintStack() //?? 敏感
		}
	}()
	return funcB()
}

//if we can use this way to instead of error passing
func funcA_with_Error() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic recover! p:", p)
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			debug.PrintStack()
		}
	}()
	return funcB()
}

func funcB() error {
	// simulation
	panic("foo")
	return errors.New("success")
}

func Test_Panic_No_error(t *testing.T) {
	err := funcA()
	if err == nil {
		fmt.Printf("err is nil\\n")
	} else {
		fmt.Printf("err is %v\\n", err)
	}
}

//panic异常处理机制不会自动将错误信息传递给error，所以要在funcA函数中进行显式的传递
func Test_Panic_with_error(t *testing.T) {
	err := funcA_with_Error()
	if err == nil {
		fmt.Printf("err is nil\\n")
	} else {
		fmt.Printf("err is %v\\n", err)
	}
}
