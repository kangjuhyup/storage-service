package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var UploadPath string
var RootPwd string
var RedisHost string
var RedisPort int

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env 파일 호출 실패: %v", err)
	}

	UploadPath = os.Getenv("UPLOAD_PATH")
	RootPwd = os.Getenv("ROOT_PWD")
	RedisHost = os.Getenv("REDIS_HOST")

	redisPortStr := os.Getenv("REDIS_PORT")
	RedisPort, err = strconv.Atoi(redisPortStr)
	if err != nil {
		log.Fatalf("REDIS_PORT 값을 정수로 변환할 수 없습니다: %v", err)
	}

	// 필수 환경 변수 확인
	if UploadPath == "" {
		log.Fatalf("UPLOAD_PATH 세팅 없음")
	}

	if RootPwd == "" {
		log.Fatalf("ROOT_PWD 세팅 없음")
	}

	if RedisHost == "" {
		log.Fatalf("REDIS_HOST 세팅 없음")
	}

	if RedisPort == 0 {
		log.Fatalf("REDIS_PORT 세팅 없음")
	}
}
