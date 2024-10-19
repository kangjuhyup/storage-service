package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kangjuhyup/storage-service/core/handler"
)

func main() {
	router := gin.Default()

	// 파일박스 관련 엔드포인트
	router.POST("/:box", handler.CreateBox)
	router.PATCH("/:box", handler.UpdateBoxMetadata)
	router.DELETE("/:box", handler.DeleteBox)
	router.GET("/:box", handler.ListFilesInBox)

	// 파일 관련 엔드포인트
	router.PUT("/:box/:file", handler.UploadFile)
	router.PATCH("/:box/:file", handler.UpdateFileMetadata)
	router.GET("/:box/:file", handler.DownloadFile)
	router.DELETE("/:box/:file", handler.DeleteFile)

	// 서버 실행
	router.Run(":8080")
}
