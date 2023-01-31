package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/**
** context包是Go 语言中用来设置截止日期、同步信号，传递请求相关值的结构体，是开发常用的并发控制技术。
 */
func Test_Context(t *testing.T) {

}

//Deadline(): 工作的截止时间，没有设置deadline则ok==false。
//Done(): 需要在select-case语句中使用（case <-context.Done(): ）。当context被关闭后，Done()返回一个被关闭的通道（关闭的通道依然是可以读的
//，所以goroutine可以收到关闭请求）；当context还未关闭时，Done()返回nil。
//Err(): 描述context关闭的原因，其原因由context实现控制。例如：因deadline关闭：context deadline exceeded；因主动关闭：context canceled。没有关闭时，返回nil。
//Value(): 特别的用于一种context：不用于控制呈树状分布的goroutine，而是用于在树状分布的goroutine之间传递信息。Value()方法根据key值查询map中的Value。

//它做了2件事：1. 创建过期时间为1s的上下文。 2. 将context传入handle函数中，函数使用500ms的时间处理请求

func Test_TimeOut_Cancel(t *testing.T) {
	_, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// releases resources if slowOperation completes before timeout elapsess
	defer cancel()

}

func Test_No_Out_Of_Date(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

	//process request with 500ms
	//main context deadline exceeded

}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

// context过期时间为1s，处理时间为0.5秒(select中的过期时间)，函数有足够的时间完成处理，
//也就是<-time.After(duration):会在<-ctx.Done()之前完成，故输出process request with 500ms。
//再过0.5s，<-ctx.Done()完成，这时候输出main context deadline exceeded。
//
func Test_Out_Of_Date(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 5000*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

}
