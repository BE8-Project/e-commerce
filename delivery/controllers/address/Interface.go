package address

import "github.com/labstack/echo/v4"

type AddressController interface {
	Insert() echo.HandlerFunc
	GetByUserID() echo.HandlerFunc
}