package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	// Load local environment file first if present, then fall back to .env.
	_ = godotenv.Load(".env.local")

	// Load default .env file if available.
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using environment variables")
	} else {
		log.Println("Loaded .env")
	}

	// Read environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" {
		dbHost = "localhost"
		log.Println("DB_HOST not set; defaulting to localhost")
	}

	log.Println("DB_HOST:", dbHost)
	log.Println("DB_PORT:", dbPort)
	log.Println("DB_NAME:", dbName)
	log.Println("DB_USER:", dbUser)
	log.Println("DB_PASSWORD:", dbPassword)

	// Build DSN
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	// Open DB connection
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Verify connection
	for i := 0; i < 10; i++ {
		err = database.Ping()

		if err == nil {
			break
		}

		log.Println("Waiting for MySQL...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal(err)
	}

	DB = database

	log.Println("Connected to MySQL")
}

func CreateTable() {
	log.Println("Creating MySQL table games")

	query := `
	CREATE TABLE IF NOT EXISTS games (
		id INT AUTO_INCREMENT PRIMARY KEY,
		player_choice VARCHAR(20),
		computer_choice VARCHAR(20),
		result VARCHAR(20),
		played_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Games table ready")
}
