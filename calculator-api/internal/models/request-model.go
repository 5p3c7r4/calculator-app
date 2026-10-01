package models

// CalculatorRequest represents the request structure for calculator operations
type CalculatorRequest struct {
	Operation string   `json:"operation" binding:"required,oneof=add subtract multiply divide exp sqrt perc"`
	A         float64  `json:"a" binding:"required"`
	B         *float64 `json:"b,omitempty"`
}
