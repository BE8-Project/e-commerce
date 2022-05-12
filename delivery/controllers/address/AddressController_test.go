package address

import (
	"e-commerce/delivery/helpers/response"
	middlewares "e-commerce/delivery/middleware"
	"e-commerce/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	token_merchant string
	token_user string
)

func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token_merchant, _ = middlewares.CreateToken(2, "merchant")
		token_user, _ = middlewares.CreateToken(3, "user")
	})
}

func TestInsert(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"address": "Jl. Kebon Jeruk No. 1",
			"city": "Bandung",
			"country": "Indonesia",
			"zip_code": 40257,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/address")
		controller := NewAddressController(&mockAddress{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "success create Address!", resp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}{"user_id":float64(2), "created_at": "0001-01-01T00:00:00Z"}), resp.Data)
	})

	t.Run("Status Invalid", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"address": "Jl. Kebon Jeruk No. 1",
			"city": "Bandung",
			"country": "Indonesia",
			"zip_code": "40257",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewAddressController(&mockAddress{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "invalid request", resp.Message)
	})

	t.Run("Status Validate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"address": "",
			"city": "Bandung",
			"country": "Indonesia",
			"zip_code": 40257,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/address")

		controller := NewAddressController(&mockAddress{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		if res.Body.Bytes() == nil {
			validator := validator.New()
			err := validator.Struct(context.ParamValues())
			if err != nil {
				assert.Equal(t, 400, res.Code)
				assert.Equal(t, "invalid request", res.Body.String())
				assert.Nil(t, err)
			}
		}
	})

	t.Run("Status BadRequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"address": "Jl. Kebon Jeruk No. 1",
			"city": "Bandung",
			"country": "Indonesia",
			"zip_code": 40257,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewAddressController(&mockError{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "assert.AnError general error for testing", resp.Message)
	})
}

func TestGetByUserID(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/address")

		controller := NewAddressController(&mockAddress{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetByUserID())(context)

		type Response struct {

			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success get Address!", resp.Message)
		assert.Equal(t, []interface {}([]interface {}{map[string]interface {}{"address":"Jl. Kebon Jeruk No. 1", "city":"Bandung", "country":"Indonesia", "id":float64(1), "zip_code":float64(40257)}}), resp.Data)
	})

	t.Run("Status Not Found", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/users/address", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewAddressController(&mockError{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetByUserID())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.Product
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Address not found!", resp.Message)
		assert.Equal(t, []response.Product([]response.Product(nil)), resp.Data)
	})


}

type mockAddress struct {}

func (ma *mockAddress) Insert(address *entity.Address) (response.InsertAddress, error) {
	return response.InsertAddress{
		UserID: address.UserID,
		CreatedAt: address.CreatedAt,
	}, nil
}

func (ma *mockAddress) GetByUserID(userID uint) []response.Address {
	return []response.Address{
		{
			ID: 1,
			Address: "Jl. Kebon Jeruk No. 1",
			City: "Bandung",
			Country: "Indonesia",
			ZipCode: 40257,
		},
	}
}

type mockError struct {}

func (ma *mockError) Insert(address *entity.Address) (response.InsertAddress, error) {
	return response.InsertAddress{}, assert.AnError
}

func (ma *mockError) GetByUserID(userID uint) []response.Address {
	return []response.Address{}
}
