package product

import (
	"e-commerce/delivery/helpers/request"
	"e-commerce/delivery/helpers/response"
	middlewares "e-commerce/delivery/middleware"
	"e-commerce/entity"
	repoProduct "e-commerce/repository/product"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type productController struct {
	Connect repoProduct.ProductModel
	Validate *validator.Validate
}

func NewProductController(conn repoProduct.ProductModel) *productController {
	return &productController{
		Connect: conn,
		Validate: validator.New(),
	}
}

func (u *productController) Insert() echo.HandlerFunc  {
	return func (c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)
		var request request.InsertProduct

		if !u.Connect.CheckRole(uint(userID)) {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := u.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		product := entity.Product{
			Name:        request.Name,
			Price:       request.Price,
			Description: request.Description,
			Image:       request.Image,
			Stock:       request.Stock,
			CategoryID:  request.CategoryID,
			UserID:      uint(userID),
		}

		result, err := u.Connect.Insert(&product)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success create Product!", result))
	}
}

func (u *productController) GetAll(c echo.Context) error {
	results := u.Connect.GetAll()

	if len(results) == 0 {
		return c.JSON(http.StatusNotFound, response.StatusNotFound("Products not found"))
	}

	return c.JSON(http.StatusOK, response.StatusOK("success get all Product!", results))
}

func (u *productController) GetBySlug(c echo.Context) error {
	slug := c.Param("slug")

	result := u.Connect.GetBySlug(slug)

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, response.StatusNotFound("Product not found"))
	}

	return c.JSON(http.StatusOK, response.StatusOK("success get Product!", result))
}

func (u *productController) Update() echo.HandlerFunc {
	return func (c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)
		slug := c.Param("slug")
		var request request.UpdateProduct
	
		if !u.Connect.CheckRole(uint(userID)) {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}
	
		if !u.Connect.CheckSlug(slug) {
			return c.JSON(http.StatusNotFound, response.StatusNotFound("Product not found"))
		}
	
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := u.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
	
		product := entity.Product{
			Name:        request.Name,
			Price:       request.Price,
			Description: request.Description,
			Image:       request.Image,
			Stock:       request.Stock,
			CategoryID:  request.CategoryID,
		}
	
		result, err := u.Connect.Update(slug, &product)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
	
		return c.JSON(http.StatusOK, response.StatusOK("success update Product!", result))
	}
}

func (u *productController) Delete() echo.HandlerFunc {
	return func (c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)
		slug := c.Param("slug")
	
		if !u.Connect.CheckRole(uint(userID)) {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}
	
		if !u.Connect.CheckSlug(slug) {
			return c.JSON(http.StatusNotFound, response.StatusNotFound("Product not found"))
		}
	
		result := u.Connect.Delete(slug)
		
		return c.JSON(http.StatusOK, response.StatusOK("success delete Product!", result))
	}
}

func (u *productController) GetByCategory(c echo.Context) error {
	id := c.Param("id")

	results := u.Connect.GetByCategory(id)

	if len(results) == 0 {
		return c.JSON(http.StatusNotFound, response.StatusNotFound("Products not found"))
	}

	return c.JSON(http.StatusOK, response.StatusOK("success get all Product!", results))
}

func (u *productController) GetBySearch(c echo.Context) error {
	search := c.QueryParam("name")

	results := u.Connect.GetBySearch(search)

	if len(results) == 0 {
		return c.JSON(http.StatusNotFound, response.StatusNotFound("Products not found"))
	}

	return c.JSON(http.StatusOK, response.StatusOK("success get all Product!", results))
}

func (u *productController) GetAllMerchant() echo.HandlerFunc {
	return func (c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)

		if !u.Connect.CheckRole(uint(userID)) {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}

		results := u.Connect.GetAllMerchant(uint(userID))

		if len(results) == 0 {
			return c.JSON(http.StatusNotFound, response.StatusNotFound("Products not found"))
		}

		return c.JSON(http.StatusOK, response.StatusOK("success get all Product!", results))
	}
}
