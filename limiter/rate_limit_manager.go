package limiter

import "sync"

type RateLimitManager struct {
	Buckets    map[string]*TokenBucket
	Capacity   int // default cap all bucket
	RefillRate int // default refill rate all bucket
	mu         sync.RWMutex
}

func NewRateLimitManager(capacity int, refillRate int) *RateLimitManager {
	return &RateLimitManager{
		Buckets:    make(map[string]*TokenBucket), // empty map
		Capacity:   capacity,
		RefillRate: refillRate,
		mu:         sync.RWMutex{},
	}
}

func (rlm *RateLimitManager) GetOrCreate(ip string) *TokenBucket {
	// Check with read lock
	rlm.mu.RLock()
	bucket, exists := rlm.Buckets[ip]
	rlm.mu.RUnlock()

	// if its exists return bucket

	if exists {
		return bucket
	}

	// if doesn't exists create new bucket
	rlm.mu.Lock()
	defer rlm.mu.Unlock()

	// Double check
	bucket, exists = rlm.Buckets[ip]
	if exists {
		return bucket
	}

	// Create new bucket
	rlm.Buckets[ip] = NewTokenBucket(rlm.Capacity, rlm.RefillRate)

	return rlm.Buckets[ip]

}
