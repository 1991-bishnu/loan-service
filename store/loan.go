package store

import (
	"context"
	"fmt"

	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/gorm"
)

type Loan interface {
	Insert(ctx context.Context, loan *entity.Loan) (err error)
	GetByIDAndUserID(ctx context.Context, id string, userID string) (loan *entity.Loan, err error)
	GetByID(ctx context.Context, id string) (loan *entity.Loan, err error)
	Update(ctx context.Context, loan *entity.Loan) (err error)
}

type loan struct {
	db *gorm.DB
}

func NewLoan(db *gorm.DB) Loan {
	return &loan{db: db}
}

func (obj *loan) Insert(ctx context.Context, loan *entity.Loan) (err error) {
	err = obj.db.WithContext(ctx).Create(loan).Error
	if err != nil {
		return fmt.Errorf("loan create failed. Error: %w", err)
	}
	return nil
}

func (obj *loan) GetByIDAndUserID(ctx context.Context, id string, userID string) (loan *entity.Loan, err error) {
	loan = &entity.Loan{}

	whereClouse := &entity.Loan{
		BaseModel: entity.BaseModel{
			ID: id,
		},
	}

	err = obj.db.WithContext(ctx).Scopes(ScopeNotDeleted(), ScopeUserID(userID)).Find(loan, whereClouse).Error
	if err != nil {
		return nil, fmt.Errorf("loan not found. Error: %w", err)
	}
	return loan, nil
}

func (obj *loan) GetByID(ctx context.Context, id string) (loan *entity.Loan, err error) {
	loan = &entity.Loan{}

	whereClouse := &entity.Loan{
		BaseModel: entity.BaseModel{
			ID: id,
		},
	}

	err = obj.db.WithContext(ctx).Scopes(ScopeNotDeleted()).Find(loan, whereClouse).Error
	if err != nil {
		return nil, fmt.Errorf("loan not found. Error: %w", err)
	}
	return loan, nil
}

func (obj *loan) Update(ctx context.Context, loan *entity.Loan) (err error) {
	err = obj.db.WithContext(ctx).Save(loan).Error
	if err != nil {
		return fmt.Errorf("loan update failed. Error: %w", err)
	}
	return nil
}
