package packages

import (
	"Email-API/config"
	"net/http"
)

type EmailHandlerDeps struct {
	Config *config.Config
}

type EmailHandler struct {
	Email    string
	Password string
	Adress   string
}

func NewEmailHandler(e EmailHandlerDeps) *EmailHandler {
	return &EmailHandler{
		Email:    e.Config.Email,
		Password: e.Config.Password,
		Adress:   e.Config.Adress,
	}
}

func (e *EmailHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (e *EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
