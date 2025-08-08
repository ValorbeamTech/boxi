package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Method 1: Token Bucket Rate Limiter (using golang.org/x/time/rate)
func TokenBucketRateLimit(rps rate.Limit, burst int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rps, burst)
	
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.Header("X-RateLimit-Limit", fmt.Sprintf("%.0f", float64(rps)))
			c.Header("X-RateLimit-Remaining", "0")
			c.Header("Retry-After", "1")
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Method 2: Per-IP Rate Limiter with cleanup
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	// Clean up old entries every minute
	go func() {
		for {
			time.Sleep(time.Minute)
			i.CleanupStaleEntries()
		}
	}()

	return i
}

func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()
	return limiter
}

func (i *IPRateLimiter) CleanupStaleEntries() {
	i.mu.Lock()
	defer i.mu.Unlock()

	for ip, limiter := range i.ips {
		// Remove limiters that haven't been used recently
		if limiter.Tokens() == float64(i.b) {
			delete(i.ips, ip)
		}
	}
}

func IPRateLimit(rateLimiter *IPRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := rateLimiter.GetLimiter(ip)

		if !limiter.Allow() {
			c.Header("X-RateLimit-Limit", fmt.Sprintf("%.0f", float64(rateLimiter.r)))
			c.Header("X-RateLimit-Remaining", "0")
			c.Header("Retry-After", "1")
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":   "Rate limit exceeded",
				"message": "Too many requests from your IP address",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Method 3: Simple Time Window Rate Limiter
type WindowRateLimiter struct {
	requests map[string][]time.Time
	mu       sync.RWMutex
	limit    int
	window   time.Duration
}

func NewWindowRateLimiter(limit int, window time.Duration) *WindowRateLimiter {
	wrl := &WindowRateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}

	// Cleanup old entries
	go func() {
		ticker := time.NewTicker(window)
		defer ticker.Stop()
		for range ticker.C {
			wrl.cleanup()
		}
	}()

	return wrl
}

func (wrl *WindowRateLimiter) Allow(identifier string) bool {
	wrl.mu.Lock()
	defer wrl.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-wrl.window)

	// Clean old requests for this identifier
	requests := wrl.requests[identifier]
	validRequests := make([]time.Time, 0)
	for _, reqTime := range requests {
		if reqTime.After(cutoff) {
			validRequests = append(validRequests, reqTime)
		}
	}

	if len(validRequests) >= wrl.limit {
		return false
	}

	// Add current request
	validRequests = append(validRequests, now)
	wrl.requests[identifier] = validRequests

	return true
}

func (wrl *WindowRateLimiter) cleanup() {
	wrl.mu.Lock()
	defer wrl.mu.Unlock()

	cutoff := time.Now().Add(-wrl.window)
	for identifier, requests := range wrl.requests {
		validRequests := make([]time.Time, 0)
		for _, reqTime := range requests {
			if reqTime.After(cutoff) {
				validRequests = append(validRequests, reqTime)
			}
		}

		if len(validRequests) == 0 {
			delete(wrl.requests, identifier)
		} else {
			wrl.requests[identifier] = validRequests
		}
	}
}

func WindowRateLimit(limiter *WindowRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !limiter.Allow(ip) {
			c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.limit))
			c.Header("X-RateLimit-Window", limiter.window.String())
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":   "Rate limit exceeded",
				"message": fmt.Sprintf("Maximum %d requests per %v", limiter.limit, limiter.window),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Example usage
// func main() {
	// r := gin.Default()

	// Method 1: Simple token bucket (global rate limit)
	// 10 requests per second with burst of 20
	// r.Use(TokenBucketRateLimit(10, 20))

	// Method 2: Per-IP rate limiting
	// ipLimiter := NewIPRateLimiter(5, 10) // 5 RPS per IP, burst of 10
	// r.Use(IPRateLimit(ipLimiter))

	// Method 3: Time window rate limiting
	// windowLimiter := NewWindowRateLimiter(100, time.Minute) // 100 requests per minute
	// r.Use(WindowRateLimit(windowLimiter))

	// Apply rate limiting to specific routes only
	// api := r.Group("/api")
	// {
		// Different rate limit for API endpoints
		// ipLimiter := NewIPRateLimiter(2, 5) // 2 RPS per IP, burst of 5
		// api.Use(IPRateLimit(ipLimiter))

	// 	api.GET("/users", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{"message": "Users endpoint"})
	// 	})

	// 	api.POST("/users", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{"message": "Create user"})
	// 	})
	// }

	// Routes without rate limiting
// 	r.GET("/health", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
// 	})

// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
// 	})

// 	r.Run(":8080")
// }

// Advanced: Rate limiter with Redis backend
/*
import (
	"github.com/go-redis/redis/v8"
	"context"
	"strconv"
)

type RedisRateLimiter struct {
	client *redis.Client
	limit  int
	window time.Duration
}

func NewRedisRateLimiter(client *redis.Client, limit int, window time.Duration) *RedisRateLimiter {
	return &RedisRateLimiter{
		client: client,
		limit:  limit,
		window: window,
	}
}

func (rrl *RedisRateLimiter) Allow(ctx context.Context, key string) (bool, error) {
	pipe := rrl.client.Pipeline()
	
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, rrl.window)
	
	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}
	
	count, err := incr.Result()
	if err != nil {
		return false, err
	}
	
	return count <= int64(rrl.limit), nil
}

func RedisRateLimit(limiter *RedisRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := fmt.Sprintf("rate_limit:%s", c.ClientIP())
		
		allowed, err := limiter.Allow(c.Request.Context(), key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}
		
		if !allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		
		c.Next()
	}
}
*/