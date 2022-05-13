package user

import (
	"e-commerce/delivery/helpers/response"
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

var token_user string

func TestRegister(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "fajar",
			"username": "fajar123",
			"email":    "fajar123@gmail.com",
			"hp":       "098765433212",
			"password": "qwerty",
			"role":     1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")
		controller := NewUserController(&mockUser{})
		controller.Register()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 201, resp.Code)
		assert.Equal(t, "success register User!", resp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}{"name": "fajar", "username": "fajar123", "email": "fajar123@gmail.com", "hp": "098765433212", "created_at": "0001-01-01T00:00:00Z"}), resp.Data)
	})
	t.Run("Status Invalid", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "fajar",
			"username": "fajar123",
			"email":    "fajar123@gmail.com",
			"hp":       0234567723,
			"password": "qwerty",
			"role":     0,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")
		controller := NewUserController(&mockUser{})
		controller.Register()(context)

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
			"name":     "fajar",
			"username": "fajar123",
			"email":    "fajar123@gmail.com",
			"hp":       "098765433212",
			"password": "qwerty",
			"role":     1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")
		controller := NewUserController(&mockErorInputUser{})
		controller.Register()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "kesalahan input", resp.Message)
	})
	t.Run("Status BadRequest", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "fajar",
			"username": "fajar123",
			"email":    "fajar123@gmail.com",
			"hp":       "098765433212",
			"password": "qwerty",
			"role":     1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")
		controller := NewUserController(&mockErorInputUser{})
		controller.Register()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "kesalahan input", resp.Message)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Status Ok", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"hp":       "098765433212",
			"username": "fajar12",
			"email":    "fajar123@gmail.com",
			"password": "123456",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")
		controller := NewUserController(&mockUser{})
		controller.Login()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success login!", resp.Message)
		data := resp.Data.(map[string]interface{})
		token_user = data["token"].(string)
	})
	t.Run("Status Invalid", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"hp":       "098765433212",
			"username": "fajar12",
			"email":    "fajar123@gmail.com",
			"password": 123456,
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")
		controller := NewUserController(&mockUser{})
		controller.Register()(context)

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
	t.Run("Status Unauthorized", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"hp":       "098765433212",
			"username": "fajar12",
			"email":    "fajar123@gmail.com",
			"password": "123456",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/login")
		controller := NewUserController(&mockAutorizUser{})
		controller.Login()(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 401, resp.Code)
		assert.Equal(t, "user or password is wrong", resp.Message)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewUserController(&mockUser{})

		controller.GetUser(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success get User!", resp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}{"name": "fajar", "username": "fajar123", "email": "fajar123@gmail.com", "hp": "098765433212", "created_at": "0001-01-01T00:00:00Z"}), resp.Data)
	})
	t.Run("Status Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		controller := NewUserController(&mockErorUser{})

		controller.GetUser(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 404, resp.Code)
		assert.Equal(t, "User not found", resp.Message)
		assert.Nil(t, resp.Data)
	})

}
func TestUpdate(t *testing.T) {
	t.Run("Status OK", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "fajar",
			"username": "fajar123",
			"email":    "fajar123@gmail.com",
			"hp":       "098765433212",
			"password": "qwerty",
			"role":     0,
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:username")
		context.SetParamNames("username")
		context.SetParamValues("fajar123")
		controller := NewUserController(&mockUser{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success update User!", resp.Message)
		assert.Equal(t, map[string]interface{}{"name": "fajar", "updated_at": "0001-01-01T00:00:00Z"}, resp.Data)
	})
	t.Run("invalid request", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "fajar",
			"username": "fajar123",
			"email":    "fajar123@gmail.com",
			"hp":       98765433212,
			"password": "qwerty",
			"role":     0,
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:username")
		context.SetParamNames("username")
		context.SetParamValues("fajar123")
		controller := NewUserController(&mockUser{})

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
	t.Run("Status Invalid", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "fajar",
			"username": "fajar123",
			"email":    "fajar123@gmail.com",
			"hp":       "98765433212",
			"password": "qwerty",
			"role":     0,
		})

		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:username")
		context.SetParamNames("username")
		context.SetParamValues("fajar123")
		controller := NewUserController(&mockErorInputUser{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Update())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "kesalahan input", resp.Message)
	})

}
func TestDelete(t *testing.T) {
	t.Run("Status OK Delet", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token_user)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")
		controller := NewUserController(&mockUser{})

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")})(controller.Delete())(context)

		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    interface{}
		}

		var resp Response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, "success delete User!", resp.Message)
		assert.Equal(t, map[string]interface{}(map[string]interface{}{"deleted_at": interface{}(nil), "name": ""}), resp.Data)
	})
}

