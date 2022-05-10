package category

import (
	"e-commerce/delivery/helpers/response"
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
	token_user string
)

func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token_admin, _ = middlewares.CreateToken(1, "admin")
		token_user, _ = middlewares.CreateToken(3, "user")
	})
}

func TestInsert(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/categories")
		controller := NewCategoryController(&mockCategory{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Insert())(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "success create Category!", resp.Message)
		assert.Equal(t, map[string]interface {}(map[string]interface {}{"name":"Product 1", "created_at":"0001-01-01T00:00:00Z"}), resp.Data)
	})

	t.Run("Status Forbidden", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "Product 1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/categories")
		controller := NewCategoryController(&mockErrorForbidden{})

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
			"name": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/categories")
		controller := NewCategoryController(&mockCategory{})

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
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_admin)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/admin/categories")
		controller := NewCategoryController(&mockError{})

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

// Membuat request get all category
func TestGetAll(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewCategoryController(&mockCategory{})

		controller.GetAll(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.Category
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success get all Category!", resp.Message)
		assert.Equal(t, []response.Category{
			{
				Name: "Product 1",
				Slug: "product-1",
			},

		}, resp.Data)
	})

	t.Run("Status Not Found", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewCategoryController(&mockError{})

		controller.GetAll(context)

		type Response struct {
			Code    int   `json:"code"`
			Message string `json:"message"`
			Data 	[]response.Category
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "Categories not found", resp.Message)
		assert.Equal(t, []response.Category([]response.Category(nil)), resp.Data)
	})
}

type mockCategory struct {}

func (mc *mockCategory) Insert(category *entity.Category) (response.InsertCategory, error) {
	return response.InsertCategory{
		Name: category.Name,
		CreatedAt: category.CreatedAt,
	}, nil
}

func (mc *mockCategory) GetAll() []response.Category {
	return []response.Category{
		{
			Name: "Product 1",
			Slug: "product-1",
		},
	}
}

type mockError struct {}

func (mc *mockError) Insert(category *entity.Category) (response.InsertCategory, error) {
	return response.InsertCategory{}, errors.New("kesalahan input")
}

func (mc *mockError) GetAll() []response.Category {
	return []response.Category{}
}

type mockErrorForbidden struct {}

func (mc *mockErrorForbidden) Insert(category *entity.Category) (response.InsertCategory, error) {
	return response.InsertCategory{}, errors.New("You are not allowed to access this resource")
}

func (mc *mockErrorForbidden) GetAll() []response.Category {
	return []response.Category{}
}