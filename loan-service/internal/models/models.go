package models

import "time"

// Loan states
const (
	Proposed  = "proposed"
	Approved  = "approved"
	Invested  = "invested"
	Disbursed = "disbursed"
)

type Loan struct {
	ID               string            `json:"id"`
	BorrowerID       string            `json:"borrower_id" binding:"required"`
	PrincipalAmount  float64           `json:"principal_amount" binding:"required"`
	Rate             float64           `json:"rate" binding:"required"`
	ROI              float64           `json:"roi" binding:"required"`
	AgreementLetter  string            `json:"agreement_letter"`
	State            string            `json:"state"`
	ApprovalInfo     *ApprovalInfo     `json:"approval_info,omitempty"`
	Investments      []Investment      `json:"investments,omitempty"`
	DisbursementInfo *DisbursementInfo `json:"disbursement_info,omitempty"`
}

type ApprovalInfo struct {
	PictureProof string    `json:"picture_proof" binding:"required"`
	EmployeeID   string    `json:"employee_id" binding:"required"`
	ApprovalDate time.Time `json:"approval_date" binding:"required"`
}

type Investment struct {
	InvestorID string  `json:"investor_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

type DisbursementInfo struct {
	SignedAgreementLetter string    `json:"signed_agreement_letter" binding:"required"`
	EmployeeID            string    `json:"employee_id" binding:"required"`
	DisbursementDate      time.Time `json:"disbursement_date" binding:"required"`
}
