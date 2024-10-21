package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kangjuhyup/storage-service/config"
	"github.com/kangjuhyup/storage-service/util"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type AuthService struct{}

func (f *AuthService) GetAuth(c *gin.Context, rootPwd string) (string, error) {
	redisClient, exists := c.Get("redisClient")
	if !exists {
		return "", fmt.Errorf("Redis client not found in context")
	}

	rdb := redisClient.(*redis.Client)

	if config.RootPwd == rootPwd {
		authToken := "AUTH_" + util.GenerateRandomString(10)
		err := rdb.Set(ctx, "auth:"+authToken, "authenticated", time.Hour).Err()
		if err != nil {
			return "", err
		}

		return authToken, nil
	} else {
		return "", fmt.Errorf("비밀번호 불일치")
	}
}
