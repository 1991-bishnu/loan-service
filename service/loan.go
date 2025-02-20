package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/1991-bishnu/loan-service/constant"
	"github.com/1991-bishnu/loan-service/custom_error"
	"github.com/1991-bishnu/loan-service/db/entity"
	"github.com/1991-bishnu/loan-service/model"
	"github.com/1991-bishnu/loan-service/store"
	"github.com/1991-bishnu/loan-service/util"
)

type Loan interface {
	Create(ctx context.Context, req *model.CreateLoanReq) (res *model.CreateLoanRes, err error)
	Retrieve(ctx context.Context, req *model.RetrieveLoanReq) (res *model.RetrieveLoanRes, err error)
	Approve(ctx context.Context, req *model.ApproveLoanReq) (err error)
}
type loan struct {
	loanStore     store.Loan
	userStore     store.User
	employeeStore store.Employee
	documentStore store.Document
}

func NewLoan(
	loanStore store.Loan,
	userStore store.User,
	employeeStore store.Employee,
	documentStore store.Document) Loan {
	return &loan{
		loanStore:     loanStore,
		userStore:     userStore,
		employeeStore: employeeStore,
		documentStore: documentStore,
	}
}

func (obj *loan) Create(ctx context.Context, req *model.CreateLoanReq) (res *model.CreateLoanRes, err error) {
	user, err := obj.userStore.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found, error: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found, loan_id: %s error: %w", req.UserID, custom_error.UserNotFound)
	}

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

	ctxWithTO, cancel := context.WithTimeout(ctx, 30*time.Second)
	go func(ctx context.Context, cancel context.CancelFunc, loan *entity.Loan) {
		defer cancel()
		// TODO: Call field officer system to notify regarding new loan application
	}(ctxWithTO, cancel, loan)

	res = &model.CreateLoanRes{
		LoanID: loan.ID,
	}

	return res, nil
}

func (obj *loan) Retrieve(ctx context.Context, req *model.RetrieveLoanReq) (res *model.RetrieveLoanRes, err error) {
	res = &model.RetrieveLoanRes{}

	var loan *entity.Loan
	if req.UserID != "" {
		loan, err = obj.loanStore.GetByIDAndUserID(ctx, req.LoanID, req.UserID)
	} else {
		loan, err = obj.loanStore.GetByID(ctx, req.LoanID)
	}

	if err != nil {
		return nil, fmt.Errorf("loan retrieval failed, error: %w", err)
	}
	if loan == nil {
		return nil, fmt.Errorf("loan not found, loan_id: %s error: %w", req.LoanID, custom_error.LoanNotFound)
	}

	res.LoanID = loan.ID
	res.Status = loan.Status.String
	res.PrincipalAmount = loan.PrincipalAmount.Int64
	res.TotalInterest = loan.TotalInterest.Int64
	res.ROI = loan.ROI.Float64
	res.CreatedAt = loan.CreatedAt.String()
	if loan.DisbursedAt.Valid {
		res.DisbursedAt = loan.DisbursedAt.Time.String()
	}
	if loan.ApprovedAt.Valid {
		res.ApprovedAt = loan.ApprovedAt.Time.String()
	}

	documents, err := obj.documentStore.GetByLoanID(ctx, loan.ID)
	if err != nil {
		return nil, fmt.Errorf("document retrieval failed, error: %w", err)
	}

	for _, document := range documents {
		doc := model.Document{
			EmployeeID: document.SubmitedBy.String,
			Type:       document.Type.String,
			DocumentID: document.ID,
			URL:        document.URL.String,
			CreatedAt:  document.CreatedAt.String(),
		}
		res.Documents = append(res.Documents, doc)
	}

	return res, nil
}

func (obj *loan) Approve(ctx context.Context, req *model.ApproveLoanReq) (err error) {
	// TODO: Can be ran as concurrent
	loan, err := obj.loanStore.GetByID(ctx, req.LoanID)
	if err != nil {
		return fmt.Errorf("loan not found, error: %w", err)
	}
	if loan == nil {
		return fmt.Errorf("loan not found. loan_id: %s error: %w", req.LoanID, custom_error.LoanNotFound)

	}

	// TODO: Can be ran as concurrent
	employee, err := obj.employeeStore.GetByID(ctx, req.EmployeeID)
	if err != nil {
		return fmt.Errorf("employee not found, error: %w", err)
	}
	if employee == nil {
		return fmt.Errorf("employee not found. employee_id: %s error: %w", req.EmployeeID, custom_error.EmployeeNotFound)

	}

	if loan.Status.String != constant.LoanStatusProposed {
		return fmt.Errorf("loan_id: %s error: %w", req.LoanID, custom_error.ErrInvalidTransition)
	}

	// TODO: Optimize this
	stateMachine := util.NewStateMachine(loan.Status.String)
	nextStatus, err := stateMachine.GetNextStage()
	if err != nil {
		return fmt.Errorf("stateMachine.GetNextStage, error: %w", err)
	}

	// Insert document
	document := &entity.Document{
		BaseModel: entity.BaseModel{
			ID: util.GeneratePID(constant.PrefixDocument),
		},
		LoanID:     loan.ID,
		SubmitedBy: sql.NullString{String: req.EmployeeID, Valid: true},
		Type:       sql.NullString{String: constant.DocumentTypePictureProof, Valid: true},
		URL:        sql.NullString{String: req.ImageURL, Valid: true},
	}

	if err := obj.documentStore.Insert(ctx, document); err != nil {
		return fmt.Errorf("document insert failed, error: %w", err)
	}

	loan.Status = sql.NullString{String: nextStatus, Valid: true}
	loan.ApprovedAt = sql.NullTime{Time: time.Now(), Valid: true}
	if err := obj.loanStore.Update(ctx, loan); err != nil {
		return fmt.Errorf("loan update failed, error: %w", err)
	}

	ctxWithTO, cancel := context.WithTimeout(ctx, 30*time.Second)
	go func(ctx context.Context, cancel context.CancelFunc, loan *entity.Loan) {
		defer cancel()
		// TODO: Call lender system to notify regarding new loan application and ready for investment
	}(ctxWithTO, cancel, loan)

	return nil
}
