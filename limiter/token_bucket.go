package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	MaxTokenLimit   int
	CurrentToken    int
	RefillTokenRate int
	LastRefillTime  time.Time
	Mu              sync.Mutex
}

func NewTokenBucket(capacity int, refillRate int) *TokenBucket {
	return &TokenBucket{
		MaxTokenLimit:   capacity,
		CurrentToken:    capacity,
		RefillTokenRate: refillRate,
		LastRefillTime:  time.Now(),
		Mu:              sync.Mutex{},
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.Mu.Lock()
	defer tb.Mu.Unlock()

	elapsed := time.Since(tb.LastRefillTime).Seconds()
	tokensToAdd := int(elapsed * float64(tb.RefillTokenRate))

	tb.CurrentToken = min(tb.CurrentToken+tokensToAdd, tb.MaxTokenLimit)
	tb.LastRefillTime = time.Now()

	if tb.CurrentToken > 0 {
		tb.CurrentToken--
		return true
	}
	return false
}
