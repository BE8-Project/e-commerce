package product

import "github.com/labstack/echo/v4"

type ProductController interface {
	Insert() echo.HandlerFunc
	GetAll(c echo.Context) error
	GetBySlug(c echo.Context) error
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByCategory(c echo.Context) error
	GetBySearch(c echo.Context) error
	GetAllMerchant() echo.HandlerFunc
}