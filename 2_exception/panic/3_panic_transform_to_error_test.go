package panic

import (
	"errors"
	"fmt"
	"runtime/debug"
	"testing"
)

//我们在调用recover的延迟函数中以最合理的方式响应该异常：
//

//panic异常处理机制不会自动将错误信息传递给error，所以要在funcA函数中进行显式的传递
func Test_Panic_transfer_2_error(t *testing.T) {
	err := errorPass()
	if err == nil {
		fmt.Printf("err is nil\\n")
	} else {
		fmt.Printf("err is %v\\n", err)
	}
}

//显式的传递Error
func errorPass() (err error) {
	defer func() {
		//todo recover bad inflowence
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
	return throwPanicWithError()
}

func throwPanicWithError() error {
	// simulation
	panic("foo")
	return errors.New("success")
}

func Test_Panic_No_error(t *testing.T) {
	err := errorNoPass()
	if err == nil {
		fmt.Printf("err is nil\\n")
	} else {
		fmt.Printf("err is %v\\n", err)
	}
}

func errorNoPass() error {

	defer func() {

		if p := recover(); p != nil {
			//打印堆栈的异常调用信息和关键的业务信息，以便这些问题保留可见；
			fmt.Printf("panic recover! p: %v", p)
			debug.PrintStack() //?? 敏感

			//将异常转换为错误，以便调用者让程序恢复到健康状态并继续安全运行
		}
	}()

	return throwPanicWithError()
}
