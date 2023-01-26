package init_main

import (
	"testing"
)

//2、导包顺序决定init函数执行顺序：
//
//go中不同包中init函数的执行顺序是根据包的导入关系决定的。
func Test_diff_pack_Init(t *testing.T) {

}
