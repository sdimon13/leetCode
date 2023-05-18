package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	envPath := filepath.Join("..", "..", ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	return db, nil
}

func CloseDB() error {
	if db != nil {
		err := db.Close()
		if err != nil {
			return fmt.Errorf("failed to close database connection: %v", err)
		}
	}
	return nil
}
