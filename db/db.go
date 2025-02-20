package db

import (
	"fmt"

	"github.com/1991-bishnu/loan-service/config"
	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Init(conf *config.AppConfig) error {
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
		return fmt.Errorf("failed to connect to the DB error: %w", err)
	}
	dbInstance = db
	return nil
}

func GetDB() *gorm.DB {
	return dbInstance
}

func MigrateDB() error {
	err := dbInstance.AutoMigrate(
		&entity.Document{},
		&entity.Employee{},
		&entity.Investor{},
		&entity.Investment{},
		&entity.Loan{},
		&entity.User{},
	)
	if err != nil {
		return fmt.Errorf("DB Migration failed error: %w", err)
	}
	return nil
}
