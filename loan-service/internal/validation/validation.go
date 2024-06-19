package validation

import (
	"errors"
	"loan-service/internal/models"
)

func ValidateLoanCreation(loan *models.Loan) error {
	if loan.BorrowerID == "" {
		return errors.New("borrower_id is required")
	}
	if loan.PrincipalAmount <= 0 {
		return errors.New("principal_amount must be greater than zero")
	}
	if loan.Rate <= 0 {
		return errors.New("rate must be greater than zero")
	}
	if loan.ROI <= 0 {
		return errors.New("roi must be greater than zero")
	}
	return nil
}
