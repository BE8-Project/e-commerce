package routes

import (
	address "e-commerce/delivery/controllers/address"
	"e-commerce/delivery/controllers/cart"
	category "e-commerce/delivery/controllers/category"
	"e-commerce/delivery/controllers/order"
	product "e-commerce/delivery/controllers/product"
	user "e-commerce/delivery/controllers/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, connUser user.UserController, connCategory category.CategoryController, connProduct product.ProductController, connCart cart.CartController, connAddress address.AddressController, connOrder order.OrderController) {
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
	customer.GET("/carts", connCart.GetAll)
	customer.PUT("/carts/:id", connCart.Update())
	customer.DELETE("/carts/:id", connCart.Delete())

	customer.POST("/address", connAddress.Insert())

	admin := e.Group("/admin", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	admin.POST("/categories", connCategory.Insert())

	merchant := e.Group("/merchants", middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("$p4ssw0rd")}))
	merchant.POST("/products", connProduct.Insert())
	merchant.PUT("/products/:slug", connProduct.Update())
	merchant.DELETE("/products/:slug", connProduct.Delete())
	merchant.GET("/products", connProduct.GetAllMerchant())
	customer.POST("/carts", connCart.Insert())

	order := e.Group("/orders", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	order.POST("", connOrder.Insert())
	order.GET("/:order_id", connOrder.GetStatus())
	order.GET("/:order_id/cancel", connOrder.Cancel())
}
