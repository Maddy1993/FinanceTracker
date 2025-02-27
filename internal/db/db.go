package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB wraps a gorm.DB pointer
type DB struct {
	Conn *gorm.DB
}

// ConnectPostgres opens a connection to a Postgres DB
func ConnectPostgres() (*DB, error) {
	// For local dev, you might set environment variables like:
	//   export DB_HOST=localhost
	//   export DB_PORT=5432
	//   export DB_USER=postgres
	//   export DB_PASS=mysecret
	//   export DB_NAME=myapp_dev

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASS", "mysecret")
	name := getEnv("DB_NAME", "myapp_dev")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgres successfully.")

	return &DB{
		Conn: dbConn,
	}, nil
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
