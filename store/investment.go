package store

import (
	"context"
	"fmt"

	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/gorm"
)

type Investment interface {
	Insert(ctx context.Context, investment *entity.Investment) (err error)
	GetByLoanID(ctx context.Context, loadID string) ([]*entity.Investment, error)
}

type investment struct {
	db *gorm.DB
}

func NewInvestment(db *gorm.DB) Investment {
	return &investment{db: db}
}

func (obj *investment) Insert(ctx context.Context, investment *entity.Investment) (err error) {
	err = obj.db.WithContext(ctx).Create(investment).Error
	if err != nil {
		return fmt.Errorf("investment create failed. Error: %w", err)
	}
	return nil
}

func (obj *investment) GetByLoanID(ctx context.Context, id string) ([]*entity.Investment, error) {
	var investments []*entity.Investment

	whereClouse := &entity.Investment{
		BaseModel: entity.BaseModel{
			ID: id,
		},
	}

	err := obj.db.WithContext(ctx).Scopes(ScopeNotDeleted()).Find(investments, whereClouse).Error
	if err != nil {
		return nil, fmt.Errorf("investment not found. Error: %w", err)
	}

	return investments, nil
}
