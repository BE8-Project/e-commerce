package routes

import (
	category "e-commerce/delivery/controllers/category"
	user "e-commerce/delivery/controllers/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, connUser user.UserController, connCategory category.CategoryController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	e.POST("/register", connUser.Register())
	e.POST("/login", connUser.Login())
	e.GET("/categories", connCategory.GetAll)

	customer := e.Group("/users", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	customer.GET("/:username", connUser.GetUser)
	customer.PUT("/:username", connUser.Update())
	customer.DELETE("/:username", connUser.Delete())

	admin := e.Group("/admin", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	admin.POST("/categories", connCategory.Insert())
}