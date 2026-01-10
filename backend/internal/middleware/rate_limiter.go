package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

// RateLimiter implements a simple in-memory rate limiter using token bucket algorithm
type RateLimiter struct {
	requests map[string]*clientInfo
	mu       sync.RWMutex
	rate     int           // requests per duration
	duration time.Duration // time window
	cleanup  time.Duration // cleanup interval for old entries
}

type clientInfo struct {
	tokens    int
	lastReset time.Time
}

// NewRateLimiter creates a new rate limiter
// rate: number of requests allowed per duration
// duration: time window for rate limiting
func NewRateLimiter(rate int, duration time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string]*clientInfo),
		rate:     rate,
		duration: duration,
		cleanup:  duration * 2,
	}

	// Start cleanup goroutine
	go rl.cleanupRoutine()

	return rl
}

// cleanupRoutine periodically removes old entries
func (rl *RateLimiter) cleanupRoutine() {
	ticker := time.NewTicker(rl.cleanup)
	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for key, info := range rl.requests {
			if now.Sub(info.lastReset) > rl.duration*2 {
				delete(rl.requests, key)
			}
		}
		rl.mu.Unlock()
	}
}

// Allow checks if a request from the given key should be allowed
func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	info, exists := rl.requests[key]

	if !exists {
		rl.requests[key] = &clientInfo{
			tokens:    rl.rate - 1,
			lastReset: now,
		}
		return true
	}

	// Reset tokens if duration has passed
	if now.Sub(info.lastReset) >= rl.duration {
		info.tokens = rl.rate - 1
		info.lastReset = now
		return true
	}

	// Check if tokens available
	if info.tokens > 0 {
		info.tokens--
		return true
	}

	return false
}

// RateLimitMiddleware returns an Echo middleware function for rate limiting
func RateLimitMiddleware(limiter *RateLimiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Use IP address as the key
			key := c.RealIP()

			if !limiter.Allow(key) {
				return c.JSON(http.StatusTooManyRequests, map[string]interface{}{
					"success": false,
					"error":   "too many requests, please try again later",
					"status":  http.StatusTooManyRequests,
				})
			}

			return next(c)
		}
	}
}

// RateLimitByRouteMiddleware creates a rate limiter specific to a route
// Useful for applying different limits to different endpoints
func RateLimitByRouteMiddleware(rate int, duration time.Duration) echo.MiddlewareFunc {
	limiter := NewRateLimiter(rate, duration)
	return RateLimitMiddleware(limiter)
}

// Predefined rate limiters for common use cases
var (
	// LoginRateLimiter: 5 attempts per minute per IP
	LoginRateLimiter = NewRateLimiter(5, time.Minute)

	// RegisterRateLimiter: 3 registrations per minute per IP
	RegisterRateLimiter = NewRateLimiter(3, time.Minute)

	// PasswordResetRateLimiter: 3 requests per 5 minutes per IP
	PasswordResetRateLimiter = NewRateLimiter(3, 5*time.Minute)

	// APIRateLimiter: 100 requests per minute per IP (general API)
	APIRateLimiter = NewRateLimiter(100, time.Minute)

	// RefreshTokenRateLimiter: 10 refreshes per minute per IP
	RefreshTokenRateLimiter = NewRateLimiter(10, time.Minute)
)
