package model

import "time"

type CreateLoanReq struct {
	UserID          string `json:"user_id"`
	PrincipalAmount int64  `json:"principal_amount"`
}

type CreateLoanRes struct {
	LoanID string `json:"loan_id"`
}

type RetrieveLoanRes struct {
	LoanID          string    `json:"loan_id"`
	Status          string    `json:"status"`
	PrincipalAmount int64     `json:"principal_amount"`
	TotalInterest   int64     `json:"total_interest"`
	ROI             float64   `json:"roi"`
	DisbursedDate   time.Time `json:"disbursed_date"`
}
