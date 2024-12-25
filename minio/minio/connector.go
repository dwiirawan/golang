// minio/connector.go

package minio

import (
	"context"
	"fmt"
	config "minio/configs"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var s3Client *minio.Client

func InitMinio() {

	// Requests are always secure (HTTPS) by default.
	// Set secure=false to enable insecure (HTTP) access.
	// This boolean value is the last argument for New().
	conn, err := minio.New(config.MinioEndpt, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		fmt.Println(err)
	}

	found, err := conn.BucketExists(context.Background(), config.MinioBucket)
	if err != nil {
		fmt.Println(err)
	}
	if found {
		fmt.Println("Connection to Minio successful.")
		fmt.Println("Endpoint: " + config.MinioEndpt)
		fmt.Println("Bucket  : " + config.MinioBucket)
	}

	s3Client = conn
}

func GetMinio() *minio.Client {
	if s3Client == nil {
		InitMinio()
	}
	return s3Client
}

func HandlePanic() {
	r := recover()

	if r != nil {
		fmt.Println("RECOVER :", r)
	}
}

func DownloadFileFromMinio(objectname string, filePath string) error {
	defer HandlePanic()

	// Download and save the object as a file in the local filesystem.
	err := GetMinio().FGetObject(context.Background(), config.MinioBucket, objectname, filePath, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetPresignedURLFromMinio(objectname string) string {
	defer HandlePanic()
	reqParams := make(url.Values)

	// Gernerate presigned get object url.
	presignedURL, err := GetMinio().PresignedGetObject(context.Background(), config.MinioBucket, objectname, time.Second*24*60*60, reqParams)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return presignedURL.String()
}

func UploadFileInMinio(objectname string, filePath string, contentType string) string {
	defer HandlePanic()

	// Upload the test file with FPutObject
	info, err := GetMinio().FPutObject(context.Background(), config.MinioBucket, objectname, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", objectname, info.Size)
	return info.ETag
}
