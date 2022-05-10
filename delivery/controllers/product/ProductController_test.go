package product

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"e-commerce/delivery/helpers/response"
	middlewares "e-commerce/delivery/middleware"
	"e-commerce/entity"

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

// Membuat request untuk insert product
func TestInsert(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       100000,
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockProduct{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "success create Product!", resp.Message)
		assert.Equal(t, map[string]interface {}(map[string]interface {}{"name":"Product 1", "created_at":"0001-01-01T00:00:00Z"}), resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       100000,
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockError{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "You are not allowed to access this resource", resp.Message)
	})

	t.Run("Status Invalid", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       "100000",
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockProduct{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "invalid request", resp.Message)
	})

	t.Run("Status BadRequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       100000,
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockErrorInput{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "kesalahan input", resp.Message)
	})
}

// Membuat request untuk get all product
func TestGetAll(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewProductController(&mockProduct{})

		controller.GetAll(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.Product
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success get all Product!", resp.Message)
		assert.Equal(t, []response.Product{
			{
				ID: 1,
				Name: "Product 1",
				Price: 100000,
				Image: "Product 1",
			},

		}, resp.Data)
	})

	t.Run("Status Not Found", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewProductController(&mockError{})

		controller.GetAll(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.Product
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Products not found", resp.Message)
		assert.Equal(t, []response.Product([]response.Product(nil)), resp.Data)
	})
}

// Membuat request untuk get product by slug
func TestGetBySlug(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewProductController(&mockProduct{})

		controller.GetBySlug(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	response.Product
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success get Product!", resp.Message)
		assert.Equal(t, response.Product{
				ID: 1,
				Name: "Product 1",
				Price: 100000,
				Image: "Product 1",
		}, resp.Data)
	})

	t.Run("Status Not Found", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewProductController(&mockError{})

		controller.GetBySlug(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Product not found", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

// Membuat request untuk update product
func TestUpdate(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       100000,
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockProduct{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success update Product!", resp.Message)
		assert.Equal(t, map[string]interface {}(map[string]interface {}{"name":"Product 1", "updated_at":"0001-01-01T00:00:00Z"}), resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       100000,
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockError{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "You are not allowed to access this resource", resp.Message)
	})

	t.Run("Status Invalid", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       "100000",
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockProduct{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "invalid request", resp.Message)
	})

	t.Run("Status BadRequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       100000,
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockErrorInput{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "kesalahan input", resp.Message)
	})

	t.Run("Status NotFound", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       "100000",
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockErrorNotfound{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Product not found", resp.Message)
		assert.Nil(t, resp.Data)
	})
}

// Membuat request untuk delete product
func TestDelete(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockProduct{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Delete())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success delete Product!", resp.Message)
		assert.Equal(t, map[string]interface {}(map[string]interface {}{"deleted_at":interface {}(nil), "name":""}), resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       100000,
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockError{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Delete())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "You are not allowed to access this resource", resp.Message)
	})

	t.Run("Status NotFound", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       "100000",
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockErrorNotfound{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Delete())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Product not found", resp.Message)
	})

	t.Run("Status BadRequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
			"price":       100000,
			"description": "Product 1",
			"image":       "Product 1",
			"stock":       1,
			"category_id": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/merchants/products")
		controller := NewProductController(&mockErrorInput{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "kesalahan input", resp.Message)
	})
}

// Membuat request untuk get product by category
func TestGetByCategory(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/products/:slug", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewProductController(&mockProduct{})

		controller.GetByCategory(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.Product
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success get all Product!", resp.Message)
		assert.Equal(t, []response.Product{
			{
				ID: 1,
				Name: "Product 1",
				Price: 100000,
				Image: "Product 1",
			},

		}, resp.Data)
	})

	t.Run("Status Not Found", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewProductController(&mockError{})

		controller.GetByCategory(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.Product
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Products not found", resp.Message)
		assert.Equal(t, []response.Product([]response.Product(nil)), resp.Data)
	})
}

// Membuat request untuk get product by Search
func TestGetBySearch(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/search", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewProductController(&mockProduct{})

		controller.GetBySearch(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.Product
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success get all Product!", resp.Message)
		assert.Equal(t, []response.Product{
			{
				ID: 1,
				Name: "Product 1",
				Price: 100000,
				Image: "Product 1",
			},

		}, resp.Data)
	})

	t.Run("Status Not Found", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/search", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewProductController(&mockError{})

		controller.GetBySearch(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.Product
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Products not found", resp.Message)
		assert.Equal(t, []response.Product([]response.Product(nil)), resp.Data)
	})
}

// Membuat request untuk get product by user
func TestGetByUser(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/merchants/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_merchant)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewProductController(&mockProduct{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetAllMerchant())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.ProductMerchant
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success get all Product!", resp.Message)
		assert.Equal(t, []response.ProductMerchant{
			{
				ID: 1,
				Name: "Product 1",
				Price: 100000,
				Image: "Product 1",
			},

		}, resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/merchants/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		controller := NewProductController(&mockError{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetAllMerchant())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 403, resp.Code)
		assert.Equal(t, "You are not allowed to access this resource", resp.Message)
	})

	t.Run("Status Not Found", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewProductController(&mockErrorNotfound{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.GetAllMerchant())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.ProductMerchant
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Products not found", resp.Message)
		assert.Equal(t, []response.ProductMerchant([]response.ProductMerchant(nil)), resp.Data)
	})
}

