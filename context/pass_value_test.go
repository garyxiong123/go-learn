package context

import (
	"context"
	"fmt"
	"testing"
)

//context.WithValue(): 从父上下文中创建一个子上下文，传值的子上下文使用 context.valueCtx。
//单线程，能否跨线程
func Test_Pass_Value(t *testing.T) {

	context.WithValue(context.Background(), "name", "gary")
	context.WithValue(context.Background(), "age", "33")
	name := context.Background().Value("name")

	println(name)
	context.TODO().Value("ddd")
	println(context.TODO())
	fmt.Println(context.TODO())

}
