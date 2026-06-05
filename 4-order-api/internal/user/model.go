package user

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Phone    string `json:"phone" gorm:"uniqueIndex,notNull"`
	Role     string `json:"role" gorm:"notNull"`
	IsVerif  bool
}
