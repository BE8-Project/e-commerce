package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"
	"strings"

	"gorm.io/gorm"
)

type categoryModel struct {
	DB *gorm.DB
}

func NewCategoryModel(db *gorm.DB) *categoryModel {
	return &categoryModel{
		DB: db,
	}
}

func (cm *categoryModel) Insert(category *entity.Category) (response.InsertCategory, error) {
	category.Name = strings.Title(strings.ToLower(category.Name))
	category.Slug = strings.Replace(strings.ToLower(category.Name), " ", "-", -1)
	
	record := cm.DB.Create(category)

	if record.RowsAffected == 0 {
		return response.InsertCategory{}, record.Error
	} else {
		return response.InsertCategory{
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
		}, nil
	}
}

func (cm *categoryModel) GetAll() []response.Category {
	var categories []response.Category
	record := cm.DB.Find(&categories)

	if record.RowsAffected == 0 {
		return []response.Category{}
	} else {
		return categories
	}
}