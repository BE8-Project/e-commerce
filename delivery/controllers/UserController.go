package user

import (
	"e-commerce/delivery/helpers/request"
	"e-commerce/delivery/helpers/response"
	middlewares "e-commerce/delivery/middleware"
	"e-commerce/entity"
	repoUser "e-commerce/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	Connect repoUser.UserModel
}

func NewUserController(conn repoUser.UserModel) *userController {
	return &userController{
		Connect: conn,
	}
}

func (u *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request request.InsertUser

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		user := entity.User{
			Name:     request.Name,
			Username: request.Username,
			Email:    request.Email,
			HP:       request.HP,
			Password: request.Password,
			Role:     request.Role,
		}

		result, err := u.Connect.Insert(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success register User!", result))
	}
}

func (u *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request request.Login

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		var input []string
		input = append(input, request.Email, request.Username, request.HP)

		login, err := u.Connect.Login(input, request.Password)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.StatusUnauthorized(err))
		}

		result := response.LoginDetail{User: login}
		if result.Token == "" {
			token, _ := middlewares.CreateToken(login.ID, login.Username)
			result.Token = token
		}

		return c.JSON(http.StatusOK, response.StatusOK("success login!", result))
	}
}

func (u *userController) GetUser(c echo.Context) error {
	username := c.Param("username")

	result := u.Connect.GetOne(username)

	if result.Name == "" {
		return c.JSON(http.StatusNotFound, response.StatusNotFound("User not found"))
	}

	return c.JSON(http.StatusOK, response.StatusOK("success get user!", result))
}

func (u *userController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")

		var request request.InsertUser
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		user := entity.User{
			Name:     request.Name,
			Username: request.Username,
			Email:    request.Email,
			HP:       request.HP,
			Password: request.Password,
			Role:     0,
		}

		result, err := u.Connect.Update(&user, username)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		return c.JSON(http.StatusOK, response.StatusOK("success update Product!", result))
	}
}

func (u *userController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		UsrParam := c.Param("username")

		// if username != UsrParam {
		// 	log.Warn("Status For bidden")
		// 	return c.JSON(http.StatusBadRequest, response.StatusForbidden("You are not allowed to access this resource"))
		// }

		result := u.Connect.Delete(UsrParam)

		return c.JSON(http.StatusOK, response.StatusOK("success delete User!", result))
	}
}
