// configs/config.go

package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (

	// Minio
	MinioEndpt     = getEnv("MINIO_ENDPOINT", "")
	MinioBucket    = getEnv("MINIO_BUCKET", "")
	MinioAccessKey = getEnv("MINIO_ACCESS_KEY", "")
	MinioSecretKey = getEnv("MINIO_SECRET_KEY", "")

	// Similarly, We can fetch more env variables here,
	// for API keys, Database credentials etc
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
