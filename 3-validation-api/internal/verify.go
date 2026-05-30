package internal

import (
	"Email-API/config"
	"Email-API/packages"
	"net/http"
	"net/smtp"

	"github.com/go-playground/validator/v10"
	"github.com/jordan-wright/email"
)

type EmailHandlerDeps struct {
	Config *config.Config
}

type EmailHandler struct {
	Email       string
	Password    string
	Adress      string
	hashedEmail *EmailWithHash
}

func NewEmailHandler(e EmailHandlerDeps) *EmailHandler {
	return &EmailHandler{
		Email:    e.Config.Email,
		Password: e.Config.Password,
		Adress:   e.Config.Adress,
	}
}

func (e *EmailHandler) ReciveEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := packages.DecodeJSON[EmailRequest](r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		valid := validator.New()
		err = valid.Struct(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		hashedEmail := NewEmailWithHash(body.Email)
		e.hashedEmail = hashedEmail
		mail := e.hashedEmail.Email
		hash := e.hashedEmail.Hash
		text := "http://localhost:8081/verify/" + hash
		e.WritingEmail(mail, text)

		// сохраняем email с хешем, записывая ее в буфер, откуда потом прочитаем
		// var buf bytes.Buffer
		// err = json.NewEncoder(&buf).Encode(hashedEmail)
		// if err != nil {
		// 	log.Println(err)
		// }

	}
}

func (em *EmailHandler) WritingEmail(mail, text string) {
	e := email.NewEmail()
	e.From = "BabyMelo <testbabymail@mail.ru>"
	e.To = []string{mail}
	e.Text = []byte(text)
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", em.Email, em.Password, em.Adress))
}

func (e *EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if hash == e.hashedEmail.Hash {
			packages.ResponceJSON(w, "You are welcome!", 200)
		}
	}
}
