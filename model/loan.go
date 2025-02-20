package model

type CreateLoanReq struct {
	UserID          string `json:"user_id"`
	PrincipalAmount int64  `json:"principal_amount"`
}

type CreateLoanRes struct {
	LoanID string `json:"loan_id"`
}

type RetrieveLoanReq struct {
	UserID string
	LoanID string
}

type RetrieveLoanRes struct {
	LoanID          string  `json:"loan_id"`
	Status          string  `json:"status"`
	PrincipalAmount int64   `json:"principal_amount"`
	TotalInterest   int64   `json:"total_interest"`
	ROI             float64 `json:"roi"`
	DisbursedAt     string  `json:"disbursed_at"`
	ApprovedAt      string  `json:"approved_at"`
}
