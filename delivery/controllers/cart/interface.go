package cart

import "github.com/labstack/echo/v4"

type CartController interface {
	Insert() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
