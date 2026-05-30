package internal

import (
	"Email-API/config"
	"Email-API/packages"
	"log"
	"net/http"
	"net/smtp"

	"github.com/go-playground/validator/v10"
	"github.com/jordan-wright/email"
)

type EmailHandlerDeps struct {
	Config *config.Config
	Repo   *LocalRepo
}

type EmailHandler struct {
	Email    string
	Password string
	Adress   string
	Repo     *LocalRepo
}

func NewEmailHandler(e EmailHandlerDeps) *EmailHandler {
	return &EmailHandler{
		Email:    e.Config.Email,
		Password: e.Config.Password,
		Adress:   e.Config.Address,
		Repo:     e.Repo,
	}
}

func (e *EmailHandler) ReciveEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := packages.DecodeJSON[EmailRequest](r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		valid := validator.New()
		err = valid.Struct(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Закидываю емеил в функцию, которая генерит хэш и сует все это в мапу.
		model := NewEmailWithHash(body.Email)
		mail := model.Email
		hash := model.Hash
		e.Repo.EmailAndHash[hash] = model
		text := "http://localhost:8081/verify/" + hash

		err = e.WritingEmail(mail, text)
		if err != nil {
			packages.ResponceJSON(w, "Coudn't send mail. Retry please.", http.StatusInternalServerError)
			return
		}
		packages.ResponceJSON(w, "Mail with registration-URL already send to your email address", http.StatusOK)

		err = e.Repo.Save()
		if err != nil {
			log.Println("Email and hash didn't save. Reason: ", err)
		}

	}
}

func (em *EmailHandler) WritingEmail(mail, text string) error {
	e := email.NewEmail()
	e.From = "BabyMelo <testbabymail@mail.ru>"
	e.To = []string{mail}
	e.Text = []byte(text)
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", em.Email, em.Password, em.Adress))
	return err
}

func (e *EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if _, ok := e.Repo.EmailAndHash[hash]; ok {
			packages.ResponceJSON(w, "You are welcome!", 200)
		} else {
			packages.ResponceJSON(w, "Wrong register hash.", http.StatusUnauthorized)
		}
		e.Repo.Delete(hash)
		e.Repo.Save()
	}
}
