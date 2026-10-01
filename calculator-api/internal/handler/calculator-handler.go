package handler

import (
	"errors"
	"fmt"
	"net/http"

	"calculator-api/internal/models"
	"calculator-api/internal/services"

	"github.com/gin-gonic/gin"
)

// CalculatorHandler handles calculator HTTP requests
type CalculatorHandler struct {
	service *services.CalculatorService
}

// NewCalculatorHandler creates a new calculator handler
func NewCalculatorHandler(service *services.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{service: service}
}

// Calculate handles calculator operations
func (h *CalculatorHandler) Calculate(c *gin.Context) {
	var req models.CalculatorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %s", err.Error())})
		return
	}

	if req.A == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter A is required"})
		return
	}

	// Validate that B is provided for operations that need it
	needsB := req.Operation != "sqrt" && req.Operation != "perc"
	if needsB && req.B == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter B is required for this operation"})
		return
	}

	// Perform calculation
	result, err := h.service.PerformOperation(req)
	if err != nil {
		if errors.Is(err, services.ErrDivisionByZero) || errors.Is(err, services.ErrSqrtOfNegative) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid operation"})
		}
		return
	}

	c.JSON(http.StatusOK, result)
}