type mockProduct struct {
}

func (u *mockProduct) Insert(product *entity.Product) (response.InsertProduct, error) {
	return response.InsertProduct{
		Name: product.Name,
		CreatedAt: product.CreatedAt,
	}, nil
}

func (u *mockProduct) GetAll() []response.Product {
	return []response.Product{
		{
			ID: 1,
			Name: "Product 1",
			Price: 100000,
			Image: "Product 1",
		},
	}
}

func (u *mockProduct) GetBySlug(slug string) response.Product {
	return response.Product{
		ID: 1,
		Name: "Product 1",
		Price: 100000,
		Image: "Product 1",
	}
}

func (u *mockProduct) CheckSlug(slug string) bool {
	return true
}

func (u *mockProduct) Update(slug string, product *entity.Product) (response.UpdateProduct, error) {
	return response.UpdateProduct{
		Name: product.Name,
		UpdatedAt: product.UpdatedAt,
	}, nil
}

func (u *mockProduct) Delete(slug string) response.DeleteProduct {
	return response.DeleteProduct{}
}

func (u *mockProduct) GetByCategory(slug string) []response.Product {
	return []response.Product{
		{
			ID: 1,
			Name: "Product 1",
			Price: 100000,
			Image: "Product 1",
		},
	}
}

func (u *mockProduct) GetBySearch(search string) []response.Product {
	return []response.Product{
		{
			ID: 1,
			Name: "Product 1",
			Price: 100000,
			Image: "Product 1",
		},
	}
}

func (u *mockProduct) GetAllMerchant(user_id uint) []response.ProductMerchant {
	return []response.ProductMerchant{
		{
			ID: 1,
			Name: "Product 1",
			Price: 100000,
			Image: "Product 1",
		},
	}
}

func (u *mockProduct) CheckRole(id uint) bool {
	return true
}

// Bagian Error
type mockError struct {
}

func (u *mockError) Insert(product *entity.Product) (response.InsertProduct, error) {
	return response.InsertProduct{}, errors.New("kesalahan input")
}

func (u *mockError) GetAll() []response.Product {
	return []response.Product{}
}

func (u *mockError) GetBySlug(slug string) response.Product {
	return response.Product{}
}

func (u *mockError) CheckSlug(slug string) bool {
	return true
}

func (u *mockError) Update(slug string, product *entity.Product) (response.UpdateProduct, error) {
	return response.UpdateProduct{}, errors.New("kesalahan input")
}

func (u *mockError) Delete(slug string) response.DeleteProduct {
	return response.DeleteProduct{}
}

func (u *mockError) GetByCategory(slug string) []response.Product {
	return []response.Product{}
}

func (u *mockError) GetBySearch(search string) []response.Product {
	return []response.Product{}
}

func (u *mockError) GetAllMerchant(user_id uint) []response.ProductMerchant {
	return []response.ProductMerchant{}
}

func (u *mockError) CheckRole(id uint) bool {
	return false
}

type mockErrorInput struct {
}

func (u *mockErrorInput) Insert(product *entity.Product) (response.InsertProduct, error) {
	return response.InsertProduct{}, errors.New("kesalahan input")
}

func (u *mockErrorInput) GetAll() []response.Product {
	return []response.Product{}
}

func (u *mockErrorInput) GetBySlug(slug string) response.Product {
	return response.Product{}
}

func (u *mockErrorInput) CheckSlug(slug string) bool {
	return true
}

func (u *mockErrorInput) Update(slug string, product *entity.Product) (response.UpdateProduct, error) {
	return response.UpdateProduct{}, errors.New("kesalahan input")
}

func (u *mockErrorInput) Delete(slug string) response.DeleteProduct {
	return response.DeleteProduct{}
}

func (u *mockErrorInput) GetByCategory(slug string) []response.Product {
	return []response.Product{}
}

func (u *mockErrorInput) GetBySearch(search string) []response.Product {
	return []response.Product{}
}

func (u *mockErrorInput) GetAllMerchant(user_id uint) []response.ProductMerchant {
	return []response.ProductMerchant{}
}

func (u *mockErrorInput) CheckRole(id uint) bool {
	return true
}

type mockErrorNotfound struct {
}

func (u *mockErrorNotfound) Insert(product *entity.Product) (response.InsertProduct, error) {
	return response.InsertProduct{}, errors.New("kesalahan input")
}

func (u *mockErrorNotfound) GetAll() []response.Product {
	return []response.Product{}
}

func (u *mockErrorNotfound) GetBySlug(slug string) response.Product {
	return response.Product{}
}

func (u *mockErrorNotfound) CheckSlug(slug string) bool {
	return false
}

func (u *mockErrorNotfound) Update(slug string, product *entity.Product) (response.UpdateProduct, error) {
	return response.UpdateProduct{}, errors.New("kesalahan input")
}

func (u *mockErrorNotfound) Delete(slug string) response.DeleteProduct {
	return response.DeleteProduct{}
}

func (u *mockErrorNotfound) GetByCategory(slug string) []response.Product {
	return []response.Product{}
}

func (u *mockErrorNotfound) GetBySearch(search string) []response.Product {
	return []response.Product{}
}

func (u *mockErrorNotfound) GetAllMerchant(user_id uint) []response.ProductMerchant {
	return []response.ProductMerchant{}
}

func (u *mockErrorNotfound) CheckRole(id uint) bool {
	return true
}