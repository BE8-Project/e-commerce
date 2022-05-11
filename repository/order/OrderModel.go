package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"

	"gorm.io/gorm"
)

type orderModel struct {
	DB *gorm.DB
}

func NewCategoryModel(db *gorm.DB) *orderModel {
	return &orderModel{
		DB: db,
	}
}

func (om *orderModel) Insert(order *entity.Order) (response.InsertOrder, error) {
	record := om.DB.Create(order)

	if record.RowsAffected == 0 {
		return response.InsertOrder{}, record.Error
	} else {
		return response.InsertOrder{
			TrackingNumber: order.TrackingNumber,
			PaymentType:    order.PaymentType,
			Total:          order.Total,
			Status:         order.Status,
			CreatedAt: order.CreatedAt,
		}, nil
	}
}