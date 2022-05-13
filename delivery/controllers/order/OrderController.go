package order

import (
	"e-commerce/config"
	"e-commerce/delivery/helpers/request"
	"e-commerce/delivery/helpers/response"
	middlewares "e-commerce/delivery/middleware"
	"e-commerce/entity"
	repoOrder "e-commerce/repository/order"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type orderController struct {
	Connect repoOrder.OrderModel
	Validate *validator.Validate
}

func NewOrderController(conn repoOrder.OrderModel) *orderController {
	return &orderController{
		Connect: conn,
		Validate: validator.New(),
	}
}

func Random() string {
	time.Sleep(500 * time.Millisecond)
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func Gopay(order_id string, total float64) *coreapi.ChargeReqWithMap {
	req := &coreapi.ChargeReqWithMap{
		"payment_type": "gopay",
		"transaction_details": map[string]interface{}{
			"order_id":     order_id,
			"gross_amount": total,
		},
	}
	
	return req
}

func (u *orderController) Insert() echo.HandlerFunc {
	return func (c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)
		var request request.InsertOrder

		if !u.Connect.CheckRole(uint(userID)) {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := u.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		order := entity.Order{
			AddressID:   request.AddressID,
			PaymentType: request.PaymentType,
			Total:       request.Total,
			TrackingNumber: "DM-"+Random(),
			Status:      "pending",
			UserID: 	uint(userID),
		}

		result, err := u.Connect.Insert(&order)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		config.SetupGlobalMidtransConfigApi()
		midtrans.SetPaymentAppendNotification("https://midtrans-java.herokuapp.com/notif/append1")
		midtrans.SetPaymentOverrideNotification("https://midtrans-java.herokuapp.com/notif/override")

		resp, _ := coreapi.ChargeTransactionWithMap(Gopay(result.OrderID, result.Total))

		var message []interface{}
		var transaction_id string
		for key, value := range resp {
			if key == "actions" {
				message = value.([]interface{})
			}

			if key == "transaction_id" {
				transaction_id = value.(string)
			}
		}

		var action map[string]interface{}
		for key, value := range message {
			if key == 1 {
				action = value.(map[string]interface{})
			}
		}

		var data map[string]interface{} = make(map[string]interface{})
		data["order_id"] = result.OrderID
		data["payment_type"] = "gopay"
		data["total"] = result.Total
		data["status"] = "pending"
		data["payment_simulator"] = action["url"]
		data["payment_url"] = "https://api.sandbox.midtrans.com/v2/gopay/"+transaction_id+"/qr-code"
		data["created_at"] = result.CreatedAt

		return c.JSON(http.StatusCreated, response.StatusCreated("success create Order!", data))
	}
}

func (u *orderController) GetStatus() echo.HandlerFunc {
	return func (c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)
		order_id := c.Param("order_id")

		if !u.Connect.CheckRole(uint(userID)) {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}

		config.SetupGlobalMidtransConfigApi()
		midtrans.SetPaymentAppendNotification("https://midtrans-java.herokuapp.com/notif/append1")
		midtrans.SetPaymentOverrideNotification("https://midtrans-java.herokuapp.com/notif/override")

		resp, _ := coreapi.CheckTransaction(order_id)

		update := entity.Order{
			TrackingNumber: resp.OrderID,
			Status: resp.TransactionStatus,
		}

		result, err := u.Connect.Update(order_id, &update)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusOK, response.StatusOK("success get Status!", result))
	}
}

func (u *orderController) Cancel() echo.HandlerFunc {
	return func (c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)
		order_id := c.Param("order_id")

		if !u.Connect.CheckRole(uint(userID)) {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}

		config.SetupGlobalMidtransConfigApi()
		midtrans.SetPaymentAppendNotification("https://midtrans-java.herokuapp.com/notif/append1")
		midtrans.SetPaymentOverrideNotification("https://midtrans-java.herokuapp.com/notif/override")

		resp, _ := coreapi.CancelTransaction(order_id)

		update := entity.Order{
			TrackingNumber: resp.OrderID,
			Status: resp.TransactionStatus,
		}

		result, err := u.Connect.Update(order_id, &update)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusOK, response.StatusOK("success cancel Order!", result))
	}
}