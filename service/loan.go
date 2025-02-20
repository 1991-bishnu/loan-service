package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/1991-bishnu/loan-service/constant"
	"github.com/1991-bishnu/loan-service/db/entity"
	"github.com/1991-bishnu/loan-service/model"
	"github.com/1991-bishnu/loan-service/store"
	"github.com/1991-bishnu/loan-service/util"
)

type Loan interface {
	Create(ctx context.Context, req *model.CreateLoanReq) (res *model.CreateLoanRes, err error)
	Retrieve(ctx context.Context, req *model.RetrieveLoanReq) (res *model.RetrieveLoanRes, err error)
}
type loan struct {
	loanStore store.Loan
	userStore store.User
}

func NewLoan(loanStore store.Loan, userStore store.User) Loan {
	return &loan{
		loanStore: loanStore,
		userStore: userStore,
	}
}

func (obj *loan) Create(ctx context.Context, req *model.CreateLoanReq) (res *model.CreateLoanRes, err error) {
	// Validate user data from db
	user, err := obj.userStore.GetByID(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found, error: %w", err)
	}

	// Create a record in DB
	loan := &entity.Loan{
		BaseModel: entity.BaseModel{
			ID:     util.GeneratePID(constant.PrefixLoan),
			Status: sql.NullString{String: constant.LoanStatusProposed, Valid: true},
		},
		UserID:          user.ID,
		PrincipalAmount: sql.NullInt64{Int64: int64(req.PrincipalAmount), Valid: true},
	}

	if err := obj.loanStore.Insert(ctx, loan); err != nil {
		return nil, fmt.Errorf("loan insert failed, error: %w", err)
	}

	// Generate response data
	res = &model.CreateLoanRes{
		LoanID: loan.ID,
	}

	return res, nil
}

func (obj *loan) Retrieve(ctx context.Context, req *model.RetrieveLoanReq) (res *model.RetrieveLoanRes, err error) {
	// Initialize the response model
	res = &model.RetrieveLoanRes{}

	// Fetch the loan details based on the provided loan ID and user ID
	var loan *entity.Loan
	if req.UserID != "" {
		loan, err = obj.loanStore.GetByIDAndUserID(ctx, req.LoanID, req.UserID)
	} else {
		loan, err = obj.loanStore.GetByID(ctx, req.LoanID)
	}

	if err != nil {
		return nil, fmt.Errorf("loan retrieval failed, error: %w", err)
	}

	// Create the response model
	res.LoanID = loan.ID
	res.Status = loan.Status.String
	res.PrincipalAmount = loan.PrincipalAmount.Int64
	res.TotalInterest = loan.TotalInterest.Int64
	res.ROI = loan.ROI.Float64
	if loan.DisbursedAt.Valid {
		res.DisbursedAt = loan.DisbursedAt.Time.String()
	}
	if loan.ApprovedAt.Valid {
		res.ApprovedAt = loan.ApprovedAt.Time.String()
	}

	return res, nil
}
