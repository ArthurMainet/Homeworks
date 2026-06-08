package auth

import (
	"Email-API/packages/api"
	"Email-API/packages/responce"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type AuthHandlerDeps struct {
	AuthService *AuthService
}

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/login/phone", handler.LoginByPhone())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := api.HandleReq[RegisterRequest](w, r)
		if err != nil {
			log.Println(err)
			HandleErrors(w, err)
			return
		}

		err = handler.AuthService.Register(body.Email, body.Password, body.Phone)
		if err != nil {
			log.Println(err)
			HandleErrors(w, err)
			return
		}

		responce.ResponceJSON(w, "Succesful register.", http.StatusOK)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := api.HandleReq[LoginRequest](w, r)
		if err != nil {
			HandleErrors(w, err)
			return
		}

		user, err := handler.AuthService.Repo.FindByEmail(body.Email)
		if err != nil {
			http.Error(w, "user not found", http.StatusBadRequest)
			return
		}

		token, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			HandleErrors(w, err)
			return
		}
		userId := strconv.Itoa(int(user.ID))

		http.SetCookie(w, &http.Cookie{
			Name:     "userid",
			Value:    userId,
			MaxAge:   900,
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "accessToken",
			Value:    token,
			MaxAge:   900,
			HttpOnly: true,
		})

		responce.ResponceJSON(w,
			"Succeful login. We send mail to your email-address to get verif. Please check",
			http.StatusOK)

	}
}

func (handler *AuthHandler) LoginByPhone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := api.HandleReq[PhoneLoginRequest](w, r)
		if err != nil {
			HandleErrors(w, err)
			return
		}

		session, err := handler.AuthService.PhoneLogin(body.Phone)
		if err != nil {
			HandleErrors(w, err)
			return
		}

		responce.ResponceJSON(w, session, 200)

	}
}

func HandleErrors(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrUserNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, ErrInvalidPassword):
		http.Error(w, err.Error(), http.StatusUnauthorized)
	case errors.Is(err, ErrEmailVerified):
		http.Error(w, err.Error(), http.StatusConflict)
	case errors.Is(err, ErrUserAlreadyRegistered):
		http.Error(w, err.Error(), http.StatusConflict)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
