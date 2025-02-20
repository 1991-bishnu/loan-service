package service

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/1991-bishnu/loan-service/constant"
	"github.com/1991-bishnu/loan-service/custom_error"
	"github.com/1991-bishnu/loan-service/db/entity"
	"github.com/1991-bishnu/loan-service/model"
	storeMock "github.com/1991-bishnu/loan-service/store/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInvest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoanStore := storeMock.NewMockLoan(ctrl)
	mockUserStore := storeMock.NewMockUser(ctrl)
	mockEmployeeStore := storeMock.NewMockEmployee(ctrl)
	mockInvestorStore := storeMock.NewMockInvestor(ctrl)
	mockInvestmentStore := storeMock.NewMockInvestment(ctrl)
	mockDocumentStore := storeMock.NewMockDocument(ctrl)

	loanService := NewLoan(mockLoanStore, mockUserStore, mockEmployeeStore, mockInvestorStore, mockInvestmentStore, mockDocumentStore)

	tests := []struct {
		name       string
		setupMocks func()
		ctx        context.Context
		req        *model.InvestReq
		wantRes    *model.InvestRes
		wantErr    bool
	}{
		{
			name: "Successful investment",
			setupMocks: func() {
				mockLoanStore.EXPECT().GetByID(gomock.Any(), "loan123").Return(&entity.Loan{
					BaseModel: entity.BaseModel{
						ID:     "loan123",
						Status: sql.NullString{String: constant.LoanStatusApproved, Valid: true},
					},
					PrincipalAmount: sql.NullInt64{Int64: 1000, Valid: true},
				}, nil)
				mockInvestorStore.EXPECT().GetByID(gomock.Any(), "investor123").Return(&entity.Investor{
					BaseModel: entity.BaseModel{
						ID: "investor123",
					},
				}, nil)
				mockInvestmentStore.EXPECT().GetByLoanID(gomock.Any(), gomock.Any()).Return([]*entity.Investment{}, nil)
				mockInvestmentStore.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(nil)
				mockLoanStore.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
				mockDocumentStore.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(nil)
			},
			ctx: context.Background(),
			req: &model.InvestReq{
				LoanID:       "loan123",
				InvestorID:   "investor123",
				InvestAmount: 1000,
			},
			wantRes: &model.InvestRes{
				InvestmentID: "investment123",
			},
			wantErr: false,
		},
		{
			name: "Loan not found",
			setupMocks: func() {
				mockLoanStore.EXPECT().GetByID(gomock.Any(), "loan123").Return(nil, custom_error.LoanNotFound)
			},
			ctx: context.Background(),
			req: &model.InvestReq{
				LoanID:       "loan123",
				InvestorID:   "investor123",
				InvestAmount: 1000,
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "Investor not found",
			setupMocks: func() {
				mockLoanStore.EXPECT().GetByID(gomock.Any(), "loan123").Return(&entity.Loan{
					BaseModel: entity.BaseModel{
						ID:     "loan123",
						Status: sql.NullString{String: constant.LoanStatusApproved, Valid: true},
					},
				}, nil)
				mockInvestorStore.EXPECT().GetByID(gomock.Any(), "investor123").Return(nil, custom_error.InvestorNotFound)
			},
			ctx: context.Background(),
			req: &model.InvestReq{
				LoanID:       "loan123",
				InvestorID:   "investor123",
				InvestAmount: 1000,
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "Invalid transition",
			setupMocks: func() {
				mockLoanStore.EXPECT().GetByID(gomock.Any(), "loan123").Return(&entity.Loan{
					BaseModel: entity.BaseModel{
						ID:     "loan123",
						Status: sql.NullString{String: constant.LoanStatusProposed, Valid: true},
					},
				}, nil)
				mockInvestorStore.EXPECT().GetByID(gomock.Any(), "investor123").Return(&entity.Investor{
					BaseModel: entity.BaseModel{
						ID: "investor123",
					},
				}, nil)
			},
			ctx: context.Background(),
			req: &model.InvestReq{
				LoanID:       "loan123",
				InvestorID:   "investor123",
				InvestAmount: 1000,
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "Document insert failure",
			setupMocks: func() {
				mockLoanStore.EXPECT().GetByID(gomock.Any(), "loan123").Return(&entity.Loan{
					BaseModel: entity.BaseModel{
						ID:     "loan123",
						Status: sql.NullString{String: constant.LoanStatusApproved, Valid: true},
					},
					PrincipalAmount: sql.NullInt64{Int64: 1000, Valid: true},
				}, nil)
				mockInvestorStore.EXPECT().GetByID(gomock.Any(), "investor123").Return(&entity.Investor{BaseModel: entity.BaseModel{ID: "investor123"}}, nil)
				mockInvestmentStore.EXPECT().GetByLoanID(gomock.Any(), gomock.Any()).Return([]*entity.Investment{}, nil)
				mockDocumentStore.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(fmt.Errorf("insert failed"))
			},
			ctx: context.Background(),
			req: &model.InvestReq{
				LoanID:       "loan123",
				InvestorID:   "investor123",
				InvestAmount: 1000,
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "Investment  insert failure",
			setupMocks: func() {
				mockLoanStore.EXPECT().GetByID(gomock.Any(), "loan123").Return(&entity.Loan{
					BaseModel: entity.BaseModel{
						ID:     "loan123",
						Status: sql.NullString{String: constant.LoanStatusApproved, Valid: true},
					},
					PrincipalAmount: sql.NullInt64{Int64: 1000, Valid: true},
				}, nil)
				mockInvestorStore.EXPECT().GetByID(gomock.Any(), "investor123").Return(&entity.Investor{BaseModel: entity.BaseModel{ID: "investor123"}}, nil)
				mockInvestmentStore.EXPECT().GetByLoanID(gomock.Any(), gomock.Any()).Return([]*entity.Investment{}, nil)
				mockDocumentStore.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(nil)
				mockInvestmentStore.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(fmt.Errorf("insert failed"))
			},
			ctx: context.Background(),
			req: &model.InvestReq{
				LoanID:       "loan123",
				InvestorID:   "investor123",
				InvestAmount: 1000,
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "Loan update failure",
			setupMocks: func() {
				mockLoanStore.EXPECT().GetByID(gomock.Any(), "loan123").Return(&entity.Loan{
					BaseModel: entity.BaseModel{
						ID:     "loan123",
						Status: sql.NullString{String: constant.LoanStatusApproved, Valid: true},
					},
					PrincipalAmount: sql.NullInt64{Int64: 1000, Valid: true},
				}, nil)
				mockInvestorStore.EXPECT().GetByID(gomock.Any(), "investor123").Return(&entity.Investor{BaseModel: entity.BaseModel{ID: "investor123"}}, nil)
				mockInvestmentStore.EXPECT().GetByLoanID(gomock.Any(), gomock.Any()).Return([]*entity.Investment{}, nil)
				mockDocumentStore.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(nil)
				mockInvestmentStore.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(nil)
				mockLoanStore.EXPECT().Update(gomock.Any(), gomock.Any()).Return(fmt.Errorf("update failed"))

			},
			ctx: context.Background(),
			req: &model.InvestReq{
				LoanID:       "loan123",
				InvestorID:   "investor123",
				InvestAmount: 1000,
			},
			wantRes: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			gotRes, err := loanService.Invest(tt.ctx, tt.req)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotEmpty(t, gotRes.InvestmentID)
			}
		})
	}
}

func TestCreateLoan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoanStore := storeMock.NewMockLoan(ctrl)
	mockUserStore := storeMock.NewMockUser(ctrl)
	mockEmployeeStore := storeMock.NewMockEmployee(ctrl)
	mockDocumentStore := storeMock.NewMockDocument(ctrl)
	mockInvestmentStore := storeMock.NewMockInvestment(ctrl)
	mockInvestorStore := storeMock.NewMockInvestor(ctrl)

	loanService := NewLoan(mockLoanStore, mockUserStore, mockEmployeeStore, mockInvestorStore, mockInvestmentStore, mockDocumentStore)

	ctx := context.Background()
	req := &model.CreateLoanReq{
		UserID:          "user123",
		PrincipalAmount: 10000,
	}

	mockUserStore.EXPECT().GetByID(ctx, req.UserID).Return(&entity.User{BaseModel: entity.BaseModel{ID: "user123"}}, nil)
	mockLoanStore.EXPECT().Insert(ctx, gomock.Any()).Return(nil)

	res, err := loanService.Create(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.LoanID)
}

func TestRetrieveLoan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoanStore := storeMock.NewMockLoan(ctrl)
	mockUserStore := storeMock.NewMockUser(ctrl)
	mockEmployeeStore := storeMock.NewMockEmployee(ctrl)
	mockDocumentStore := storeMock.NewMockDocument(ctrl)
	mockInvestmentStore := storeMock.NewMockInvestment(ctrl)
	mockInvestorStore := storeMock.NewMockInvestor(ctrl)

	loanService := NewLoan(mockLoanStore, mockUserStore, mockEmployeeStore, mockInvestorStore, mockInvestmentStore, mockDocumentStore)

	ctx := context.Background()
	req := &model.RetrieveLoanReq{
		LoanID: "loan123",
		UserID: "user123",
	}

	mockLoanStore.EXPECT().GetByIDAndUserID(ctx, req.LoanID, req.UserID).Return(&entity.Loan{
		BaseModel: entity.BaseModel{ID: "loan123"},
		UserID:    "user123",
	}, nil)
	mockDocumentStore.EXPECT().GetByLoanIDAndType(ctx, "loan123", constant.DocumentTypeAgreementBorrower).Return(&entity.Document{
		BaseModel: entity.BaseModel{ID: "user123"},
		URL:       sql.NullString{String: "http://example.com/agreement.pdf", Valid: true},
	}, nil)

	res, err := loanService.Retrieve(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "loan123", res.LoanID)
	assert.Equal(t, "http://example.com/agreement.pdf", res.AgreementURL)
}

