package main

import (
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func Test_Log_Info_Output(t *testing.T) {

	//panic("123")
	logx.Alert("---Alert")

	logx.Info("----info")

	logx.Debug("---debug")

	logx.Errorf("----error") //业务错误，能错 =》 warn

	logx.Severef("---Sever") //系统 错误， 不能容错
}
