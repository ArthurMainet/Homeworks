package user

import (
	"Email-API/packages/db"
	"fmt"
	"log"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	DB *db.DB
}

func NewUserRepository(db *db.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) FindByEmail(email string) (*UserModel, error) {
	var user UserModel
	result := repo.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// func (repo *UserRepository) FindByPhone(phone string) (bool, error) {
// 	var user UserModel
// 	result := repo.DB.First(&user, "phone = ?", phone)
// 	if result.Error != nil {
// 		return false, result.Error
// 	}
// 	return true, nil
// }

func (repo *UserRepository) Create(user *UserModel) error {
	result := repo.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *UserRepository) Update(user *UserModel) (*UserModel, error) {
	result := pr.DB.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		log.Println("Can't update: ", result.Error)
		return nil, result.Error
	}
	pr.DB.Save(&user)
	return user, nil
}

func (pr *UserRepository) GetByAddress(method, address string) (*UserModel, error) {
	var product UserModel
	result := pr.DB.First(&product, fmt.Sprintf("%s = ?", method), address)
	if result.Error != nil {
		log.Println("Can't find by Address: ", result.Error)
		return nil, result.Error
	}
	return &product, nil
}
