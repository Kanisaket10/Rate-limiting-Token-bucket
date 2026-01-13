package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   int
	tokens     int
	refillRate int // tokens per second
	lastRefill time.Time
	mu         sync.Mutex
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	newTokens := int(elapsed * float64(tb.refillRate))

	if newTokens > 0 {
		tb.tokens += newTokens
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
		tb.lastRefill = now
	}

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

// Global limiter
var limiter = NewTokenBucket(4, 2)

func rateLimiter(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		if !limiter.Allow() {
			message := Message{
				Status: "Request Failed",
				Body:   "Too many requests, try again later.",
			}

			writer.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(writer).Encode(message)
			return
		}

		next(writer, request)
	}
}
