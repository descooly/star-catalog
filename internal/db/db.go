package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/descooly/star-catalog/internal/service"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file found, using environment variables, error: %v", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("missing required database environment variables")
	}

	constr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	var err error

	for i := 0; i < 30; i++ {
		DB, err = gorm.Open(postgres.Open(constr), &gorm.Config{})
		if err == nil {
			log.Println("successfully connected to database")
			break
		}
		log.Printf("waiting for database %v", err)
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after 30 attempts: %w", err)
	}

	if err := DB.AutoMigrate(&service.StarInfo{}, &service.Star{}); err != nil {
		log.Printf("couldnt migrate: %v", err)
	}

	return DB, nil
}
