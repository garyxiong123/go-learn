package sign

import (
	"testing"
)

func TestEncrypto1112(t *testing.T) {
	key, err := NewSigningKey()
	if err != nil {
		return
	}
	println(key)

	message := []byte("Hello world")
	println(message)

}
