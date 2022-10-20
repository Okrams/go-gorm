package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database string
var DSN = "host=localhost user=postgres password=scrack123 dbname=go-gorm-example port=5434"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB Connected")
	}
}
