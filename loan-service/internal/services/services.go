package services

import (
	"crypto/rand"
	"encoding/hex"
	"loan-service/internal/mailer"
	"loan-service/internal/models"
	"log"
)

func GenerateID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func SendInvestorEmails(loan *models.Loan) {
	for _, investment := range loan.Investments {
		toEmail := investment.InvestorID // Assuming InvestorID is the email address
		subject := "Loan Investment Agreement"
		body := "Dear Investor,\n\nThank you for your investment. Please find the agreement letter here: " + loan.AgreementLetter + "\n\nBest regards,\nLoan Service Team"
		err := mailer.SendEmail(toEmail, subject, body)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", toEmail, err)
		} else {
			log.Printf("Email sent to %s successfully", toEmail)
		}
	}
}
