package entity

import (
	"database/sql"
)

type Investment struct {
	BaseModel
	LoanID     string `gorm:"type:varchar(255);not null"`
	InvestorID string `gorm:"type:varchar(255);not null"`

	Amount sql.NullInt64   `gorm:"type:int"`
	Profit sql.NullInt64   `gorm:"type:int"`
	ROI    sql.NullFloat64 `gorm:"type:float"`
}
