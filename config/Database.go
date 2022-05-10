package config

import (
	"e-commerce/entity"
	"fmt"

	"github.com/labstack/gommon/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config AppConfig) *gorm.DB {
	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
	)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to database", err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Product{})

	// db.Create(&entity.User{
	// 	Name: "Admin Website",
	// 	Username: "admin",
	// 	HP: "081234567892",
	// 	Email: "admin@gmail.com",
	// 	Password: "$2a$14$fqChZ4CqMd9uvfE7MU6y4OvjTsBHoIBSbN/Iymyu9fCBJ9/VoCXum", // password
	// 	Role: 2,
	// })

	// db.Create(&entity.User{
	// 	Name: "Merchant Website",
	// 	Username: "merchant",
	// 	HP: "081234567891",
	// 	Email: "merchant@gmail.com",
	// 	Password: "$2a$14$fqChZ4CqMd9uvfE7MU6y4OvjTsBHoIBSbN/Iymyu9fCBJ9/VoCXum", // password
	// 	Role: 1,
	// })

	// db.Create(&entity.User{
	// 	Name: "User Website",
	// 	Username: "user",
	// 	HP: "081234567890",
	// 	Email: "user@gmail.com",
	// 	Password: "$2a$14$fqChZ4CqMd9uvfE7MU6y4OvjTsBHoIBSbN/Iymyu9fCBJ9/VoCXum", // password
	// 	Role: 0,
	// })
}
