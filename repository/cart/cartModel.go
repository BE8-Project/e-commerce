package cart

import (
	"e-commerce/delivery/helpers/request"
	"e-commerce/entity"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type cartModel struct {
	DB *gorm.DB
}

func NewCartModel(db *gorm.DB) *cartModel {
	return &cartModel{
		DB: db,
	}
}

func (u *cartModel) Checkid(id uint, idcart uint) error {
	var cart entity.Cart
	u.DB.Where("id = ?", idcart).Find(&cart)

	if cart.UserID != id {
		return errors.New("error")
	}
	fmt.Println(cart.UserID)
	fmt.Println(id)

	return nil
}

func (u *cartModel) Insert(cart request.InsertCart, idUser uint) error {
	var product entity.Product
	if err := u.DB.Where("id = ?", cart.ProductID).Find(&product).Error; err != nil {
		return errors.New("Product Not Fount")
	}
	cartInsert := entity.Cart{
		Name:      product.Name,
		UserID:    idUser,
		Quantity:  cart.Quantity,
		Price:     product.Price,
		Image:     product.Image,
		ProductID: product.ID,
	}

	if err := u.DB.Create(&cartInsert).Error; err != nil {
		return err
	}
	return nil
}

func (u *cartModel) Update(cart request.UpdateCart, id uint) error {
	if err := u.DB.Model(&entity.Cart{}).Where("id = ?", id).Update("quantity", cart.Quantity).Error; err != nil {
		return err
	}
	return nil
}

func (u *cartModel) Delete(id uint) error {
	var cart entity.Cart
	if err := u.DB.Where("id = ?", id).Delete(&cart).Error; err != nil {
		log.Warn(err)
		return err
	}
	return nil
}
func (u *cartModel) GetAll(id uint) ([]entity.Cart, error) {
	var cart []entity.Cart
	if err := u.DB.Where("user_id=?", id).Find(&cart).Error; err != nil {
		log.Warn(err)
		return []entity.Cart{}, err
	}
	return cart, nil
}
