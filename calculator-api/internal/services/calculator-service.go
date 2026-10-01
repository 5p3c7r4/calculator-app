package services

import (
	"calculator-api/internal/models"
	"errors"
	"log/slog"
	"math"
)

var (
	ErrDivisionByZero = errors.New("you can not divide by zero")
	ErrSqrtOfNegative = errors.New("you can not get square root of a negative number")
	ErrUnsuportedOperation = errors.New("unsupported operation")
)

// CalculatorService handles calculator operations
type CalculatorService struct{}

// NewCalculatorService creates a new calculator service
func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

// CalculatorResponse represents the response structure for calculator operations
type CalculatorResponse struct {
	Result float64 `json:"result"`
}

// PerformOperation executes the requested calculator operation
func (s *CalculatorService) PerformOperation(req models.CalculatorRequest) (CalculatorResponse, error) {
	slog.Info("Starting calculation", "operation", req.Operation, "a", *req.A, "b", req.B)
	
	switch req.Operation {
	case "add":
		result := *req.A + *req.B
		slog.Info("Addition operation completed", "result", result)
		return CalculatorResponse{Result: result}, nil
	case "subtract":
		result := *req.A - *req.B
		slog.Info("Subtraction operation completed", "result", result)
		return CalculatorResponse{Result: result}, nil
	case "multiply":
		result := *req.A * *req.B
		slog.Info("Multiplication operation completed", "result", result)
		return CalculatorResponse{Result: result}, nil
	case "divide":
		if *req.B == 0 {
			slog.Error("Division by zero attempted", "a", *req.A, "b", *req.B)
			return CalculatorResponse{}, ErrDivisionByZero
		}
		result := *req.A / *req.B
		slog.Info("Division operation completed", "result", result)
		return CalculatorResponse{Result: result}, nil
	case "power":
		result := math.Pow(*req.A, *req.B)
		slog.Info("Power operation completed", "result", result)
		return CalculatorResponse{Result: result}, nil
	case "sqrt":
		if *req.A < 0 {
			slog.Error("Square root of negative number attempted", "a", *req.A)
			return CalculatorResponse{}, ErrSqrtOfNegative
		}
		result := math.Sqrt(*req.A)
		slog.Info("Square root operation completed", "result", result)
		return CalculatorResponse{Result: result}, nil
	case "percentage":
		result := *req.A * (*req.B / 100)
		slog.Info("Percentage operation completed", "result", result)
		return CalculatorResponse{Result: result}, nil
	default:
		slog.Error("Unsupported operation attempted", "operation", req.Operation)
		return CalculatorResponse{}, ErrUnsuportedOperation
	}
}