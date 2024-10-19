package service

import (
	"os"
	"path/filepath"
)

const uploadPath = "./storage"

// FileService - 파일 관련 서비스
type FileService struct{}

// UploadFile - 파일 업로드 로직
func (f *FileService) UploadFile(boxName, fileName string, fileContent []byte) error {
	boxPath := filepath.Join(uploadPath, boxName)
	filePath := filepath.Join(boxPath, fileName)

	// 파일박스가 없으면 생성
	if err := os.MkdirAll(boxPath, 0755); err != nil {
		return err
	}

	// 파일 저장
	return os.WriteFile(filePath, fileContent, 0644)
}

// DeleteFile - 파일 삭제 로직
func (f *FileService) DeleteFile(boxName, fileName string) error {
	filePath := filepath.Join(uploadPath, boxName, fileName)
	return os.Remove(filePath)
}

// GetFilePath - 파일 경로 반환
func (f *FileService) GetFilePath(boxName, fileName string) (string, error) {
	filePath := filepath.Join(uploadPath, boxName, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", err
	}
	return filePath, nil
}
