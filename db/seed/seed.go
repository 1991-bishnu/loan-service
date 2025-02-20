package seed

import (
	"database/sql"
	"fmt"

	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
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

	employees := []entity.Employee{
		{
			BaseModel: entity.BaseModel{ID: "emp_1"},
			Name:      sql.NullString{String: "John Doe", Valid: true},
			Email:     sql.NullString{String: "john@example.com", Valid: true},
		},
		{
			BaseModel: entity.BaseModel{ID: "emp_2"},
			Name:      sql.NullString{String: "Jane Smith", Valid: true},
			Email:     sql.NullString{String: "jane@example.com", Valid: true},
		},
	}
	for _, employee := range employees {
		if err := db.Create(&employee).Error; err != nil {
			return fmt.Errorf("could not seed employees: %v", err)
		}
	}

	investors := []entity.Investor{
		{
			BaseModel: entity.BaseModel{ID: "invtr_1"},
			Name:      sql.NullString{String: "John Doe", Valid: true},
			Email:     sql.NullString{String: "john@example.com", Valid: true},
		},
		{
			BaseModel: entity.BaseModel{ID: "invtr_2"},
			Name:      sql.NullString{String: "Jane Smith", Valid: true},
			Email:     sql.NullString{String: "jane@example.com", Valid: true},
		},
	}
	for _, investor := range investors {
		if err := db.Create(&investor).Error; err != nil {
			return fmt.Errorf("could not seed investors: %v", err)
		}
	}

	return nil
}
