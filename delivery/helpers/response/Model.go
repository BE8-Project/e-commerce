package response

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	HP        string    `json:"hp"`
	CreatedAt time.Time `json:"created_at"`
}

type Login struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type DeleteUser struct {
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UpdateUser struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InsertCategory struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Category struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type InsertProduct struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateProduct struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteProduct struct {
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

type ProductMerchant struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock uint    `json:"stock"`
	Image string  `json:"image"`
	// Status uint `json:"status"`
}
type CartInsert struct {
	ID        uint    `json:"id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
	UserID    uint    `json:"userID"`
	Image     string  `json:"image"`
	ProductID uint    `json:"produkId"`
}
