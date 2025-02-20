package seed

import (
	"database/sql"
	"fmt"

	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// Seed Users
	users := []entity.User{
		{
			BaseModel: entity.BaseModel{ID: "usr_1"},
			Name:      sql.NullString{String: "John Doe", Valid: true},
			Email:     sql.NullString{String: "john@example.com", Valid: true},
		},
		{
			BaseModel: entity.BaseModel{ID: "usr_2"},
			Name:      sql.NullString{String: "Jane Smith", Valid: true},
			Email:     sql.NullString{String: "jane@example.com", Valid: true},
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return fmt.Errorf("could not seed users: %v", err)
		}
	}
	return nil
}