type mockUser struct{}

func (u *mockUser) Insert(user *entity.User) (response.User, error) {
	return response.User{
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		HP:        user.HP,
		CreatedAt: user.CreatedAt,
	}, nil
}
func (u *mockUser) Login(custom []string, password string) (response.Login, error) {
	return response.Login{
		ID:       2,
		Name:     "fajar",
		Username: "fajar12",
	}, nil
}

func (u *mockUser) GetOne(username string) response.User {
	return response.User{
		Name:     "fajar",
		Username: "fajar123",
		Email:    "fajar123@gmail.com",
		HP:       "098765433212",
	}
}
func (u *mockUser) Update(newUser *entity.User, username string) (response.UpdateUser, error) {
	return response.UpdateUser{
		Name:      newUser.Name,
		UpdatedAt: newUser.UpdatedAt,
	}, nil
}

func (u *mockUser) Delete(username string) response.DeleteUser {
	return response.DeleteUser{}
}

type mockErorInputUser struct{}

func (u *mockErorInputUser) Insert(user *entity.User) (response.User, error) {
	return response.User{}, errors.New("kesalahan input")
}
func (u *mockErorInputUser) Login(custom []string, password string) (response.Login, error) {
	return response.Login{
		ID:       0,
		Name:     "fajar",
		Username: "fajar12",
	}, nil
}

func (u *mockErorInputUser) GetOne(username string) response.User {
	return response.User{
		Name:     "fajar",
		Username: "fajar123",
		Email:    "fajar123@gmail.com",
		HP:       "098765433212",
	}
}
func (u *mockErorInputUser) Update(newUser *entity.User, username string) (response.UpdateUser, error) {
	return response.UpdateUser{}, errors.New("kesalahan input")
}

func (u *mockErorInputUser) Delete(username string) response.DeleteUser {
	return response.DeleteUser{}
}

type mockErorUser struct{}

func (u *mockErorUser) Insert(user *entity.User) (response.User, error) {
	return response.User{}, errors.New("kesalahan input")
}
func (u *mockErorUser) Login(custom []string, password string) (response.Login, error) {
	return response.Login{
		ID:       0,
		Name:     "fajar",
		Username: "fajar12",
	}, nil
}

func (u *mockErorUser) GetOne(username string) response.User {
	return response.User{}
}
func (u *mockErorUser) Update(newUser *entity.User, username string) (response.UpdateUser, error) {
	return response.UpdateUser{
		Name:      newUser.Name,
		UpdatedAt: newUser.UpdatedAt,
	}, nil
}

func (u *mockErorUser) Delete(username string) response.DeleteUser {
	return response.DeleteUser{}
}

type mockAutorizUser struct{}

func (u *mockAutorizUser) Insert(user *entity.User) (response.User, error) {
	return response.User{}, errors.New("kesalahan input")
}
func (u *mockAutorizUser) Login(custom []string, password string) (response.Login, error) {
	return response.Login{}, errors.New("user or password is wrong")
}

func (u *mockAutorizUser) GetOne(username string) response.User {
	return response.User{
		Name:     "fajar",
		Username: "fajar123",
		Email:    "fajar123@gmail.com",
		HP:       "098765433212",
	}
}
func (u *mockAutorizUser) Update(newUser *entity.User, username string) (response.UpdateUser, error) {
	return response.UpdateUser{
		Name:      newUser.Name,
		UpdatedAt: newUser.UpdatedAt,
	}, nil
}

func (u *mockAutorizUser) Delete(username string) response.DeleteUser {
	return response.DeleteUser{}
}
