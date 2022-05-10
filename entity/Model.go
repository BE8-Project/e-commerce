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

type Category struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Slug     string `gorm:"type:varchar(100);not null;unique"`
	Products []Product `gorm:"foreignkey:CategoryID"`
}

type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"`
	Slug        string `gorm:"type:varchar(100);not null;unique"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Stock 	    uint     `gorm:"type:int;not null"`
	Description string `gorm:"type:text;not null"`
	Image       string `gorm:"type:varchar(100);not null"`
	UserID 		uint   `gorm:"type:int;not null"`
	CategoryID  uint  `gorm:"type:int;not null"`
}