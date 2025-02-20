package entity

import (
	"database/sql"
)

type Investment struct {
	BaseModel
	LoanID     string `gorm:"type:varchar(255);not null"`
	InvestorID string `gorm:"type:varchar(255);not null"`

	Amount sql.NullInt16   `gorm:"type:int"`
	Profit sql.NullInt16   `gorm:"type:int"`
	ROI    sql.NullFloat64 `gorm:"type:float"`
}
