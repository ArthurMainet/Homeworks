package verify

import (
	"Email-API/config"
	"Email-API/internal/user"
	"errors"
	"fmt"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// type EmailSender interface {
// 	ReciveEmail(email string) error
// 	WritingEmail(mail, text string) error
// 	AproveVerif(email string)
// }

type EmailServiceDeps struct {
	Repo           *LocalRepo
	EmailConf      *config.EmailConfig
	UserRepository *user.UserRepository
}

type EmailService struct {
	Repo           *LocalRepo
	EmailConf      *config.EmailConfig
	UserRepository *user.UserRepository
}

func NewEmailService(deps *EmailServiceDeps) *EmailService {
	return &EmailService{
		Repo:           deps.Repo,
		EmailConf:      deps.EmailConf,
		UserRepository: deps.UserRepository,
	}
}

func (e *EmailService) ReciveEmail(email string) error {

	// Закидываю емеил в функцию, которая генерит хэш и сует все это в мапу.
	model := NewEmailWithHash(email)
	mail := model.Email
	hash := model.Hash
	e.Repo.EmailAndHash[hash] = model
	text := "http://localhost:8081/verify/" + hash
	fmt.Println(hash)

	err := e.WritingEmail(mail, text)
	if err != nil {
		return errors.New("Send message error.")
	}

	err = e.Repo.SaveEmailHash()
	if err != nil {
		log.Println("Email and hash didn't save. Reason: ", err)
	}

	return nil
}

func (em *EmailService) WritingEmail(mail, text string) error {
	e := email.NewEmail()
	e.From = "BabyMelo <testbabymail@mail.ru>"
	e.To = []string{mail}
	e.Text = []byte(text)
	err := e.Send(mail, smtp.PlainAuth("", em.EmailConf.Email, em.EmailConf.Password, em.EmailConf.Address))
	return err
}

func (em *EmailService) AprooveVerif(address string) {
	user, err := em.UserRepository.FindByEmail(address)
	if err != nil {
		log.Printf("Error to find of user with addrees %s. Reason: %v", user.Email, err)
		return
	}
	user.IsVerif = true
	_, err = em.UserRepository.Update(user)
	if err != nil {
		log.Printf("Error to update verif status of user with address %s. Reason: %v", user.Email, err)
		return
	}
	log.Printf("Succesful verif. User - %s", user.Email)
}
