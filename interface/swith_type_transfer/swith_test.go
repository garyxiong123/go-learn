package swith_type_transfer

import (
	"github.com/garyxiong123/go-learn/go-zero-demo/greet/common/errorx"
	"testing"
)

func TestSwitch(t *testing.T) {

	err := getErr()
	switch e := err.(type) {
	case *errorx.CodeError:
		println(e.Data())
	default:

	}
}

func getErr() error {

	return nil
}
