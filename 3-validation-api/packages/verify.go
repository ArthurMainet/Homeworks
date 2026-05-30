package packages

import "net/http"

type EmailHandlerDeps struct {
	Email    string
	Password string
	Adress   string
}

type EmailHandler struct {
	Email    string
	Password string
	Adress   string
}

func NewEmailHandler(e EmailHandlerDeps) *EmailHandler {
	return &EmailHandler{
		Email:    e.Email,
		Password: e.Password,
		Adress:   e.Adress,
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
