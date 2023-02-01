package orm

import (
	"testing"
)

// 非内置的结构体就需要通过标签embedded 将其嵌入，并且还可以通过标签 embeddedPrefix 来为 db 中的字段名添加前缀：
func Test_embedded(t *testing.T) {
	//basic.Blog对象
}
