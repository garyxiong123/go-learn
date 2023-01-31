package unitTest

import (
	"testing"
)

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		println(i)
	}
	return
}

func BenchmarkBigLen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		println(i)
	}
}
