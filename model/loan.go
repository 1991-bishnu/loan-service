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
	LoanID          string     `json:"loan_id"`
	Status          string     `json:"status"`
	PrincipalAmount int64      `json:"principal_amount"`
	TotalInterest   int64      `json:"total_interest"`
	ROI             float64    `json:"roi"`
	DisbursedAt     string     `json:"disbursed_at"`
	ApprovedAt      string     `json:"approved_at"`
	CreatedAt       string     `json:"created_at"`
	Documents       []Document `json:"documents"`
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
