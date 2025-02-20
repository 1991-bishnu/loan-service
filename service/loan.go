package service

import (
	"context"

	"github.com/1991-bishnu/loan-service/model"
	"github.com/1991-bishnu/loan-service/store"
)

type Loan interface {
	Create(ctx context.Context, req *model.CreateLoanReq) (res *model.CreateLoanRes, err error)
	Retrieve(ctx context.Context, id string) (res *model.RetrieveLoanRes, err error)
}
type loan struct {
	userStore *store.User
}

func NewLoan(userStore *store.User) Loan {
	return &loan{
		userStore: userStore,
	}
}

func (obj *loan) Create(ctx context.Context, req *model.CreateLoanReq) (res *model.CreateLoanRes, err error) {

	return res, nil
}

func (obj *loan) Retrieve(ctx context.Context, id string) (res *model.RetrieveLoanRes, err error) {
	// user, err := obj.userStore.GetByID(id)
	// if err != nil {
	// 	return userResponse, fmt.Errorf("Error in store function, error: %w", err)
	// }

	// userResponse.name = user.name

	return res, nil

}
