package hash_digist

import (
	"crypto/sha256"
	"testing"
)

func Test_Sha256(t *testing.T) {

	// 预加密数据
	text := "hello dalgurak"
	// hash签名
	hashText := sha256.Sum256([]byte(text))

	println(hashText)

}
