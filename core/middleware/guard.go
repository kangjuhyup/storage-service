package middleware

import (
	"fmt"
	"net/http"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// AuthGuard - 인증을 위한 미들웨어 (Guard 역할)
func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("x-auth-token")

		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid auth token"})
			c.Abort()
			return
		}

		redisClient, exists := c.Get("redisClient")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis client not found in context"})
			c.Abort()
			return
		}

		// Redis에서 토큰 검증
		if !isValidToken(redisClient.(*redis.Client), authToken) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func isValidToken(rdb *redis.Client, token string) bool {
	ctx := context.Background()

	val, err := rdb.Get(ctx, "auth:"+token).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		fmt.Println("Redis error:", err)
		return false
	}

	if val == "authenticated" {
		return true
	}
	return false
}
