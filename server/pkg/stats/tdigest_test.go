package stats

import (
	"math/rand"
	"testing"
)

func BenchmarkTDigest_Add(b *testing.B) {
	digest := NewTDigest(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := rand.Float64()
		digest.Add(r)
	}
}

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rand.Float64()
	}
}
