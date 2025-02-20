package store

import (
	"context"
	"fmt"

	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/gorm"
)

type Document interface {
	Insert(ctx context.Context, loan *entity.Document) (err error)
	GetByLoanID(ctx context.Context, loanID string) (documents []*entity.Document, err error)
}

type document struct {
	db *gorm.DB
}

func NewDocument(db *gorm.DB) Document {
	return &document{db: db}
}

func (obj *document) Insert(ctx context.Context, document *entity.Document) (err error) {
	err = obj.db.WithContext(ctx).Create(document).Error
	if err != nil {
		return fmt.Errorf("document create failed. Error: %w", err)
	}
	return nil
}

func (obj *document) GetByLoanID(ctx context.Context, loanID string) ([]*entity.Document, error) {
	var documents []*entity.Document

	whereClouse := &entity.Document{
		LoanID: loanID,
	}
	err := obj.db.WithContext(ctx).Scopes(ScopeNotDeleted()).Find(&documents, whereClouse).Error
	if err != nil {
		return nil, fmt.Errorf("document not found. Error: %w", err)
	}

	return documents, nil
}
