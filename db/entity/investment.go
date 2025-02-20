package entity

import (
	"database/sql"
	"time"
)

type Investment struct {
	ID         string
	LoanID     string
	InvestorID string
	Status     sql.NullString

	Amount sql.NullInt16
	Profit sql.NullInt16
	ROI    sql.NullFloat64

	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
