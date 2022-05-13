package cart

import (
	"e-commerce/delivery/helpers/request"
	middlewares "e-commerce/delivery/middleware"
	"e-commerce/entity"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	token_admin string
	token_user  string
)

func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token_admin, _ = middlewares.CreateToken(3, "admin")
		token_user, _ = middlewares.CreateToken(3, "user")
	})
}

func TestInsert(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_id": 1,
			"quantity":   3,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/carts")
		controller := NewCartController(&mockCart{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "success add to Cart!", resp.Message)
		assert.Equal(t, "Succes", resp.Data)
	})

	t.Run("Status Invalid", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_id": 1,
			"quantity":   "3",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/carts")
		controller := NewCartController(&mockCart{})

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

	t.Run("Status BadRequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_id": 1,
			"quantity":   "3",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/carts")
		controller := NewCartController(&mockErorCart{})

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
}
func TestGetAll(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewCartController(&mockCart{})
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetAll())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    []entity.Cart
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success get cart!", resp.Message)
		assert.Equal(t, []entity.Cart{
			{
				UserID:    3,
				Name:      "laptop lenovo",
				Quantity:  3,
				Price:     1000,
				Image:     "laptop",
				ProductID: 2,
			},
		}, resp.Data)
	})

	t.Run("Status Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewCartController(&mockErorCart{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetAll())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    []entity.Cart
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Cart not found", resp.Message)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_id": 1,
			"quantity":   3,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewCartController(&mockBiddenErorCart{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetAll())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "You are not allowed to access this resource", resp.Message)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"quantity": 3,
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/carts/:id")
		context.SetParamNames("id")
		context.SetParamValues("3")
		controller := NewCartController(&mockCart{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "success update to Cart!", resp.Message)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"quantity": 3,
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/carts/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		controller := NewCartController(&mockBiddenErorCart{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "You are not allowed to access this resource", resp.Message)
	})

	t.Run("Status Invalid", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"quantity": "3",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/carts/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		controller := NewCartController(&mockCart{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

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

	t.Run("Status BadRequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewCartController(&mockCart{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "", resp.Message)
	})
}

// // Membuat request untuk delete product
func TestDelete(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/carts/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		controller := NewCartController(&mockCart{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "success delete to Cart!", resp.Message)
	})
}

type mockCart struct{}

func (mc *mockCart) Checkid(id uint, idcart uint) error {
	return nil
}
func (mc *mockCart) Insert(cart request.InsertCart, idUser uint) error {
	return nil
}
func (mc *mockCart) GetAll(id uint) ([]entity.Cart, error) {
	return []entity.Cart{
		{
			UserID:    3,
			Name:      "laptop lenovo",
			Quantity:  3,
			Price:     1000,
			Image:     "laptop",
			ProductID: 2,
		},
	}, nil
}
func (mc *mockCart) Update(cart request.UpdateCart, id uint) error {
	return nil
}
func (mc *mockCart) Delete(id uint) error {
	return nil
}

type mockErorCart struct{}

func (mc *mockErorCart) Checkid(id uint, idcart uint) error {
	return errors.New("eror")
}
func (mc *mockErorCart) Insert(cart request.InsertCart, idUser uint) error {
	return errors.New("Eror Insert")
}
func (mc *mockErorCart) GetAll(id uint) ([]entity.Cart, error) {
	return []entity.Cart{}, nil
}
func (mc *mockErorCart) Update(cart request.UpdateCart, id uint) error {
	return errors.New("Eror Update")
}
func (mc *mockErorCart) Delete(id uint) error {
	return errors.New("Eror Delete")
}

type mockBiddenErorCart struct{}

func (mc *mockBiddenErorCart) Checkid(id uint, idcart uint) error {
	return errors.New("eror")
}
func (mc *mockBiddenErorCart) Insert(cart request.InsertCart, idUser uint) error {
	return errors.New("Eror Insert")
}
func (mc *mockBiddenErorCart) GetAll(id uint) ([]entity.Cart, error) {
	return []entity.Cart{}, errors.New("Eror Insert")
}
func (mc *mockBiddenErorCart) Update(cart request.UpdateCart, id uint) error {
	return errors.New("Eror Update")
}
func (mc *mockBiddenErorCart) Delete(id uint) error {
	return errors.New("Eror Delete")
}

type mockErorInputCart struct{}

func (mc *mockErorInputCart) Checkid(id uint, idcart uint) error {
	return nil
}
func (mc *mockErorInputCart) Insert(cart request.InsertCart, idUser uint) error {
	return errors.New("Eror Insert")
}
func (mc *mockErorInputCart) GetAll(id uint) ([]entity.Cart, error) {
	return []entity.Cart{}, errors.New("eror Get All")
}
func (mc *mockErorInputCart) Update(cart request.UpdateCart, id uint) error {
	return errors.New("Eror Update")
}
func (mc *mockErorInputCart) Delete(id uint) error {
	return errors.New("Eror Delete")
}