func TestApproveLoan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoanStore := storeMock.NewMockLoan(ctrl)
	mockUserStore := storeMock.NewMockUser(ctrl)
	mockEmployeeStore := storeMock.NewMockEmployee(ctrl)
	mockDocumentStore := storeMock.NewMockDocument(ctrl)
	mockInvestmentStore := storeMock.NewMockInvestment(ctrl)
	mockInvestorStore := storeMock.NewMockInvestor(ctrl)

	loanService := NewLoan(mockLoanStore, mockUserStore, mockEmployeeStore, mockInvestorStore, mockInvestmentStore, mockDocumentStore)

	ctx := context.Background()
	req := &model.ApproveLoanReq{
		LoanID:     "loan123",
		EmployeeID: "employee123",
		ImageURL:   "http://example.com/image.jpg",
	}

	mockLoanStore.EXPECT().GetByID(ctx, req.LoanID).Return(&entity.Loan{
		BaseModel: entity.BaseModel{
			ID:     "loan123",
			Status: sql.NullString{String: constant.LoanStatusProposed, Valid: true},
		},
	}, nil)
	mockEmployeeStore.EXPECT().GetByID(ctx, req.EmployeeID).Return(&entity.Employee{
		BaseModel: entity.BaseModel{
			ID: "employee123",
		},
	}, nil)
	mockDocumentStore.EXPECT().Insert(ctx, gomock.Any()).Return(nil)
	mockLoanStore.EXPECT().Update(ctx, gomock.Any()).Return(nil)

	err := loanService.Approve(ctx, req)
	assert.NoError(t, err)
}
