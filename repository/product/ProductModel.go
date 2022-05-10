package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"
	"strings"

	"gorm.io/gorm"
)

type productModel struct {
	DB *gorm.DB
}

func NewProductModel(db *gorm.DB) *productModel {
	return &productModel{
		DB: db,
	}
}

func (u *productModel) CheckRole(id uint) bool {
	var user entity.User
	u.DB.Where("id = ?", id).Find(&user)

	if user.Role == 1 {
		return true
	}

	return false
}

func (u *productModel) Insert(product *entity.Product) (response.InsertProduct, error) {
	product.Name = strings.Title(strings.ToLower(product.Name))
	product.Slug = strings.Replace(strings.ToLower(product.Name), " ", "-", -1)
	
	record := u.DB.Create(product)

	if record.RowsAffected == 0 {
		return response.InsertProduct{}, record.Error
	} else {
		return response.InsertProduct{
			Name: product.Name,
			CreatedAt: product.CreatedAt,
		}, nil
	}
}

func (u *productModel) GetAll() []response.Product {
	var products []response.Product
	
	record := u.DB.Where("stock > ?", 0).Find(&products)

	if record.RowsAffected == 0 {
		return []response.Product{}
	} else {
		return products
	}
}

func (u *productModel) GetBySlug(slug string) response.Product {
	var product response.Product

	record := u.DB.Where("slug = ? and stock > ?", slug, 0).Find(&product)

	if record.RowsAffected == 0 {
		return response.Product{}
	} else {
		return product
	}
}

func (u *productModel) CheckSlug(slug string) bool {
	var product entity.Product
	u.DB.Where("slug = ?", slug).Find(&product)

	if product.ID == 0 {
		return false
	}

	return true
}

func (u *productModel) Update(slug string, product *entity.Product) (response.UpdateProduct, error) {
	product.Name = strings.Title(strings.ToLower(product.Name))
	product.Slug = strings.Replace(strings.ToLower(product.Name), " ", "-", -1)

	record := u.DB.Where("slug = ?", slug).Updates(product)

	if record.RowsAffected == 0 {
		return response.UpdateProduct{}, record.Error
	} else {
		return response.UpdateProduct{
			Name: product.Name,
			UpdatedAt: product.UpdatedAt,
		}, nil
	}
}

func (u *productModel) Delete(slug string) response.DeleteProduct {
	var product entity.Product

	record := u.DB.Where("slug = ?", slug).Delete(&product)

	if record.RowsAffected == 0 {
		return response.DeleteProduct{}
	} else {
		return response.DeleteProduct{
			Name: product.Name,
			DeletedAt: product.DeletedAt,
		}
	}
}

func (u *productModel) GetByCategory(category_id string) []response.Product {
	var products []response.Product

	record := u.DB.Where("category_id = ? and stock > ?", category_id, 0).Find(&products)

	if record.RowsAffected == 0 {
		return []response.Product{}
	} else {
		return products
	}
}

func (u *productModel) GetBySearch(search string) []response.Product {
	var products []response.Product

	record := u.DB.Where("name LIKE ? and stock > ?", "%"+search+"%", 0).Find(&products)

	if record.RowsAffected == 0 {
		return []response.Product{}
	} else {
		return products
	}
}

func (u *productModel) GetAllMerchant(user_id uint) []response.ProductMerchant {
	var products []entity.Product

	record := u.DB.Where("user_id = ? and stock > ?", user_id, 0).Find(&products)

	var results []response.ProductMerchant
	for _, product := range products {
		results = append(results, response.ProductMerchant{
			ID: product.ID,
			Name: product.Name,
			Price: product.Price,
			Stock: product.Stock,
			Image: product.Image,
		})
	}

	if record.RowsAffected == 0 {
		return []response.ProductMerchant{}
	} else {
		return results
	}
}