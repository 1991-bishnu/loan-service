package store

import (
	"fmt"

	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/gorm"
)

type User interface {
	GetByID(id string) (user *entity.User, err error)
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return &user{db: db}
}

func (obj *user) GetByID(id string) (user *entity.User, err error) {
	user = &entity.User{}
	if err := obj.db.Last(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found. Error: %w", err)
	}

	return user, nil
}
