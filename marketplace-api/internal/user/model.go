package user

import (
	"Email-API/internal/order"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Phone    string `json:"phone" gorm:"uniqueIndex,notNull"`
	Role     string `json:"role" gorm:"notNull"`
	IsVerif  bool
	Order    []order.Order `json:"order" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
