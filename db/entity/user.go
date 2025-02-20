package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID     string
	Status sql.NullString

	Name  sql.NullString
	Email sql.NullString

	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
