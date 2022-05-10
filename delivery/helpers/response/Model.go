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
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Category struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}