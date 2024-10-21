package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kangjuhyup/storage-service/config"
	"github.com/redis/go-redis/v9"
)

// RedisMiddleware - Redis 클라이언트를 설정하고 context에 추가하는 미들웨어
func RedisMiddleware() gin.HandlerFunc {
	// Redis 클라이언트 초기화
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + strconv.Itoa(config.RedisPort), // Redis 서버 주소
		Password: "",                                                      // Redis 비밀번호
		DB:       0,                                                       // 기본 DB
	})

	return func(c *gin.Context) {
		// Redis 클라이언트를 요청의 context에 저장
		c.Set("redisClient", rdb)

		// 요청을 다음 미들웨어 또는 핸들러로 전달
		c.Next()
	}
}
