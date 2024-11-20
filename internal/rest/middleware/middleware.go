package middleware

import (
	"net"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	visitors = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 5) // 초당 1요청, 최대 5버스트
		visitors[ip] = limiter

		// IP 삭제를 위한 타이머
		go func() {
			time.Sleep(1 * time.Hour)
			mu.Lock()
			delete(visitors, ip)
			mu.Unlock()
		}()
	}
	return limiter
}

func RateLimitMiddleware(c *gin.Context) {
	ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	limiter := getVisitor(ip)

	if !limiter.Allow() {
		c.JSON(429, gin.H{"error": "Too many requests"})
		c.Abort()
		return
	}
	c.Next()
}
