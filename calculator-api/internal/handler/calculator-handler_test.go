// Backend tests for calculator handler
// File: calculator-api/internal/handler/calculator-handler_test.go

package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"calculator-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCalculatorHandlerCalculate(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create calculator service
	service := services.NewCalculatorService()
	handler := NewCalculatorHandler(service)

	// Create router and call handler
	router := gin.New()
	router.POST("/api/v1/calculator", handler.Calculate)

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
		expectedResult float64
		expectError    bool
		errorMessage   string
	}{
		// Addition tests
		{
			name: "add positive numbers",
			requestBody: map[string]interface{}{
				"operation": "add",
				"a":         5.0,
				"b":         3.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 8.0,
			expectError:    false,
		},
		{
			name: "add negative numbers",
			requestBody: map[string]interface{}{
				"operation": "add",
				"a":         -5.0,
				"b":         -3.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: -8.0,
			expectError:    false,
		},

		// Subtraction tests
		{
			name: "subtract positive numbers",
			requestBody: map[string]interface{}{
				"operation": "subtract",
				"a":         10.0,
				"b":         3.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 7.0,
			expectError:    false,
		},

		// Multiplication tests
		{
			name: "multiply positive numbers",
			requestBody: map[string]interface{}{
				"operation": "multiply",
				"a":         4.0,
				"b":         5.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 20.0,
			expectError:    false,
		},
		{
			name: "multiply by zero",
			requestBody: map[string]interface{}{
				"operation": "multiply",
				"a":         10.0,
				"b":         0.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 0.0,
			expectError:    false,
		},

		// Division tests
		{
			name: "divide positive numbers",
			requestBody: map[string]interface{}{
				"operation": "divide",
				"a":         10.0,
				"b":         2.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 5.0,
			expectError:    false,
		},
		{
			name: "divide by zero",
			requestBody: map[string]interface{}{
				"operation": "divide",
				"a":         10.0,
				"b":         0.0,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
			errorMessage:   "you can not divide by zero",
		},
		{
			name: "divide negative numbers",
			requestBody: map[string]interface{}{
				"operation": "divide",
				"a":         -10.0,
				"b":         2.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: -5.0,
			expectError:    false,
		},

		// Power tests
		{
			name: "power positive numbers",
			requestBody: map[string]interface{}{
				"operation": "power",
				"a":         2.0,
				"b":         3.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 8.0,
			expectError:    false,
		},
		{
			name: "power zero exponent",
			requestBody: map[string]interface{}{
				"operation": "power",
				"a":         5.0,
				"b":         0.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 1.0,
			expectError:    false,
		},

		// Square root tests
		{
			name: "square root positive number",
			requestBody: map[string]interface{}{
				"operation": "sqrt",
				"a":         16.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 4.0,
			expectError:    false,
		},
		{
			name: "square root zero",
			requestBody: map[string]interface{}{
				"operation": "sqrt",
				"a":         0.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 0.0,
			expectError:    false,
		},
		{
			name: "square root negative number",
			requestBody: map[string]interface{}{
				"operation": "sqrt",
				"a":         -4.0,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
			errorMessage:   "you can not get square root of a negative number",
		},

		// Percentage tests
		{
			name: "percentage calculation",
			requestBody: map[string]interface{}{
				"operation": "percentage",
				"a":         50.0,
				"b":         20.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 10.0,
			expectError:    false,
		},
		{
			name: "percentage of zero",
			requestBody: map[string]interface{}{
				"operation": "percentage",
				"a":         0.0,
				"b":         50.0,
			},
			expectedStatus: http.StatusOK,
			expectedResult: 0.0,
			expectError:    false,
		},

		// Error cases
		{
			name: "invalid operation",
			requestBody: map[string]interface{}{
				"operation": "invalid",
				"a":         5.0,
				"b":         3.0,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "missing required B parameter for operation that needs it",
			requestBody: map[string]interface{}{
				"operation": "add",
				"a":         5.0,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
			errorMessage:   "Parameter B is required for this operation",
		},
		{
			name: "invalid JSON request",
			requestBody: map[string]interface{}{
				"operation": "add",
				"a":         "invalid",
				"b":         3.0,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
			errorMessage:   "Invalid request body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert request body to JSON
			jsonBody, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err)

			// Create HTTP request
			req := httptest.NewRequest("POST", "/api/v1/calculator", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectError {
				// Check error response
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Contains(t, response["error"], tt.errorMessage)
			} else {
				// Check successful response
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, response["result"])
			}
		})
	}
}
