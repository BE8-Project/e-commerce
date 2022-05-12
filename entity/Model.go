package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(35);not null"`
	Username string    `gorm:"type:varchar(35);not null;unique"`
	Email    string    `gorm:"type:varchar(100);not null;unique"`
	HP       string    `gorm:"type:varchar(20);not null;unique"`
	Password string    `gorm:"type:varchar(255);not null"`
	Role     int       `gorm:"type:int;not null"`
	Products []Product `gorm:"foreignkey:UserID"`
	Addresses []Address `gorm:"foreignkey:UserID"`
	Orders    []Order   `gorm:"foreignkey:UserID"`
}

type Category struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);not null"`
	Slug     string    `gorm:"type:varchar(100);not null;unique"`
	Products []Product `gorm:"foreignkey:CategoryID"`
}

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null"`
	Slug        string  `gorm:"type:varchar(100);not null;unique"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Stock       uint    `gorm:"type:int;not null"`
	Description string  `gorm:"type:text;not null"`
	Image       string  `gorm:"type:varchar(100);not null"`
	UserID      uint    `gorm:"type:int;not null"`
	CategoryID  uint    `gorm:"type:int;not null"`
	Cart        []Cart  `gorm:"foreignkey:ProductID"`
}

type Cart struct {
	gorm.Model
	UserID    uint    `gorm:"type:int;not null"`
	Name      string  `gorm:"type:varchar(100);not null"`
	Quantity  uint    `gorm:"type:int;not null"`
	Price     float64 `gorm:"type:decimal(10,2);not null"`
	Image     string  `gorm:"type:varchar(100);not null"`
	ProductID uint    `gorm:"type:int;not null"`
}

type Address struct {
	gorm.Model
	UserID uint `gorm:"type:int;not null"`
	Address string `gorm:"type:varchar(100);not null"`
	City string `gorm:"type:varchar(100);not null"`
	Country string `gorm:"type:varchar(100);not null"`
	ZipCode int `gorm:"type:varchar(100);not null"`
	Orders []Order `gorm:"foreignkey:AddressID"`
}

type Order struct {
	gorm.Model
	UserID uint `gorm:"type:int;not null"`
	AddressID uint `gorm:"type:int;not null"`
	TrackingNumber string `gorm:"type:varchar(100);not null"`
	PaymentType string `gorm:"type:varchar(100);not null"`
	Total float64 `gorm:"type:decimal(10,2);not null"`
	Status string `gorm:"type:varchar(100);not null;default:'pending'"`
}