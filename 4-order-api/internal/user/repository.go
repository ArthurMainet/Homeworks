package user

import "Email-API/packages/db"

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

func (repo *UserRepository) Create(user *UserModel) error {
	result := repo.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
