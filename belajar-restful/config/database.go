// package main
package config

import (
	"belajar-restful/entities"

	// MySQL
	// "gorm.io/driver/mysql"

	// Postgresql
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

// Mysql
// var Database_URI = "root:admin123@tcp(localhost:3306)/restful_golang?charset=utf8mb4&parseTime=True&loc=Local"
// Postgress
var dsn = "host=localhost user=postgres password=admin dbname=restful_golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func Connect() error {
	var err error

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic("Database tidak terhubung")
	}

	Database.AutoMigrate(&entities.User{})

	return nil
}
