// package main
package config

import (
	"belajar-restful/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var Database_URI = "root:admin123@tcp(localhost:3306)/restful_golang?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(Database_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic("Database tidak terhubung")
	}

	Database.AutoMigrate(&entities.User{})

	return nil
}

/*
func main() {
  Connect()
}
*/
