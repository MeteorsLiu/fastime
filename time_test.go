package fastime

import (
	"testing"
	"time"
)

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

func BenchmarkFastTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Now()
	}
}
