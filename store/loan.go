package store

import (
	"context"

	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/gorm"
)

type Loan interface {
	Insert(ctx context.Context, loan *entity.Loan) (err error)
	GetByIDAndUserID(ctx context.Context, id string, userID string) (loan *entity.Loan, err error)
	GetByID(ctx context.Context, id string) (loan *entity.Loan, err error)
}

type loan struct {
	db *gorm.DB
}

func NewLoan(db *gorm.DB) Loan {
	return &loan{db: db}
}

func (obj *loan) Insert(ctx context.Context, loan *entity.Loan) (err error) {
	err = obj.db.Create(loan).Error
	return
}

func (obj *loan) GetByIDAndUserID(ctx context.Context, id string, userID string) (loan *entity.Loan, err error) {
	loan = &entity.Loan{}
	err = obj.db.WithContext(ctx).Scopes(ScopeNotDeleted(), ScopeUserID(userID)).First(loan, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return loan, nil
}

func (obj *loan) GetByID(ctx context.Context, id string) (loan *entity.Loan, err error) {
	loan = &entity.Loan{}
	err = obj.db.WithContext(ctx).Scopes(ScopeNotDeleted()).First(loan, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return loan, nil
}
