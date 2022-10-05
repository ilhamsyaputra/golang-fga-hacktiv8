package database

import (
	"assignment-2/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONFIG")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
