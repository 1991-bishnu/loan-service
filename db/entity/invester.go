package entity

import (
	"database/sql"
)

type Investor struct {
	BaseModel

	Name  sql.NullString `gorm:"type:varchar(255)"`
	Email sql.NullString `gorm:"type:varchar(255)"`
}
