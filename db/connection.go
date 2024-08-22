package db

import (
	"fmt"
	"os"
	"time"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Init() {
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")

	dbUri := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPass)
	time.Sleep(10 * time.Second)
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrations_err := db.AutoMigrate(&Reports{}, &NerResults{})
	if migrations_err != nil {
		panic("failed migrate to database")
	}

	Conn = db
}
