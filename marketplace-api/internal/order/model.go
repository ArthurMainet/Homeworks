package order

import (
	"Email-API/internal/products"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   uint               `json:"user_id"`
	Products []products.Product `json:"products" gorm:"many2many:order_products,constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewOrder(products []int, userid uint) *Order {
	return &Order{
		UserID:   userid,
		Products: products,
	}
}
