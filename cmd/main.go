package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kangjuhyup/storage-service/config"
	"github.com/kangjuhyup/storage-service/core/handler"
	"github.com/kangjuhyup/storage-service/core/middleware"
	_ "github.com/kangjuhyup/storage-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Storage Service API
// @version 1.0
// @description This is a storage service API.
// @host localhost:3002
// @BasePath /storage

func main() {
	config.InitConfig()
	router := gin.Default()

	router.Use(middleware.RedisMiddleware())

	storage := router.Group("/storage")
	{
		storage.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// 인증 관련 엔드포인트
		storage.GET("auth", handler.GetAuthSession)

		// 파일박스 관련 엔드포인트
		storage.POST(":box", middleware.AuthGuard(), handler.CreateBox)
		storage.PATCH(":box", middleware.AuthGuard(), handler.UpdateBoxMetadata)
		storage.DELETE(":box", middleware.AuthGuard(), handler.DeleteBox)
		storage.GET(":box", middleware.AuthGuard(), handler.ListFilesInBox)

		// 파일 관련 엔드포인트
		storage.PUT(":box/:file", middleware.AuthGuard(), handler.UploadFile)
		storage.PATCH(":box/:file", middleware.AuthGuard(), handler.UpdateFileMetadata)
		storage.GET(":box/:file", handler.DownloadFile)
		storage.DELETE(":box/:file", middleware.AuthGuard(), handler.DeleteFile)
	}

	router.Run(":3002")
}
