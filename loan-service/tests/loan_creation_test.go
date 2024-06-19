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

func TestLoanCreation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routes.Setup(router)

	t.Run("CreateLoan_Success", func(t *testing.T) {
		loan := map[string]interface{}{
			"borrower_id":      "123",
			"principal_amount": 1000.0,
			"rate":             5.5,
			"roi":              6.5,
		}

		body, _ := json.Marshal(loan)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/loans", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var createdLoan models.Loan
		json.Unmarshal(w.Body.Bytes(), &createdLoan)
		assert.Equal(t, "123", createdLoan.BorrowerID)
		assert.Equal(t, 1000.0, createdLoan.PrincipalAmount)
		assert.Equal(t, 5.5, createdLoan.Rate)
		assert.Equal(t, 6.5, createdLoan.ROI)
		assert.Equal(t, models.Proposed, createdLoan.State)
	})

	t.Run("CreateLoan_InvalidData", func(t *testing.T) {
		invalidLoan := map[string]interface{}{
			"borrower_id": "123",
			// Missing principal_amount, rate, roi fields
		}

		body, _ := json.Marshal(invalidLoan)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/loans", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
