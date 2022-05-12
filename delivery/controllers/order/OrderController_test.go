package order

// import (
// 	"e-commerce/delivery/helpers/response"
// 	middlewares "e-commerce/delivery/middleware"
// 	"e-commerce/entity"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	token_merchant string
// 	token_user     string
// )

// func TestCreateToken(t *testing.T) {
// 	t.Run("Create Token", func(t *testing.T) {
// 		token_merchant, _ = middlewares.CreateToken(2, "merchant")
// 		token_user, _ = middlewares.CreateToken(3, "user")
// 	})
// }

// // Membuat request untuk insert product
// func TestInsert(t *testing.T) {
// 	t.Run("Status OK", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"name":        "Product 1",
// 			"price":       100000,
// 			"description": "Product 1",
// 			"image":       "Product 1",
// 			"stock":       1,
// 			"category_id": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/merchants/products")
// 		controller := NewProductController(&mockProduct{})

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

// 		type Response struct {
// 			Code    int    `json:"code"`
// 			Message string `json:"message"`
// 			Data    interface{}
// 		}

// 		var resp Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

// 		assert.Equal(t, 201, resp.Code)
// 		assert.Equal(t, "success create Product!", resp.Message)
// 		assert.Equal(t, map[string]interface{}(map[string]interface{}{"name": "Product 1", "created_at": "0001-01-01T00:00:00Z"}), resp.Data)
// 	})

// 	t.Run("Status Forbidden", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"name":        "Product 1",
// 			"price":       100000,
// 			"description": "Product 1",
// 			"image":       "Product 1",
// 			"stock":       1,
// 			"category_id": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/merchants/products")
// 		controller := NewProductController(&mockError{})

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

// 		type Response struct {
// 			Code    int    `json:"code"`
// 			Message string `json:"message"`
// 			Data    interface{}
// 		}

// 		var resp Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

// 		assert.Equal(t, 403, resp.Code)
// 		assert.Equal(t, "You are not allowed to access this resource", resp.Message)
// 	})

// 	t.Run("Status Invalid", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"name":        "Product 1",
// 			"price":       "100000",
// 			"description": "Product 1",
// 			"image":       "Product 1",
// 			"stock":       1,
// 			"category_id": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/merchants/products")
// 		controller := NewProductController(&mockProduct{})

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

// 		type Response struct {
// 			Code    int    `json:"code"`
// 			Message string `json:"message"`
// 			Data    interface{}
// 		}

// 		var resp Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

// 		assert.Equal(t, 400, resp.Code)
// 		assert.Equal(t, "invalid request", resp.Message)
// 	})

// 	t.Run("Status BadRequest", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"name":        "Product 1",
// 			"price":       100000,
// 			"description": "Product 1",
// 			"image":       "Product 1",
// 			"stock":       1,
// 			"category_id": 1,
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/merchants/products")
// 		controller := NewProductController(&mockErrorInput{})

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

// 		type Response struct {
// 			Code    int    `json:"code"`
// 			Message string `json:"message"`
// 			Data    interface{}
// 		}

// 		var resp Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

// 		assert.Equal(t, 400, resp.Code)
// 		assert.Equal(t, "assert.AnError general error for testing", resp.Message)
// 	})
// }

// type mockOrder struct {}

// func (mo *mockOrder) CheckRole(id uint) bool {

// }

// func (mo *mockOrder) Insert(order *entity.Order) (response.InsertOrder, error) {

// }

// func (mo *mockOrder) Update(order_id string, order *entity.Order) (response.UpdateOrder, error) {
// }

// type mockError struct {}

// func (mo *mockError) CheckRole(id uint) bool {

// }

// func (mo *mockError) Insert(order *entity.Order) (response.InsertOrder, error) {

// }

// func (mo *mockError) Update(order_id string, order *entity.Order) (response.UpdateOrder, error) {
// }