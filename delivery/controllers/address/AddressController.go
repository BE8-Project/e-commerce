package address

import (
	"e-commerce/delivery/helpers/request"
	"e-commerce/delivery/helpers/response"
	middlewares "e-commerce/delivery/middleware"
	"e-commerce/entity"
	repoAddress "e-commerce/repository/address"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type addressController struct {
	Connect repoAddress.AddressModel
	Validate *validator.Validate
}

func NewAddressController(conn repoAddress.AddressModel) *addressController {
	return &addressController{
		Connect: conn,
		Validate: validator.New(),
	}
}

func (ac *addressController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)
		var request request.InsertAddress

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := ac.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		address := entity.Address{
			UserID: uint(userID),
			Address: request.Address,
			City: request.City,
			Country: request.Country,
			ZipCode: request.ZipCode,
		}

		result, err := ac.Connect.Insert(&address)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success create Address!", result))
	}
}

func (ac *addressController) GetByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)

		addresses := ac.Connect.GetByUserID(uint(userID))
		
		if len(addresses) == 0 {
			return c.JSON(http.StatusNotFound, response.StatusNotFound("Address not found!"))
		}
		
		return c.JSON(http.StatusOK, response.StatusOK("success get Address!", addresses))
	}
}