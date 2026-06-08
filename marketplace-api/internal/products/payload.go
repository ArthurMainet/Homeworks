package products

import "github.com/lib/pq"

type ProductRequest struct {
	Name        string         `json:"name" validate:"min=3,max=200"`
	Description string         `json:"description" validate:"min=3,max=1000"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Price       float64        `json:"price" validate:"gt=0"`
}

type UpdateProductRequest struct {
	Name        *string         `json:"name" validate:"omitempty,min=3,max=200"`
	Description *string         `json:"description" validate:"omitempty,min=3,max=1000"`
	Images      *pq.StringArray `json:"images" validate:"omitempty,dive,url" gorm:"type:text[]"`
	Price       *float64        `json:"price" validate:"omitempty,gt=0"`
}
