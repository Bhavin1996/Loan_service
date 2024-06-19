package handlers

import (
	"loan-service/internal/models"
	"loan-service/internal/services"
	"loan-service/internal/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

var loans = make(map[string]*models.Loan)

// Create a new loan
func CreateLoan(c *gin.Context) {
	var loan models.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validation.ValidateLoanCreation(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan.ID = services.GenerateID()
	loan.State = models.Proposed
	loans[loan.ID] = &loan
	c.JSON(http.StatusCreated, loan)
}

// Approve a loan
func ApproveLoan(c *gin.Context) {
	id := c.Param("id")
	loan, exists := loans[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
		return
	}

	if loan.State != models.Proposed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "loan is not in proposed state"})
		return
	}

	var approvalInfo models.ApprovalInfo
	if err := c.ShouldBindJSON(&approvalInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan.ApprovalInfo = &approvalInfo
	loan.State = models.Approved
	c.JSON(http.StatusOK, loan)
}

// Invest in a loan
func InvestLoan(c *gin.Context) {
	id := c.Param("id")
	loan, exists := loans[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
		return
	}

	if loan.State != models.Approved {
		c.JSON(http.StatusBadRequest, gin.H{"error": "loan is not in approved state"})
		return
	}

	var investment models.Investment
	if err := c.ShouldBindJSON(&investment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalInvested := 0.0
	for _, inv := range loan.Investments {
		totalInvested += inv.Amount
	}
	totalInvested += investment.Amount

	if totalInvested > loan.PrincipalAmount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "total invested amount cannot exceed the loan principal"})
		return
	}

	loan.Investments = append(loan.Investments, investment)

	if totalInvested == loan.PrincipalAmount {
		loan.State = models.Invested
		services.SendInvestorEmails(loan)
	}

	c.JSON(http.StatusOK, loan)
}

// Disburse a loan
func DisburseLoan(c *gin.Context) {
	id := c.Param("id")
	loan, exists := loans[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
		return
	}

	if loan.State != models.Invested {
		c.JSON(http.StatusBadRequest, gin.H{"error": "loan is not in invested state"})
		return
	}

	var disbursementInfo models.DisbursementInfo
	if err := c.ShouldBindJSON(&disbursementInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan.DisbursementInfo = &disbursementInfo
	loan.State = models.Disbursed
	c.JSON(http.StatusOK, loan)
}
