package random

import (
	"io"
	"testing"
)

//Reader是一个全局、共享的密码用强随机数生成器。在Unix类型系统中，
//会从/dev/urandom读取；而Windows中会调用CryptGenRandom API。
var Reader io.Reader

func TestRandom(t *testing.T) {

}
