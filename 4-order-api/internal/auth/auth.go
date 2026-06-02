package auth

import (
	"Email-API/internal/user"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceDeps struct {
	Repo *user.UserRepository
}

type AuthService struct {
	Repo *user.UserRepository
}

func NewAuthService(deps *AuthServiceDeps) *AuthService {
	return &AuthService{
		Repo: deps.Repo,
	}
}

func (service *AuthService) Register(email, password, phone string) error {

	searchedUser, _ := service.Repo.FindByEmail(email)
	if searchedUser != nil {
		return errors.New("There is already a user with this email.")
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &user.UserModel{
		Email:    email,
		Password: string(hashedPass),
		Phone:    phone,
	}
	result := service.Repo.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *AuthService) Login(email, password string) error {

	searchedUser, _ := service.Repo.FindByEmail(email)
	if searchedUser == nil {
		return errors.New("Wrong email or password.")
	}
	err := bcrypt.CompareHashAndPassword([]byte(searchedUser.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return errors.New("Wrong email or password.")
	}
	return nil
}
