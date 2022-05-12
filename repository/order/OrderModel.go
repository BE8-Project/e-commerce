package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"

	"gorm.io/gorm"
)

type orderModel struct {
	DB *gorm.DB
}

func NewOrderModel(db *gorm.DB) *orderModel {
	return &orderModel{
		DB: db,
	}
}

func (om *orderModel) CheckRole(id uint) bool {
	var user entity.User
	om.DB.Where("id = ?", id).Find(&user)

	if user.Role == 0 {
		return true
	}

	return false
}

func (om *orderModel) Insert(order *entity.Order) (response.InsertOrder, error) {
	record := om.DB.Create(order)

	if record.RowsAffected == 0 {
		return response.InsertOrder{}, record.Error
	} else {
		return response.InsertOrder{
			OrderID: order.TrackingNumber,
			Total: order.Total,
			CreatedAt: order.CreatedAt,
		}, nil
	}
}

func (om *orderModel) Update(order_id string, order *entity.Order) (response.UpdateOrder, error) {
	record := om.DB.Where("tracking_number = ?", order_id).Updates(&order)

	if record.RowsAffected == 0 {
		return response.UpdateOrder{}, record.Error
	} else {
		return response.UpdateOrder{
			OrderID: order.TrackingNumber,
			Status: 	   order.Status,
			UpdatedAt: 	order.UpdatedAt,
		}, nil
	}
}