package repository

import (
	"e-commerce/delivery/helpers/response"
	"e-commerce/entity"
)

type UserModel interface {
	Insert(user *entity.User) (response.User, error)
	Login(custom []string, password string) (response.Login, error)
	GetOne(username string) response.User
	Delete(username string) response.DeleteUser
	Update(newUser *entity.User, username string) (response.UpdateUser, error)
}
