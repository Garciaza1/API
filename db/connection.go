package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// Carregar o arquivo .env
	if err := godotenv.Load("env/.env"); err != nil {
		log.Panicf("Error loading .env file: %v", err)
	}

	// Obter as variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Formatar a string de conexão
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", dbHost, dbUser, dbPass, dbName)

	// Tentar abrir a conexão com o banco de dados
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established")
}
