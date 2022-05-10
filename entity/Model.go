package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(35);not null"`
	Username string    `gorm:"type:varchar(35);not null;unique"`
	Email    string    `gorm:"type:varchar(100);not null;unique"`
	HP       string    `gorm:"type:varchar(20);not null;unique"`
	Password string    `gorm:"type:varchar(100);not null"`
	Role     int       `gorm:"type:int;not null"`
	Products []Product `gorm:"foreignkey:UserID"`
}
type Product struct {
	UserID uint `gorm:"type:int;not null"`
}
