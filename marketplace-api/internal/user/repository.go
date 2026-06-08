package user

import (
	"Email-API/internal/auth"
	"Email-API/packages/db"
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

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, auth.ErrUserNotFound
	}
	return &user, nil
}

func (repo *UserRepository) FindByPhone(phone string) (*User, error) {
	var user User
	result := repo.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) Create(user *User) error {
	result := repo.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *UserRepository) Update(user *User) (*User, error) {
	result := pr.DB.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		log.Println("Can't update: ", result.Error)
		return nil, result.Error
	}
	pr.DB.Save(&user)
	return user, nil
}
