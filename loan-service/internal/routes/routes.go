package routes

import (
	"loan-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.POST("/api/v1/loans", handlers.CreateLoan)
	router.POST("/api/v1/loans/:id/approve", handlers.ApproveLoan)
	router.POST("/api/v1/loans/:id/invest", handlers.InvestLoan)
	router.POST("/api/v1/loans/:id/disburse", handlers.DisburseLoan)
}
