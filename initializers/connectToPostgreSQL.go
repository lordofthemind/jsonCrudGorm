package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPostgreSQL() {
	var err error
	// connectionString := fmt.Sprintf(
	// 	"host=postgresql user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	// os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PORT"),
	// )
	connectionString := os.Getenv("postgreSQLConnectionURL")
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
	fmt.Println("Connected to database")
}
