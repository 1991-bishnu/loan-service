package store

import (
	"context"
	"fmt"

	"github.com/1991-bishnu/loan-service/db/entity"
	"gorm.io/gorm"
)

type Employee interface {
	GetByID(ctx context.Context, id string) (employee *entity.Employee, err error)
}

type employee struct {
	db *gorm.DB
}

func NewEmployee(db *gorm.DB) Employee {
	return &employee{db: db}
}

func (obj *employee) GetByID(ctx context.Context, id string) (employee *entity.Employee, err error) {
	employee = &entity.Employee{}

	whereClouse := &entity.Employee{
		BaseModel: entity.BaseModel{
			ID: id,
		},
	}

	err = obj.db.WithContext(ctx).Scopes(ScopeNotDeleted()).Find(employee, whereClouse).Error
	if err != nil {
		return nil, fmt.Errorf("employee not found. Error: %w", err)
	}

	return employee, nil
}
