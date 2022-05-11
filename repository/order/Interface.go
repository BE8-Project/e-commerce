package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"
)

type OrderModel interface {
	Insert(order *entity.Order) (response.InsertOrder, error)
}