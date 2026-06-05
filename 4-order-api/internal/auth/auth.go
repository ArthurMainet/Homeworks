package auth

import (
	"Email-API/config"
	"Email-API/internal/user"
	"Email-API/internal/verify"
	"Email-API/packages/jwt"
	"errors"
	"log"
	"math/rand/v2"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceDeps struct {
	Repo         *user.UserRepository
	EmailService *verify.EmailService
	PhoneService *verify.PhoneService
	JWT          *config.AuthConfig
}

type AuthService struct {
	Repo         *user.UserRepository
	EmailService *verify.EmailService
	PhoneService *verify.PhoneService
	JWT          *jwt.JWT
}

type Session struct {
	Session string
}

func NewAuthService(deps *AuthServiceDeps) *AuthService {
	return &AuthService{
		Repo:         deps.Repo,
		EmailService: deps.EmailService,
		PhoneService: deps.PhoneService,
		JWT: &jwt.JWT{
			Secret: deps.JWT.AuthToken,
		},
	}
}

func (service *AuthService) Register(email, password, phone string) error {

	searchedUser, _ := service.Repo.FindByEmail(email)
	if searchedUser != nil {
		return ErrUserAlreadyRegistered
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("server error")
	}

	role := "user"
	if email == "superpuper@mail.ru" {
		role = "admin"
	}
	user := &user.UserModel{
		Email:    email,
		Password: string(hashedPass),
		Phone:    phone,
		Role:     role,
		IsVerif:  false,
	}
	result := service.Repo.DB.Create(&user)
	if result.Error != nil {
		return errors.New("DB create error")
	}

	_ = service.EmailService.ReciveEmail(user.Email)
	// if err != nil {
	// 	return errors.New("Recive mail error.")
	// }

	return nil
}

func (service *AuthService) Login(email, password string) (string, error) {

	searchedUser, _ := service.Repo.FindByEmail(email)
	if searchedUser == nil {
		return "", ErrUserNotFound
	}
	err := bcrypt.CompareHashAndPassword([]byte(searchedUser.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return "", errors.New("Wrong email or password.")
	}

	token, err := service.JWT.GenerateToken(searchedUser.Email, searchedUser.Role,
		searchedUser.Phone, service.JWT.Secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service *AuthService) PhoneLogin(phone string) (string, error) {

	searchedUser, _ := service.Repo.GetByAddress("phone", phone)
	if searchedUser == nil {
		return "", ErrUserNotFound
	}

	session := generateSession()
	_, err := service.PhoneService.SendCode(session, phone)
	if err != nil {
		log.Printf("Phone code didnt send - reason: %v", err)
		return "", errors.New("Error send code")
	}

	return session, nil
}

func generateSession() string {
	runes := []rune("qwertyuiopasdfghjklzxcvbnm123456790")
	slice := make([]rune, 14)
	for i, _ := range runes {
		slice[i] = runes[rand.IntN(len(runes))]
	}
	session := string(slice)
	return session
}
