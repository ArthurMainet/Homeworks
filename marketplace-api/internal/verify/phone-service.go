package verify

import (
	"Email-API/config"
	"Email-API/internal/user"
	"Email-API/packages/jwt"
	"errors"
	"fmt"
	"log"
)

// type PhoneSender interface {
// 	SendCode(address string) (string, error)
// }

type PhoneServiceDeps struct {
	Repo           *LocalRepo
	UserRepository *user.UserRepository
	JWT            *config.AuthConfig
}

type PhoneService struct {
	Repo           *LocalRepo
	UserRepository *user.UserRepository
	JWT            *jwt.JWT
}

func NewPhoneService(deps *PhoneServiceDeps) *PhoneService {
	return &PhoneService{
		Repo:           deps.Repo,
		UserRepository: deps.UserRepository,
		JWT: &jwt.JWT{
			Secret: deps.JWT.AuthToken,
		},
	}
}

func (ps *PhoneService) SendCode(session, address string) (string, error) {
	data := NewSessionWithCode(session, address)
	fmt.Println(data)

	log.Printf("!! CODE HERE %d CODE HERE !!\n", data.Code)

	err := sendingCode(data.Code, address)
	if err != nil {
		return "", errors.New("Send code error")
	}

	ps.Repo.PhoneAndCode[data.Code] = data

	return data.Session, nil
}

func sendingCode(code int, address string) error {
	// здесь фукнция которая типа отсылает код "code"
	// возвращаем nil потому что все прошло успешно
	return nil
}

func (ps *PhoneService) AprooveVerif(address string) (string, error) {
	user, err := ps.UserRepository.FindByPhone(address)
	if err != nil {
		log.Printf("Error to find of user with address %s. Reason: %v", address, err)
		return "", err
	}
	user.IsVerif = true
	_, err = ps.UserRepository.Update(user)
	if err != nil {
		log.Printf("Error to update verif status of user with address %s. Reason: %v", address, err)
		return "", err
	}

	token, err := ps.JWT.GenerateToken(user.ID, user.Email, user.Role,
		user.Phone, ps.JWT.Secret)
	if err != nil {
		return "", err
	}

	log.Printf("Succesful verif. User - %s", address)
	return token, nil
}
