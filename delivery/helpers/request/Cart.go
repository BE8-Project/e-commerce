package request

type InsertCart struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
type UpdateCart struct {
	Quantity uint `json:"quantity"`
}
