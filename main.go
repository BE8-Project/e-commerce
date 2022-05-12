package main

import (
	"e-commerce/config"
	address "e-commerce/delivery/controllers/address"
	"e-commerce/delivery/controllers/cart"
	category "e-commerce/delivery/controllers/category"
	"e-commerce/delivery/controllers/order"
	product "e-commerce/delivery/controllers/product"
	user "e-commerce/delivery/controllers/user"
	"e-commerce/delivery/routes"

	addressModel "e-commerce/repository/address"
	cartModel "e-commerce/repository/cart"
	categoryModel "e-commerce/repository/category"
	orderModel "e-commerce/repository/order"
	productModel "e-commerce/repository/product"
	userModel "e-commerce/repository/user"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	config.AutoMigrate(db)

	e := echo.New()

	userModel := userModel.NewUserModel(db)
	userController := user.NewUserController(userModel)

	categoryModel := categoryModel.NewCategoryModel(db)
	categoryController := category.NewCategoryController(categoryModel)

	productModel := productModel.NewProductModel(db)
	productController := product.NewProductController(productModel)

	cartModel := cartModel.NewCartModel(db)
	cartController := cart.NewCartController(cartModel)

	addressModel := addressModel.NewAddressModel(db)
	addressController := address.NewAddressController(addressModel)

	orderModel := orderModel.NewOrderModel(db)
	orderController := order.NewOrderController(orderModel)

	routes.Route(e, userController, categoryController, productController, cartController, addressController, orderController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
