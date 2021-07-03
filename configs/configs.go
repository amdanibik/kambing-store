package configs

import (
	"app/models/category"
	"app/models/methods"
	"app/models/product"
	"app/models/rekening"
	"fmt"
	"os/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigDB struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_POST string
	DB_DTBS string
}

func InitDB() {
	config := ConfigDB{
		DB_USER: "root",
		DB_PASS: "1234Abcd",
		DB_HOST: "127.0.0.1",
		DB_POST: "3306",
		DB_DTBS: "terra",
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_POST, config.DB_DTBS)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	AutoMigrate()
}

func AutoMigrate() {
	DB.AutoMigrate(&user.User{})
	DB.AutoMigrate(&category.Category{})
	DB.AutoMigrate(&product.Product{})
	DB.AutoMigrate(&methods.Methods{})
	DB.AutoMigrate(&rekening.Rekening{})
}
