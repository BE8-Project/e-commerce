package main

import (
	"e-commerce/config"
	"e-commerce/delivery/controllers/cart"
	category "e-commerce/delivery/controllers/category"
	product "e-commerce/delivery/controllers/product"
	user "e-commerce/delivery/controllers/user"
	"e-commerce/delivery/routes"

	cartModels "e-commerce/repository/cart"
	categoryModel "e-commerce/repository/category"
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

	cartModel := cartModels.NewCartModel(db)
	cartController := cart.NewCartController(cartModel)

	routes.Route(e, userController, categoryController, productController, cartController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
