package user

import "github.com/labstack/echo/v4"

type UserController interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetUser(c echo.Context) error
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}
