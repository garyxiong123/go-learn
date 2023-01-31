package main

import (
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func main() {

	/**
	 * 系统 错误， 不能容错
	 */
	logx.Severef("---Sever")

	/**
	*	业务错误，能错 =》 warn
	 */
	logx.Errorf("----error")

	//panic("123")

	logx.Info("----info")

	logx.Debug("---debug")

	logx.Alert("---Alert")

}

type user struct {
	name string
	age  int
}

//Error, Info, Slow: 将任何类型的信息写进日志，使用 fmt.Sprint(...) 来转换为 string
//Errorf, Infof, Slowf: 将指定格式的信息写入日志
//Errorv, Infov, Slowv: 将任何类型的信息写入日志，用 json marshal 编码
//Errorw, Infow, Sloww: 写日志，并带上给定的 key:value 字段
//WithContext：将给定的 ctx 注入日志信息，例如用于记录 trace-id和span-id
//WithDuration: 将指定的时间写入日志信息中，字段名为 duration
func Test_log_object(t *testing.T) {
	user := &user{
		"gary", 123,
	}
	logx.Info("----info:", user)
	//logx.Info("----info:{} %s:%d", user.name, user.age)
	//Info call has possible formatting directive %s

	logx.Infof("----info:{} %s:--%d", user.name, user.age)

	logx.Infov(user)
	//logx.Infof("----info", user)

	//logx.Infov(user)

	logx.WithDuration(1).Infof("sssss")

}
