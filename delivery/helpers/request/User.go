package request

type InsertUser struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	HP       string `json:"hp" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     int    `json:"role" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"`
	HP       string `json:"hp"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	HP       string `json:"hp"`
	Password string `json:"password" validate:"required"`
}
