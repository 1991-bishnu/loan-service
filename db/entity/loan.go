package entity

import (
	"database/sql"
)

type Loan struct {
	BaseModel
	UserID string `gorm:"type:varchar(255);not null"`

	PrincipalAmount sql.NullInt64   `gorm:"type:int"`
	TotalInterest   sql.NullInt64   `gorm:"type:int"`
	ROI             sql.NullFloat64 `gorm:"type:float"`
	DisbursedBy     sql.NullString  `gorm:"type:varchar(255)"` // EmployeeID

	ApprovedAt  sql.NullTime `gorm:"type:timestamp"`
	DisbursedAt sql.NullTime `gorm:"type:timestamp"`
}
