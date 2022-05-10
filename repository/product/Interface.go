package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"
)

type ProductModel interface {
	CheckRole(id uint) bool
	CheckSlug(slug string) bool
	Insert(product *entity.Product) (response.InsertProduct, error)
	GetAll() []response.Product
	GetBySlug(slug string) response.Product
	Update(slug string, product *entity.Product) (response.UpdateProduct, error)
	Delete(slug string) response.DeleteProduct
	GetByCategory(category_id string) []response.Product
	GetBySearch(search string) []response.Product
	GetAllMerchant(user_id uint) []response.ProductMerchant
}