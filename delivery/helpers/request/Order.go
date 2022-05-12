package request

type InsertOrder struct {
	AddressID   uint    `json:"address_id" validate:"required"`
	PaymentType string  `json:"payment_type" validate:"required"`
	Total       float64 `json:"total" validate:"required"`
}