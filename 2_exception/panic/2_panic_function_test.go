package panic

import (
	"errors"
	"fmt"
	"testing"
)

func throwPanic() error {
	// simulation
	panic("foo")
	return errors.New("success")
}

func Test_Panic_No_Recover(t *testing.T) {
	err := throwPanic()
	if err == nil {
		fmt.Printf("err is nil\\n")
	} else {
		fmt.Printf("err is %v\\n", err)
	}
}
