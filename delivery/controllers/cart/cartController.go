package cart

import (
	"e-commerce/delivery/helpers/request"
	"e-commerce/delivery/helpers/response"
	token "e-commerce/delivery/middleware"
	repoCart "e-commerce/repository/cart"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type cartController struct {
	Connect  repoCart.CartModel
	Validate *validator.Validate
}

func NewCartController(conn repoCart.CartModel) *cartController {
	return &cartController{
		Connect:  conn,
		Validate: validator.New(),
	}
}

func (u *cartController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := token.ExtractTokenUserId(c)

		var request request.InsertCart

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := u.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		err := u.Connect.Insert(request, uint(userID))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success add to Cart!", "Succes"))
	}
}

func (u *cartController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := token.ExtractTokenUserId(c)

		cartAll, error := u.Connect.GetAll(uint(userID))
		if error != nil {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}

		if len(cartAll) == 0 {
			return c.JSON(http.StatusNotFound, response.StatusNotFound("Cart not found"))
		}

		return c.JSON(http.StatusOK, response.StatusOK("success get cart!", cartAll))
	}
}

func (u *cartController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := token.ExtractTokenUserId(c)
		cartId, _ := strconv.Atoi(c.Param("id"))
		CartID := uint(cartId)

		if res := u.Connect.Checkid(uint(userID), CartID); res != nil {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}

		var request request.UpdateCart

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := u.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		error := u.Connect.Update(request, CartID)
		if error != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(error))
		}

		cartAll, error := u.Connect.GetAll(uint(userID))
		if error != nil {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}

		if len(cartAll) == 0 {
			return c.JSON(http.StatusNotFound, response.StatusNotFound("Cart not found"))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success update to Cart!", cartAll))
	}
}

func (u *cartController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := token.ExtractTokenUserId(c)
		cartId, _ := strconv.Atoi(c.Param("id"))
		CartID := uint(cartId)

		if res := u.Connect.Checkid(uint(userID), CartID); res != nil {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}

		var request request.InsertCart

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		error := u.Connect.Delete(CartID)
		if error != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(error))
		}
		cartAll, error := u.Connect.GetAll(uint(userID))
		if error != nil {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}

		if len(cartAll) == 0 {
			return c.JSON(http.StatusNotFound, response.StatusNotFound("Cart not found"))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success delete to Cart!", cartAll))
	}
}
