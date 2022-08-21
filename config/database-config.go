package config


import (
	"MBETakeHomeTest/entity"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func SetupDatabaseConnection() *gorm.DB{
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//"host=localhost port=5432 user=postgres password=123456789 dbname=diton sslmode=disable"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",dbHost,dbPort,dbUser,dbPass,dbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // ambang Slow SQL
			LogLevel:                   logger.Silent, // tingkat Log
			IgnoreRecordNotFoundError: true,           // mengabaikan kesalahan ErrRecordNotFound  untuk logger
			Colorful:                  false,          // nonaktifkan warna
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to set connection to database")
	}

	//migrasi model
	db.AutoMigrate(&entity.User{},&entity.UserBalance{},&entity.Transaction{})
	return db
}

func CloseDBConnection(db *gorm.DB)  {
	dbPostgress, err := db.DB()
	if err != nil {
		panic("Failed to close connection from Database")
	}

	dbPostgress.Close()
}