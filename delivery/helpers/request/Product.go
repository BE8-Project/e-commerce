package request

type InsertProduct struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       uint    `json:"stock"`
	Image       string  `json:"image"`
	CategoryID  uint    `json:"category_id"`
}