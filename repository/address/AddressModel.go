package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"

	"gorm.io/gorm"
)

type addressModel struct {
	DB *gorm.DB
}

func NewAddressModel(db *gorm.DB) *addressModel {
	return &addressModel{
		DB: db,
	}
}

func (am *addressModel) Insert(address *entity.Address) (response.InsertAddress, error) {
	record := am.DB.Create(address)

	if record.RowsAffected == 0 {
		return response.InsertAddress{}, record.Error
	} else {
		return response.InsertAddress{
			UserID: address.UserID,
			CreatedAt: address.CreatedAt,
		}, nil
	}
}

func (am *addressModel) GetByUserID(userID uint) []response.Address {
	var addresses []response.Address

	record := am.DB.Where("user_id = ?", userID).Find(&addresses)

	if record.RowsAffected == 0 {
		return []response.Address{}
	} else {
		return addresses
	}
}