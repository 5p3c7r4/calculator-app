package services

import (
	"calculator-api/internal/models"
	"errors"
	"math"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrSqrtOfNegative = errors.New("square root of negative number")
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
	switch req.Operation {
	case "add":
		return CalculatorResponse{Result: *req.A + *req.B}, nil
	case "subtract":
		return CalculatorResponse{Result: *req.A - *req.B}, nil
	case "multiply":
		return CalculatorResponse{Result: *req.A * *req.B}, nil
	case "divide":
		if *req.B == 0 {
			return CalculatorResponse{}, ErrDivisionByZero
		}
		return CalculatorResponse{Result: *req.A / *req.B}, nil
	case "power":
		return CalculatorResponse{Result: math.Pow(*req.A, *req.B)}, nil
	case "sqrt":
		if *req.A < 0 {
			return CalculatorResponse{}, ErrSqrtOfNegative
		}
		return CalculatorResponse{Result: math.Sqrt(*req.A)}, nil
	case "percentage":
		return CalculatorResponse{Result: *req.A * (*req.B / 100)}, nil
	default:
		return CalculatorResponse{}, ErrUnsuportedOperation
	}
}