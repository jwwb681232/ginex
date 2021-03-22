package config

import (
	"GinRest/entity"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {
	envErr := godotenv.Load()
	if envErr != nil {
		panic("Failed to load .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbUserName,dbPassword,dbHost,dbPort,dbDatabase)
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	_ = db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB)  {
	dbSql,err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	_ = dbSql.Close()
}