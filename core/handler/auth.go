package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kangjuhyup/storage-service/core/service"
)

var authService = service.AuthService{}

func GetAuthSession(c *gin.Context) {
	rootPwd := c.GetHeader("x-root-pwd")
	if rootPwd == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "x-root-pwd header is missing"})
		return
	}

	// AuthService를 통해 인증 처리
	authToken, err := authService.GetAuth(c, rootPwd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if authToken != "" {
		// 인증 성공 시 토큰을 헤더에 설정
		c.Header("x-auth-token", authToken)
		// 인증 성공 응답
		c.JSON(http.StatusOK, gin.H{"message": "Authenticated"})
	} else {
		// 인증 실패시 메시지 반환
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
	}
}
