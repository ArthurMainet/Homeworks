package products

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" gorm:"uniqueIndex"`
	Description string         `json:"description" gorm:"uniqueIndex"`
	Images      pq.StringArray `json:"images" gorm:"type:text[],uniqueIndex"`
	Price       float64        `json:"price"`
}

func NewProduct(productReq *ProductRequest) *Product {
	return &Product{
		Name:        productReq.Name,
		Description: productReq.Description,
		Images:      productReq.Images,
		Price:       productReq.Price,
	}
}
