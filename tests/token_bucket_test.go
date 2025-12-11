package tests

import (
	"GoRateLimit/limiter"
	"sync"
	"testing"
	"time"
)

func TestAllow(t *testing.T) {
	tb := limiter.NewTokenBucket(5, 1)

	for i := 0; i < 5; i++ {
		if !tb.Allow() {
			t.Errorf("Allow() should return true")
		}
	}

	if tb.Allow() {
		t.Errorf("Allow() should return false")
	}
}

func TestRefill(t *testing.T) {
	tb := limiter.NewTokenBucket(10, 5)

	//Allow 10 token
	for i := 0; i < 10; i++ {
		tb.Allow()
	}

	//token finished
	if tb.Allow() {
		t.Errorf("Allow() should return false")
	}

	//wait 2 seconds token should be refill
	time.Sleep(2 * time.Second)

	//token should be refill
	for i := 0; i < 10; i++ {
		if !tb.Allow() {
			t.Errorf("Request %d should pass after refill", i+1)
		}
	}
}

func TestConcurrency(t *testing.T) {
	tb := limiter.NewTokenBucket(50, 10)

	var wg sync.WaitGroup
	successCount := 0
	var mu sync.Mutex //successCount should be thread safe

	//100 goroutine send request same time
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if tb.Allow() {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	//100 request but token limit is 50 so only 50 request should pass
	if successCount != 50 {
		t.Errorf("Allow() should return true")
	}
}
