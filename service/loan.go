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
	Invest(ctx context.Context, req *model.InvestReq) (res *model.InvestRes, err error)
	Disburse(ctx context.Context, req *model.DisbursReq) (err error)
}
type loan struct {
	loanStore       store.Loan
	userStore       store.User
	employeeStore   store.Employee
	investorStore   store.Investor
	investmentStore store.Investment
	documentStore   store.Document
}

func NewLoan(
	loanStore store.Loan,
	userStore store.User,
	employeeStore store.Employee,
	investorStore store.Investor,
	investmentStore store.Investment,
	documentStore store.Document) Loan {
	return &loan{
		loanStore:       loanStore,
		userStore:       userStore,
		employeeStore:   employeeStore,
		investorStore:   investorStore,
		investmentStore: investmentStore,
		documentStore:   documentStore,
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
	if loan.ID == "" {
		return nil, fmt.Errorf("loan not found, loan_id: %s error: %w", req.LoanID, custom_error.LoanNotFound)
	}

	res.LoanID = loan.ID
	res.PrincipalAmount = loan.PrincipalAmount.Int64
	res.TotalInterest = loan.TotalInterest.Int64
	res.ROI = loan.ROI.Float64
	if loan.DisbursedAt.Valid {
		res.DisbursedAt = loan.DisbursedAt.Time.String()
	}

	document, err := obj.documentStore.GetByLoanIDAndType(ctx, loan.ID, constant.DocumentTypeAgreementBorrower)
	if err != nil {
		return nil, fmt.Errorf("document retrieval failed, error: %w", err)
	}
	if document.ID != "" {
		res.AgreementURL = document.URL.String
	}

	return res, nil
}

func (obj *loan) Approve(ctx context.Context, req *model.ApproveLoanReq) (err error) {
	// TODO: Can be ran as concurrent
	loan, err := obj.loanStore.GetByID(ctx, req.LoanID)
	if err != nil {
		return fmt.Errorf("loan not found, error: %w", err)
	}
	if loan.ID == "" {
		return fmt.Errorf("loan not found. loan_id: %s error: %w", req.LoanID, custom_error.LoanNotFound)

	}

	// TODO: Can be ran as concurrent
	employee, err := obj.employeeStore.GetByID(ctx, req.EmployeeID)
	if err != nil {
		return fmt.Errorf("employee not found, error: %w", err)
	}
	if employee.ID == "" {
		return fmt.Errorf("employee not found. employee_id: %s error: %w", req.EmployeeID, custom_error.EmployeeNotFound)

	}

	if loan.Status.String != constant.LoanStatusProposed {
		return fmt.Errorf("loan_id: %s error: %w", req.LoanID, custom_error.ErrInvalidTransition)
	}

	nextStatus, err := util.GetNextStage(loan.Status.String)
	if err != nil {
		return fmt.Errorf("stateMachine.GetNextStage, error: %w", err)
	}

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

func (obj *loan) Invest(ctx context.Context, req *model.InvestReq) (res *model.InvestRes, err error) {
	res = &model.InvestRes{}
	// TODO: Can be ran as concurrent
	loan, err := obj.loanStore.GetByID(ctx, req.LoanID)
	if err != nil {
		return nil, fmt.Errorf("loan not found, error: %w", err)
	}
	if loan.ID == "" {
		return nil, fmt.Errorf("loan not found. loan_id: %s error: %w", req.LoanID, custom_error.LoanNotFound)

	}

	// TODO: Can be ran as concurrent
	investor, err := obj.investorStore.GetByID(ctx, req.InvestorID)
	if err != nil {
		return nil, fmt.Errorf("investor not found, error: %w", err)
	}
	if investor.ID == "" {
		return nil, fmt.Errorf("investor not found. investor_id: %s error: %w", req.InvestorID, custom_error.InvestorNotFound)

	}

	if loan.Status.String != constant.LoanStatusApproved {
		return nil, fmt.Errorf("loan_id: %s error: %w", req.LoanID, custom_error.ErrInvalidTransition)
	}

	investment := &entity.Investment{
		BaseModel: entity.BaseModel{
			ID: util.GeneratePID(constant.PrefixInvestment),
		},
		LoanID:     loan.ID,
		InvestorID: investor.ID,
		Amount:     sql.NullInt64{Int64: req.InvestAmount, Valid: true},
		ROI:        sql.NullFloat64{Float64: req.ROI, Valid: true},
	}

	// Get all investment and calculate with current invest amount should les/equal to principal amount
	investments, err := obj.investmentStore.GetByLoanID(ctx, req.LoanID)
	if err != nil {
		return nil, fmt.Errorf("investment not found, error: %w", err)
	}

	investments = append(investments, investment)
	var investedAmount int64
	var weightedROI float64

	for _, investment := range investments {
		investedAmount += investment.Amount.Int64
		weightedROI += (float64(investment.Amount.Int64) * investment.ROI.Float64)
	}

	if investedAmount > loan.PrincipalAmount.Int64 {
		return nil, fmt.Errorf("loan_id: %s error: %w", req.LoanID, custom_error.InvalidInvestAmount)
	}
	if investedAmount == loan.PrincipalAmount.Int64 {

		finalROI := weightedROI / float64(investedAmount)
		totalInterest := util.CalculateProfit(investedAmount, finalROI)
		nextStatus, err := util.GetNextStage(loan.Status.String)
		if err != nil {
			return nil, fmt.Errorf("stateMachine.GetNextStage, error: %w", err)
		}
		loan.Status = sql.NullString{String: nextStatus, Valid: true}
		loan.ROI = sql.NullFloat64{Float64: finalROI, Valid: true}
		loan.TotalInterest = sql.NullInt64{Int64: totalInterest, Valid: true}
	}

	investmentAgreementDocument := &entity.Document{
		BaseModel: entity.BaseModel{
			ID: util.GeneratePID(constant.PrefixDocument),
		},
		LoanID: loan.ID,
		Type:   sql.NullString{String: constant.DocumentTypeAgreementInvestor, Valid: true},
		URL:    sql.NullString{String: "dummy_url", Valid: true}, // TODO: Once lender sign the agreement, PDF upload to storage and capture storage link her
	}
	if err := obj.documentStore.Insert(ctx, investmentAgreementDocument); err != nil {
		return nil, fmt.Errorf("investment document insert failed, error: %w", err)
	}

	profit := util.CalculateProfit(req.InvestAmount, req.ROI)
	investment.Profit = sql.NullInt64{Int64: profit, Valid: true}
	investment.AgreementID = sql.NullString{String: investmentAgreementDocument.ID, Valid: true}
	if err := obj.investmentStore.Insert(ctx, investment); err != nil {
		return nil, fmt.Errorf("investment insert failed, error: %w", err)
	}

	// TODO: Apply DB transaction
	if err := obj.loanStore.Update(ctx, loan); err != nil {
		return nil, fmt.Errorf("loan update failed, error: %w", err)
	}

	if loan.Status.String == constant.LoanStatusInvested {
		ctxWithTO, cancel := context.WithTimeout(ctx, 30*time.Second)
		go func(ctx context.Context, cancel context.CancelFunc, loan *entity.Loan, investments []*entity.Investment) {
			defer cancel()
			// TODO: Send email to all investor with agreement url
			// Investment table has investor and agreement details
		}(ctxWithTO, cancel, loan, investments)
	}

	res.InvestmentID = investment.ID
	res.LoanID = loan.ID
	res.InvestAmount = investment.Amount.Int64
	res.ROI = investment.ROI.Float64
	res.Profit = investment.Profit.Int64

	return res, nil
}

func (obj *loan) Disburse(ctx context.Context, req *model.DisbursReq) (err error) {
	// TODO: Can be ran as concurrent
	loan, err := obj.loanStore.GetByID(ctx, req.LoanID)
	if err != nil {
		return fmt.Errorf("loan not found, error: %w", err)
	}
	if loan.ID == "" {
		return fmt.Errorf("loan not found. loan_id: %s error: %w", req.LoanID, custom_error.LoanNotFound)

	}

	// TODO: Can be ran as concurrent
	employee, err := obj.employeeStore.GetByID(ctx, req.EmployeeID)
	if err != nil {
		return fmt.Errorf("employee not found, error: %w", err)
	}
	if employee.ID == "" {
		return fmt.Errorf("employee not found. employee_id: %s error: %w", req.EmployeeID, custom_error.EmployeeNotFound)

	}

	if loan.Status.String != constant.LoanStatusInvested {
		return fmt.Errorf("loan_id: %s error: %w", req.LoanID, custom_error.ErrInvalidTransition)
	}

	nextStatus, err := util.GetNextStage(loan.Status.String)
	if err != nil {
		return fmt.Errorf("stateMachine.GetNextStage, error: %w", err)
	}

	document := &entity.Document{
		BaseModel: entity.BaseModel{
			ID: util.GeneratePID(constant.PrefixDocument),
		},
		LoanID:     loan.ID,
		SubmitedBy: sql.NullString{String: req.EmployeeID, Valid: true},
		Type:       sql.NullString{String: constant.DocumentTypeAgreementBorrower, Valid: true},
		URL:        sql.NullString{String: req.AgreementURL, Valid: true},
	}
	if err := obj.documentStore.Insert(ctx, document); err != nil {
		return fmt.Errorf("document insert failed, error: %w", err)
	}

	loan.Status = sql.NullString{String: nextStatus, Valid: true}
	loan.DisbursedAt = sql.NullTime{Time: time.Now(), Valid: true}
	if err := obj.loanStore.Update(ctx, loan); err != nil {
		return fmt.Errorf("loan update failed, error: %w", err)
	}

	return nil
}
