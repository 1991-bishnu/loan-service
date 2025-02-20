package db

import (
	"fmt"
	"log"

	"github.com/1991-bishnu/loan-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Init(conf *config.AppConfig) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Name,
		conf.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	dbInstance = db
	log.Print("Connected to the database")
}

func GetDB() *gorm.DB {
	return dbInstance
}
