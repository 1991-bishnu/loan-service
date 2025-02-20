package entity

import (
	"database/sql"
	"time"
)

type Loan struct {
	ID     string
	UserID string
	Status sql.NullString

	PrincipalAmount sql.NullInt16
	TotalInterest   sql.NullInt16
	ROI             sql.NullFloat64
	DisbursedBy     sql.NullString // EmployeeID

	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
