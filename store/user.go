package store

import (
	"context"
	"fmt"

	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=user.go -destination=mock/user.go -package=store
type User interface {
	GetByID(ctx context.Context, id string) (user *entity.User, err error)
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return &user{db: db}
}

func (obj *user) GetByID(ctx context.Context, id string) (user *entity.User, err error) {
	user = &entity.User{}

	whereClouse := &entity.User{
		BaseModel: entity.BaseModel{
			ID: id,
		},
	}

	err = obj.db.WithContext(ctx).Scopes(ScopeNotDeleted()).Find(user, whereClouse).Error
	if err != nil {
		return nil, fmt.Errorf("user not found. Error: %w", err)
	}

	return user, nil
}
