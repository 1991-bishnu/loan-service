package entity

import (
	"database/sql"

	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model

	ID     string         `gorm:"primaryKey;type:varchar(255)"`
	Status sql.NullString `gorm:"type:varchar(255)"`
}
