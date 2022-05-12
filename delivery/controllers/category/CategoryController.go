package category

import (
	"e-commerce/delivery/helpers/request"
	"e-commerce/delivery/helpers/response"
	middlewares "e-commerce/delivery/middleware"
	"e-commerce/entity"
	repoCategory "e-commerce/repository/category"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type categoryController struct {
	Connect repoCategory.CategoryModel
	Validate *validator.Validate
}

func NewCategoryController(conn repoCategory.CategoryModel) *categoryController {
	return &categoryController{
		Connect: conn,
		Validate: validator.New(),
	}
}

func (cc *categoryController) Insert() echo.HandlerFunc {
	return func (c echo.Context) error {
		userID := middlewares.ExtractTokenUserId(c)

		if uint(userID) != 1 {
			return c.JSON(http.StatusForbidden, response.StatusForbidden("You are not allowed to access this resource"))
		}
	
		var request request.InsertCategory
	
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := cc.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
	
		category := entity.Category{
			Name: request.Name,
		}
	
		result, err := cc.Connect.Insert(&category)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}
	
		return c.JSON(http.StatusCreated, response.StatusCreated("success create Category!", result))
	}
}

func (cc *categoryController) GetAll(c echo.Context) error {
	categories := cc.Connect.GetAll()

	if len(categories) == 0 {
		return c.JSON(http.StatusNotFound, response.StatusNotFound("Categories not found"))
	}

	return c.JSON(http.StatusOK, response.StatusOK("success get all Category!", categories))
}