package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	host = "127.0.0.1"
	port = "3306"
	user = "root"
	password = "fpf"
	DBName = "test"

	return gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, port, DBName,
		),
	)
}
