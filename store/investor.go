package store

import (
	"context"
	"fmt"

	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/gorm"
)

type Investor interface {
	GetByID(ctx context.Context, id string) (employee *entity.Investor, err error)
}

type investor struct {
	db *gorm.DB
}

func NewInvestor(db *gorm.DB) Investor {
	return &investor{db: db}
}

func (obj *investor) GetByID(ctx context.Context, id string) (investor *entity.Investor, err error) {
	investor = &entity.Investor{}

	whereClouse := &entity.Investor{
		BaseModel: entity.BaseModel{
			ID: id,
		},
	}

	err = obj.db.WithContext(ctx).Scopes(ScopeNotDeleted()).Find(investor, whereClouse).Error
	if err != nil {
		return nil, fmt.Errorf("investor not found. Error: %w", err)
	}

	return investor, nil
}
