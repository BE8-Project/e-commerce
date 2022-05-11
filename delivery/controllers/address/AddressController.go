package address

import (
	"e-commerce/delivery/helpers/request"
	"e-commerce/delivery/helpers/response"
	middlewares "e-commerce/delivery/middleware"
	"e-commerce/entity"
	repoAddress "e-commerce/repository/address"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type addressController struct {
	Connect repoAddress.AddressModel
}

func NewProductController(conn repoAddress.AddressModel) *addressController {
	return &addressController{
		Connect: conn,
	}
}

func (ac *addressController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)
		var request request.InsertAddress

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		address := entity.Address{
			UserID: uint(userID),
			Address: request.Address,
			City: request.City,
		}

		fmt.Println(address)

		return c.JSON(http.StatusCreated, response.StatusCreated("success create address!", address))
	}
}