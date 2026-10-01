package handler

import (
	"errors"
	"fmt"
	"log/slog"
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

	slog.Info("Received calculator request", "operation", req.Operation)

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Invalid request body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %s", err.Error())})
		return
	}

	if req.A == nil {
		slog.Error("Parameter A is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter A is required"})
		return
	}

	// Validate that B is provided for operations that need it
	needsB := req.Operation != "sqrt" && req.Operation != "perc"
	if needsB && req.B == nil {
		slog.Error("Parameter B is required for this operation", "operation", req.Operation)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter B is required for this operation"})
		return
	}

	// Perform calculation
	slog.Info("Performing calculation", "operation", req.Operation, "a", *req.A, "b", req.B)
	result, err := h.service.PerformOperation(req)
	if err != nil {
		slog.Error("Calculation failed", "error", err.Error(), "operation", req.Operation)
		if errors.Is(err, services.ErrDivisionByZero) || errors.Is(err, services.ErrSqrtOfNegative) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid operation"})
		}
		return
	}

	slog.Info("Calculation successful", "result", result.Result)
	c.JSON(http.StatusOK, result)
}
