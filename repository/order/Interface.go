package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"
)

type OrderModel interface {
	CheckRole(id uint) bool
	Insert(order *entity.Order) (response.InsertOrder, error)
	Update(order_id string, order *entity.Order) (response.UpdateOrder, error)
}