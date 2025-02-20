package entity

import (
	"database/sql"
)

type Document struct {
	BaseModel

	LoanID     string         `gorm:"type:varchar(255);not null"`
	SubmitedBy sql.NullString `gorm:"type:varchar(255)"`
	Type       sql.NullString `gorm:"type:varchar(255)"`
	URL        sql.NullString `gorm:"type:varchar(255)"`
}
