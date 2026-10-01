// Backend tests for calculator service
// File: calculator-api/internal/services/calculator-service_test.go

package services

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"calculator-api/internal/models"
)

func TestCalculatorServicePerformOperation(t *testing.T) {
	service := NewCalculatorService()

	tests := []struct {
		name          string
		request       models.CalculatorRequest
		expectedResult float64
		expectError   bool
		errorType     error
	}{
		// Addition tests
		{
			name: "add positive numbers",
			request: models.CalculatorRequest{
				Operation: "add",
				A:         new(float64(5)),
				B:         new(float64(3)),
			},
			expectedResult: 8,
			expectError:    false,
		},
		{
			name: "add negative numbers",
			request: models.CalculatorRequest{
				Operation: "add",
				A:         new(float64(-5)),
				B:         new(float64(-3)),
			},
			expectedResult: -8,
			expectError:    false,
		},
		{
			name: "add zero",
			request: models.CalculatorRequest{
				Operation: "add",
				A:         new(float64(5)),
				B:        new(float64(0)),
			},
			expectedResult: 5,
			expectError:    false,
		},

		// Subtraction tests
		{
			name: "subtract positive numbers",
			request: models.CalculatorRequest{
				Operation: "subtract",
				A:         new(float64(10)),
				B:         new(float64(3)),
			},
			expectedResult: 7,
			expectError:    false,
		},
		{
			name: "subtract with negative result",
			request: models.CalculatorRequest{
				Operation: "subtract",
				A:         new(float64(3)),
				B:         new(float64(10)),
			},
			expectedResult: -7,
			expectError:    false,
		},

		// Multiplication tests
		{
			name: "multiply positive numbers",
			request: models.CalculatorRequest{
				Operation: "multiply",
				A:         new(float64(4)),
				B:         new(float64(5)),
			},
			expectedResult: 20,
			expectError:    false,
		},
		{
			name: "multiply by zero",
			request: models.CalculatorRequest{
				Operation: "multiply",
				A:         new(float64(10)),
				B:         new(float64(0)),
			},
			expectedResult: 0,
			expectError:    false,
		},

		// Division tests
		{
			name: "divide positive numbers",
			request: models.CalculatorRequest{
				Operation: "divide",
				A:         new(float64(10)),
				B:         new(float64(2)),
			},
			expectedResult: 5,
			expectError:    false,
		},
		{
			name: "divide by zero",
			request: models.CalculatorRequest{
				Operation: "divide",
				A:         new(float64(10)),
				B:         new(float64(0)),
			},
			expectedResult: 0,
			expectError:    true,
			errorType:      ErrDivisionByZero,
		},
		{
			name: "divide negative numbers",
			request: models.CalculatorRequest{
				Operation: "divide",
				A:         new(float64(-10)),
				B:         new(float64(2)),
			},
			expectedResult: -5,
			expectError:    false,
		},

		// Power tests
		{
			name: "power positive numbers",
			request: models.CalculatorRequest{
				Operation: "power",
				A:         new(float64(2)),
				B:         new(float64(3)),
			},
			expectedResult: 8,
			expectError:    false,
		},
		{
			name: "power zero exponent",
			request: models.CalculatorRequest{
				Operation: "power",
				A:         new(float64(5)),
				B:         new(float64(0)),
			},
			expectedResult: 1,
			expectError:    false,
		},

		// Square root tests
		{
			name: "square root positive number",
			request: models.CalculatorRequest{
				Operation: "sqrt",
				A:         new(float64(16)),
			},
			expectedResult: 4,
			expectError:    false,
		},
		{
			name: "square root zero",
			request: models.CalculatorRequest{
				Operation: "sqrt",
				A:         new(float64(0)),
			},
			expectedResult: 0,
			expectError:    false,
		},
		{
			name: "square root negative number",
			request: models.CalculatorRequest{
				Operation: "sqrt",
				A:         new(float64(-4)),
			},
			expectedResult: 0,
			expectError:    true,
			errorType:      ErrSqrtOfNegative,
		},

		// Percentage tests
		{
			name: "percentage calculation",
			request: models.CalculatorRequest{
				Operation: "percentage",
				A:         new(float64(50)),
				B:         new(float64(20)),
			},
			expectedResult: 10,
			expectError:    false,
		},
		{
			name: "percentage of zero",
			request: models.CalculatorRequest{
				Operation: "percentage",
				A:         new(float64(0)),
				B:         new(float64(50)),
			},
			expectedResult: 0,
			expectError:    false,
		},

		// Unsupported operation
		{
			name: "unsupported operation",
			request: models.CalculatorRequest{
				Operation: "invalid",
				A:         new(float64(5)),
				B:         new(float64(3)),
			},
			expectedResult: 0,
			expectError:    true,
			errorType:      ErrUnsuportedOperation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.PerformOperation(tt.request)
			
			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorType, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result.Result)
			}
		})
	}
}