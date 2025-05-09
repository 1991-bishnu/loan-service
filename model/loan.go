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
	PrincipalAmount int64   `json:"principal_amount"`
	TotalInterest   int64   `json:"total_interest"`
	ROI             float64 `json:"roi"`
	DisbursedAt     string  `json:"disbursed_at"`
	AgreementURL    string  `json:"agreement_url"`
}

type Document struct {
	DocumentID string `json:"document_id"`
	Type       string `json:"type"`
	URL        string `json:"url"`
	EmployeeID string `json:"employee_id"`
	CreatedAt  string `json:"created_at"`
}

type ApproveLoanReq struct {
	LoanID     string
	EmployeeID string `json:"employee_id"`
	ImageURL   string `json:"image_url"`
}

type InvestReq struct {
	LoanID       string
	InvestorID   string  `json:"investor_id"`
	InvestAmount int64   `json:"invest_amount"`
	ROI          float64 `json:"roi"`
}

type InvestRes struct {
	InvestmentID string  `json:"investment_id"`
	LoanID       string  `json:"loan_id"`
	InvestAmount int64   `json:"invest_amount"`
	ROI          float64 `json:"roi"`
	Profit       int64   `json:"profit"`
}

type DisbursReq struct {
	LoanID       string
	EmployeeID   string `json:"employee_id"`
	AgreementURL string `json:"agreement_url"`
}
