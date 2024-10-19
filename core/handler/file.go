package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kangjuhyup/storage-service/core/service"
)

var fileService = service.FileService{}

// UploadFile - 파일 업로드 핸들러
func UploadFile(c *gin.Context) {
	box := c.Param("box")
	file := c.Param("file")

	uploadedFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get uploaded file"})
		return
	}

	fileContent, err := uploadedFile.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer fileContent.Close()

	// TODO : Metadata 를 읽어 파일 크기 제한 등 검사

	fileBytes := make([]byte, uploadedFile.Size)
	if _, err := fileContent.Read(fileBytes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	if err := fileService.UploadFile(box, file, fileBytes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": file})
}

// UpdateFileMetadata - 파일 메타데이터 수정 핸들러
func UpdateFileMetadata(c *gin.Context) {
	box := c.Param("box")
	file := c.Param("file")
	c.JSON(http.StatusOK, gin.H{"message": "File metadata updated", "box": box, "file": file})
}

// DownloadFile - 파일 다운로드 핸들러
func DownloadFile(c *gin.Context) {
	box := c.Param("box")
	file := c.Param("file")

	filePath, err := fileService.GetFilePath(box, file)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.File(filePath)
}

// DeleteFile - 파일 삭제 핸들러
func DeleteFile(c *gin.Context) {
	box := c.Param("box")
	file := c.Param("file")

	if err := fileService.DeleteFile(box, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted", "file": file})
}
