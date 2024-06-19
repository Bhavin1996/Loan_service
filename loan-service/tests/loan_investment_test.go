package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"loan-service/internal/models"
	"loan-service/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoanInvestment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routes.Setup(router)

	// Mock initial loan data for testing
	initialLoan := &models.Loan{
		ID:              "1",
		BorrowerID:      "123",
		PrincipalAmount: 1000.0,
		Rate:            5.5,
		ROI:             6.5,
		State:           models.Approved, // Set loan state to Approved for investment tests
	}
	loans := map[string]*models.Loan{"1": initialLoan}

	t.Run("InvestLoan_Success", func(t *testing.T) {
		// Prepare investment payload
		investment := models.Investment{
			InvestorID: "investor1@example.com",
			Amount:     500.0,
		}
		body, _ := json.Marshal(investment)

		// Create a new HTTP request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/loans/1/invest", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		// Assert HTTP status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Assert loan state after investment
		var updatedLoan models.Loan
		json.Unmarshal(w.Body.Bytes(), &updatedLoan)
		assert.Contains(t, updatedLoan.State, models.Invested) // Ensure loan state contains "invested"
		assert.Len(t, updatedLoan.Investments, 1)              // Ensure there is one investment
		assert.Equal(t, investment.InvestorID, updatedLoan.Investments[0].InvestorID)
		assert.Equal(t, investment.Amount, updatedLoan.Investments[0].Amount)
	})

	t.Run("InvestLoan_InvalidState", func(t *testing.T) {
		// Set loan state to Disbursed
		loans["1"].State = models.Disbursed

		// Prepare investment payload
		investment := models.Investment{
			InvestorID: "investor2@example.com",
			Amount:     700.0,
		}
		body, _ := json.Marshal(investment)

		// Create a new HTTP request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/loans/1/invest", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		// Assert HTTP status code
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
