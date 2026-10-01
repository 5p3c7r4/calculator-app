package models

// CalculatorRequest represents the request structure for calculator operations
type CalculatorRequest struct {
	Operation string   `json:"operation" binding:"required,oneof=add subtract multiply divide power sqrt percentage"`
	A         *float64  `json:"a"`
	B         *float64 `json:"b,omitempty"`
}
