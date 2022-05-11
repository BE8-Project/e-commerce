package request

type InsertOrder struct {
	PaymentType string  `json:"payment_type"`
	Total       float64 `json:"total"`
}