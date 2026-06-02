package auth

import (
	"Email-API/internal/user"
	"Email-API/packages"
	"net/http"
)

type AuthHandlerDeps struct {
	Repo        *user.UserRepository
	AuthService *AuthService
}

type AuthHandler struct {
	Repo        *user.UserRepository
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		Repo:        deps.Repo,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /register", handler.Register())
	router.HandleFunc("POST /login", handler.Login())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := packages.DecodeJSON[RegisterRequest](r.Body)
		if err != nil {
			http.Error(w, "Invdalid request.", http.StatusBadRequest)
			return
		}

		err = handler.AuthService.Register(body.Email, body.Password, body.Phone)
		if err != nil {
			http.Error(w, "There is already a user with this email.", http.StatusBadRequest)
			return
		}

		packages.ResponceJSON(w, "Succesful register.", http.StatusOK)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := packages.DecodeJSON[LoginRequest](r.Body)
		if err != nil {
			http.Error(w, "Invdalid request.", http.StatusUnauthorized)
			return
		}

		err = handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, "Wrong email or password.", http.StatusUnauthorized)
			return
		}

		packages.ResponceJSON(w, "Succeful login.", http.StatusOK)

	}
}
