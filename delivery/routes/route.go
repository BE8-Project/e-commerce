package routes

import (
	category "e-commerce/delivery/controllers/category"
	product "e-commerce/delivery/controllers/product"
	user "e-commerce/delivery/controllers/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, connUser user.UserController, connCategory category.CategoryController, connProduct product.ProductController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di API E-Commerce")
	})

	e.POST("/register", connUser.Register())
	e.POST("/login", connUser.Login())
	e.GET("/categories", connCategory.GetAll)
	e.GET("/products", connProduct.GetAll)
	e.GET("/products/:slug", connProduct.GetBySlug)
	e.GET("/products/category/:id", connProduct.GetByCategory)
	e.GET("/search", connProduct.GetBySearch)

	customer := e.Group("/users", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	customer.GET("/:username", connUser.GetUser)
	customer.PUT("/:username", connUser.Update())
	customer.DELETE("/:username", connUser.Delete())

	admin := e.Group("/admin", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	admin.POST("/categories", connCategory.Insert())

	merchant := e.Group("/merchants", middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")}))
	merchant.POST("/products", connProduct.Insert())
	merchant.PUT("/products/:slug", connProduct.Update())
	merchant.DELETE("/products/:slug", connProduct.Delete())
	merchant.GET("/products", connProduct.GetAllMerchant())
}