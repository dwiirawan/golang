package main

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "localhost:9000"
	accessKeyID := "xr9mdQy4oLoAMybJLWLd"
	secretAccessKey := "s1p1RHyCV9zkQ5BDJJheGBOVbxiRn8nzZFDFY9yr"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup

	// // List Bucket
	// s3Client, err := minio.New(endpoint, &minio.Options{
	// 	Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	// 	Secure: useSSL,
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// buckets, err := s3Client.ListBuckets(context.Background())
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// for _, bucket := range buckets {
	// 	log.Println(bucket)
	// }
}
