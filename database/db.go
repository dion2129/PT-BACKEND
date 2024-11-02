package database

import (
	"api-test/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	// err error
)
func ConnectToDb() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    dbPort := os.Getenv("DB_PORT")
    dbHost := os.Getenv("DB_HOST")
    dbUser  := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Memperbaiki format DSN untuk MySQL
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser , dbPassword, dbHost, dbPort, dbName)
    
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        logrus.Fatal(err)
        panic(err)
    }

    db.AutoMigrate(&models.Club{}, &models.Event{}, &models.User{})

    logrus.Println("db is connected")

    return db
}
