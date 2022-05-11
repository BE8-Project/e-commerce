package order

import (
	"e-commerce/delivery/helpers/request"
	"e-commerce/delivery/helpers/response"
	middlewares "e-commerce/delivery/middleware"
	"e-commerce/entity"
	repoOrder "e-commerce/repository/order"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type orderController struct {
	Connect repoOrder.OrderModel
}

func NewProductController(conn repoOrder.OrderModel) *orderController {
	return &orderController{
		Connect: conn,
	}
}

func Random() string {
	time.Sleep(500 * time.Millisecond)
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func (u *orderController) Insert() echo.HandlerFunc {
	return func (c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)
		var request request.InsertOrder

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		order := entity.Order{
			TrackingNumber: "DM-"+Random(),
			PaymentType: request.PaymentType,
			Total:       request.Total,
			AddressID:  uint(userID),
			Status:      "pending",
		}

		result, err := u.Connect.Insert(&order)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success create Order!", result))
	}
}