package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kangjuhyup/storage-service/config"
	"github.com/kangjuhyup/storage-service/core/handler"
	"github.com/kangjuhyup/storage-service/core/middleware"
)

func main() {

	config.InitConfig()
	router := gin.Default()

	router.Use(middleware.RedisMiddleware())

	// 인증 관련 엔드포인트
	router.GET("auth", handler.GetAuthSession)

	// 파일박스 관련 엔드포인트
	router.POST("storage/:box", middleware.AuthGuard(), handler.CreateBox)
	router.PATCH("storage/:box", middleware.AuthGuard(), handler.UpdateBoxMetadata)
	router.DELETE("storage/:box", middleware.AuthGuard(), handler.DeleteBox)
	router.GET("storage/:box", middleware.AuthGuard(), handler.ListFilesInBox)

	// 파일 관련 엔드포인트
	router.PUT("storage/:box/:file", middleware.AuthGuard(), handler.UploadFile)
	router.PATCH("storage/:box/:file", middleware.AuthGuard(), handler.UpdateFileMetadata)
	router.GET("storage/:box/:file", handler.DownloadFile)
	router.DELETE("storage/:box/:file", middleware.AuthGuard(), handler.DeleteFile)

	router.Run(":3002")
}
