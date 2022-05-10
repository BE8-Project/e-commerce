package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"
)

type CategoryModel interface {
	Insert(category *entity.Category) (response.InsertCategory, error)
	GetAll() []response.Category
}