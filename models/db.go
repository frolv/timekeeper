package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var db *gorm.DB

func init() {
	var err error

	params := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"))

	db, err = gorm.Open("postgres", params)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	if os.Getenv("APP_ENV") == "development" {
		db.LogMode(true)
	}

	fmt.Println("Database connection initialized")
	performMigrations()
}

func performMigrations() {
	db.AutoMigrate(&Account{})
}
