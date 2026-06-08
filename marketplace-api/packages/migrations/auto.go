package main

import (
	"Email-API/internal/order"
	"Email-API/internal/products"
	"Email-API/internal/user"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&products.Product{}, &user.User{}, &order.Order{})
	if err != nil {
		panic(err)
	}
}
