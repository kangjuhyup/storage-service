package service

import (
	"os"
	"path/filepath"
)

// BoxService - 파일박스 관련 서비스
type BoxService struct{}

// CreateBox - 파일박스 생성 로직
func (b *BoxService) CreateBox(boxName string) error {
	boxPath := filepath.Join(uploadPath, boxName)
	return os.MkdirAll(boxPath, 0755)
}

// DeleteBox - 파일박스 삭제 로직
func (b *BoxService) DeleteBox(boxName string) error {
	boxPath := filepath.Join(uploadPath, boxName)
	return os.RemoveAll(boxPath)
}

// ListFilesInBox - 파일박스 내 파일 리스트 반환
func (b *BoxService) ListFilesInBox(boxName string) ([]string, error) {
	boxPath := filepath.Join(uploadPath, boxName)
	files, err := os.ReadDir(boxPath)
	if err != nil {
		return nil, err
	}

	var fileList []string
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}

	return fileList, nil
}
