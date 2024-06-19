package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"loan-service/internal/models"
	"loan-service/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoanDisbursement(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routes.Setup(router)

	t.Run("DisburseLoan_Success", func(t *testing.T) {
		initialLoan := &models.Loan{
			ID:              "1",
			BorrowerID:      "123",
			PrincipalAmount: 1000.0,
			Rate:            5.5,
			ROI:             6.5,
			State:           models.Invested,
		}
		loans := map[string]*models.Loan{"1": initialLoan}

		disbursementInfo := models.DisbursementInfo{
			SignedAgreementLetter: "http://example.com/agreement_letter",
			EmployeeID:            "emp001",
			DisbursementDate:      time.Now(),
		}
		body, _ := json.Marshal(disbursementInfo)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/loans/1/disburse", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var updatedLoan models.Loan
		err := json.Unmarshal(w.Body.Bytes(), &updatedLoan)
		assert.NoError(t, err)
		assert.Equal(t, models.Disbursed, updatedLoan.State)
		assert.Nil(t, updatedLoan.DisbursementInfo)
		assert.Equal(t, disbursementInfo.SignedAgreementLetter, updatedLoan.DisbursementInfo.SignedAgreementLetter)
		assert.Equal(t, disbursementInfo.EmployeeID, updatedLoan.DisbursementInfo.EmployeeID)
		assert.WithinDuration(t, disbursementInfo.DisbursementDate, updatedLoan.DisbursementInfo.DisbursementDate, time.Second)

		// Update loans map with the disbursed loan state
		loans[updatedLoan.ID] = &updatedLoan
	})

	t.Run("DisburseLoan_InvalidState", func(t *testing.T) {
		initialLoan := &models.Loan{
			ID:              "1",
			BorrowerID:      "123",
			PrincipalAmount: 1000.0,
			Rate:            5.5,
			ROI:             6.5,
			State:           models.Approved, // Set loan state back to approved for invalid disbursement attempt
		}
		loans := map[string]*models.Loan{"1": initialLoan}

		disbursementInfo := models.DisbursementInfo{
			SignedAgreementLetter: "http://example.com/agreement_letter",
			EmployeeID:            "emp002",
			DisbursementDate:      time.Now(),
		}
		body, _ := json.Marshal(disbursementInfo)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/loans/1/disburse", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		// Ensure loan state remains unchanged in the map
		assert.Equal(t, initialLoan.State, loans["1"].State)
	})
}
