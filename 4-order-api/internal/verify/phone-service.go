package verify

import (
	"Email-API/internal/user"
	"Email-API/packages/jwt"
	"errors"
	"log"
)

// type PhoneSender interface {
// 	SendCode(address string) (string, error)
// }

type PhoneServiceDeps struct {
	Repo           *LocalRepo
	UserRepository *user.UserRepository
	JWT            *jwt.JWT
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
		JWT:            deps.JWT,
	}
}

func (ps *PhoneService) SendCode(session, address string) (string, error) {
	data := NewEmailWithHash("phone", session)
	repo := NewLocalRepo()

	err := sendingCode(data.Hash, address)
	if err != nil {
		return "", errors.New("Send code error")
	}
	repo.AdrressAndHash["phone"][data.Hash] = data

	return data.Hash, nil
}

func sendingCode(code, address string) error {
	// здесь фукнция которая типа отсылает код "code"
	// возвращаем nil потому что все прошло успешно
	return nil
}

func (ps *PhoneService) AprooveVerif(address string) (string, error) {
	user, err := ps.UserRepository.GetByAddress("phone", address)
	if err != nil {
		log.Printf("Error to find of user with addrees %s. Reason: %v", user.Phone, err)
		return "", nil
	}
	user.IsVerif = true
	_, err = ps.UserRepository.Update(user)
	if err != nil {
		log.Printf("Error to update verif status of user with address %s. Reason: %v", user.Phone, err)
		return "", nil
	}

	token, err := ps.JWT.GenerateToken(user.Email, user.Role,
		user.Phone, ps.JWT.Secret)
	if err != nil {
		return "", err
	}

	log.Printf("Succesful verif. User - %s", user.Phone)
	return token, nil
}
