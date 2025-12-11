package benchmark

import (
	"GoRateLimit/limiter"
	"fmt"
	"testing"
)

func BenchmarkSingleAllow(b *testing.B) {
	tb := limiter.NewTokenBucket(100, 100)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tb.Allow()
	}
}

func BenchmarkConcurrentAllow(b *testing.B) {
	tb := limiter.NewTokenBucket(1000000, 100)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tb.Allow()
		}
	})
}

func BenchmarkGetOrCreate(b *testing.B) {
	rlm := limiter.NewRateLimitManager(100, 100)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rlm.GetOrCreate(fmt.Sprintf("%d", i))
	}
}
