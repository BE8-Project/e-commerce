package request

type InsertCart struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  uint `json:"quantity" validate:"required"`
}
type UpdateCart struct {
	Quantity uint `json:"quantity" validate:"required"`
}
