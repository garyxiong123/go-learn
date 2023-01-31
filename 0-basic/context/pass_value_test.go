package context

import (
	"context"
	"fmt"
	"testing"
)

//context 用来解决 goroutine 之间退出通知、元数据传递的功能

//context.WithValue(): 从父上下文中创建一个子上下文，传值的子上下文使用 context.valueCtx。
//单线程，能否跨线程
func Test_Pass_Value(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "gary")
	name := ctx.Value("name").(string)

	println(name)

}

func Test_Pass_Value_double(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "gary")
	ctx = context.WithValue(ctx, "age", "123")
	name := ctx.Value("name").(string)

	println(name)

}

func Test_Pass_Value_1(t *testing.T) {
	ctx := context.Background()
	process(ctx)

	ctx = context.WithValue(ctx, "traceId", "qcrao-2019")
	process(ctx)

}

func process(ctx context.Context) {
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("process over. trace_id=%s\n", traceId)
	} else {
		fmt.Printf("process over. no trace_id\n")
	}
}

func Test_Pass_Value_Skip_Thread(t *testing.T) {

}
