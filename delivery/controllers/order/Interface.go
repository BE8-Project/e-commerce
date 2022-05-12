package order

import "github.com/labstack/echo/v4"

type OrderController interface {
	Insert() echo.HandlerFunc
	GetStatus() echo.HandlerFunc
	Cancel() echo.HandlerFunc
}