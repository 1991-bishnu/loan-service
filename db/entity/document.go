package entity

import (
	"database/sql"
	"time"
)

type Document struct {
	ID     string
	LoanID string

	Status sql.NullString

	SubmitedBy sql.NullString // EmployeeID
	Type       sql.NullString
	URL        sql.NullString

	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
