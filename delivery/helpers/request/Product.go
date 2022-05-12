package request

type InsertProduct struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Stock       uint    `json:"stock" validate:"required"`
	Image       string  `json:"image" validate:"required"`
	CategoryID  uint    `json:"category_id" validate:"required"`
}

type UpdateProduct struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description"`
	Stock       uint    `json:"stock" validate:"required"`
	Image       string  `json:"image"`
	CategoryID  uint    `json:"category_id"`
}