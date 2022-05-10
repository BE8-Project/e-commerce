package request

type InsertUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	HP       string `json:"hp"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	HP       string `json:"hp"`
	Password string `json:"password"`
}
