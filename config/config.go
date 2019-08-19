package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DbConfig struct {
	DbType     string
	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DatabaseConfig() DbConfig {
	return DbConfig{
		DbType:     os.Getenv("DB_TYPE"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbDatabase: os.Getenv("DB_DATABASE"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
}
