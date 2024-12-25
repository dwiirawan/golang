package minio

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
)

// func CreateBucket(minioClient *minio.Client, bucketName string) {
// 	ctx := context.Background()
// 	location := "us-east-1"

// 	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
// 	if err != nil {
// 		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
// 		if errBucketExists == nil && exists {
// 			log.Printf("Bucket %s sudah ada\n", bucketName)
// 		} else {
// 			log.Fatalf("Gagal membuat bucket: %v", err)
// 		}
// 	} else {
// 		log.Printf("Bucket %s berhasil dibuat\n", bucketName)
// 	}
// }

func UploadFile(minioClient *minio.Client, bucketName, objectName, filePath string) {
	ctx := context.Background()

	// Mengunggah file ke bucket
	_, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Fatalf("Gagal mengunggah file: %v", err)
	}

	log.Printf("File %s berhasil diunggah ke bucket %s\n", objectName, bucketName)
}

func DownloadFile(minioClient *minio.Client, bucketName, objectName, filePath string) {
	ctx := context.Background()

	// Mengunduh file dari bucket
	err := minioClient.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalf("Gagal mengunduh file: %v", err)
	}

	log.Printf("File %s berhasil diunduh ke %s\n", objectName, filePath)
}
