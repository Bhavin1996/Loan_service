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

func TestLoanApproval(t *testing.T) {
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
		State:           models.Proposed,
	}
	loans := map[string]*models.Loan{"1": initialLoan}

	t.Run("ApproveLoan_Success", func(t *testing.T) {
		// Set loan state to Proposed
		loans["1"].State = models.Proposed

		// Prepare approval information payload
		approvalInfo := models.ApprovalInfo{
			PictureProof: "http://example.com/picture_proof",
			EmployeeID:   "emp001",
			ApprovalDate: time.Now(),
		}
		loans["1"].ApprovalInfo = &approvalInfo
		body, _ := json.Marshal(approvalInfo)

		// Create a new HTTP request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/loans/1/approve", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		// Assert HTTP status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Assert loan state after approval
		var updatedLoan models.Loan
		json.Unmarshal(w.Body.Bytes(), &updatedLoan)
		assert.Equal(t, models.Approved, updatedLoan.State) // Ensure the loan state is updated to "approved"
		assert.NotNil(t, updatedLoan.ApprovalInfo)          // Ensure approval info is set
		assert.Equal(t, approvalInfo.PictureProof, updatedLoan.ApprovalInfo.PictureProof)
		assert.Equal(t, approvalInfo.EmployeeID, updatedLoan.ApprovalInfo.EmployeeID)
		assert.Equal(t, approvalInfo.ApprovalDate.Format(time.RFC3339), updatedLoan.ApprovalInfo.ApprovalDate.Format(time.RFC3339))
	})

	t.Run("ApproveLoan_NotFound", func(t *testing.T) {
		// Create a new HTTP request for a non-existent loan ID
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/loans/999/approve", nil)
		router.ServeHTTP(w, req)

		// Assert HTTP status code
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("ApproveLoan_InvalidState", func(t *testing.T) {
		// Set loan state to Approved
		loans["1"].State = models.Approved

		// Prepare approval information payload
		approvalInfo := models.ApprovalInfo{
			PictureProof: "http://example.com/picture_proof",
			EmployeeID:   "emp001",
			ApprovalDate: time.Now(),
		}
		body, _ := json.Marshal(approvalInfo)

		// Create a new HTTP request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/loans/1/approve", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		// Assert HTTP status code
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
