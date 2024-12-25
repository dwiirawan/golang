package main

import (
	"log"
	config "minio/configs"
	"minio/minio"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bucketName := config.MinioBucket
	objectName := "example.txt"
	filePath := "example.txt"

	minio.GetMinio()
	s3Client := minio.GetMinio()

	// Mengunggah file
	minio.UploadFile(s3Client, bucketName, objectName, filePath)

	// Mengunduh file
	minio.DownloadFile(s3Client, bucketName, objectName, "readme.txt")
}
