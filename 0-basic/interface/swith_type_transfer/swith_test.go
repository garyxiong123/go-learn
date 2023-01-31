package swith_type_transfer

import (
	"github.com/garyxiong123/go-learn/web/go-zero/common/errorx"
	"testing"
)

func TestSwitch(t *testing.T) {

	err := getErr()
	switch e := err.(type) {
	case *errorx.CodeError:
		println(e.Data())
	case *SystemError:
		//println(e.Data())  //case 这个地方会出现类型推导，和强制类型转换
		println(e.Data1())
	default:

	}
}

func getErr() error {

	return nil
}
