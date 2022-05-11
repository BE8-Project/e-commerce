package cart

import (
	"e-commerce/delivery/helpers/request"
	"e-commerce/entity"
)

type CartModel interface {
	Checkid(id uint, idcart uint) bool
	Insert(cart request.InsertCart, idUser uint) error
	GetAll(id uint) ([]entity.Cart, error)
	Update(cart request.UpdateCart, id uint) error
	Delete(id uint) error
}
