package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"
)

type AddressModel interface {
	Insert(address *entity.Address) (response.InsertAddress, error)
	GetByUserID(userID uint) []response.Address
}