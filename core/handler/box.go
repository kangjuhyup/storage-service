package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kangjuhyup/storage-service/core/service"
)

var boxService = service.BoxService{}

// CreateBox - 파일박스 생성
func CreateBox(c *gin.Context) {
	box := c.Param("box")
	if err := boxService.CreateBox(box); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create box"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Box created", "box": box})
}

// UpdateBoxMetadata - 파일박스 메타데이터 수정 (간단히 성공 메시지만)
func UpdateBoxMetadata(c *gin.Context) {
	box := c.Param("box")
	// 실제로 메타데이터를 처리하는 로직 추가 가능
	c.JSON(http.StatusOK, gin.H{"message": "Box metadata updated", "box": box})
}

// DeleteBox - 파일박스 삭제
func DeleteBox(c *gin.Context) {
	box := c.Param("box")

	if err := boxService.DeleteBox(box); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete box"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Box deleted", "box": box})
}

// ListFilesInBox - 파일박스 내 파일 리스트 조회
func ListFilesInBox(c *gin.Context) {
	box := c.Param("box")

	fileList, err := boxService.ListFilesInBox(box)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list files in box"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": fileList})
}
