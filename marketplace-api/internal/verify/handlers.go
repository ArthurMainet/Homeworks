package verify

import (
	"Email-API/packages/api"
	"Email-API/packages/responce"
	"fmt"
	"net/http"
)

type VerifyHandlerDeps struct {
	EmailService *EmailService
	PhoneService *PhoneService
	Repo         *LocalRepo
}

type VerifyHandler struct {
	EmailService *EmailService
	PhoneService *PhoneService
	Repo         *LocalRepo
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	service := &VerifyHandler{
		EmailService: deps.EmailService,
		PhoneService: deps.PhoneService,
	}
	//	router.HandleFunc("POST /send", mail.ReciveEmail())
	router.HandleFunc("GET /verify/email/{hash}", service.EmailVerify())
	router.HandleFunc("POST /verify/phone", service.PhoneVerify())
}

func (e *VerifyHandler) EmailVerify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		if _, ok := e.EmailService.Repo.EmailAndHash[hash]; ok {
			responce.ResponceJSON(w, "You are welcome!", 200)
			fmt.Println(e.EmailService.Repo.EmailAndHash[hash].Email)
			e.EmailService.AprooveVerif(e.EmailService.Repo.EmailAndHash[hash].Email)
		} else {
			responce.ResponceJSON(w, "Wrong register hash.", http.StatusUnauthorized)
			return
		}

		e.EmailService.Repo.DeleteEmail(hash)
		e.EmailService.Repo.SaveEmailHash()
	}
}

func (e *VerifyHandler) PhoneVerify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := api.DecodeJSON[SessionVerifRequest](r.Body)
		if err != nil {
			http.Error(w, "Invaild data", 402)
			return
		}

		if session, ok := e.PhoneService.Repo.PhoneAndCode[body.Code]; ok {
			if e.PhoneService.Repo.PhoneAndCode[body.Code].Session == body.Session {
				fmt.Println(session)
				token, err := e.PhoneService.AprooveVerif(session.Phone)
				if err != nil {
					http.Error(w, "verif user error", 502)
					return
				}

				http.SetCookie(w, &http.Cookie{
					Name:     "accessToken",
					Value:    token,
					MaxAge:   21600,
					HttpOnly: true,
				})
				responce.ResponceJSON(w, "You are welcome!", 200)

			} else {
				responce.ResponceJSON(w, "Wrong code.", http.StatusUnauthorized)
			}
		}

		e.PhoneService.Repo.DeleteSession(body.Code)
		e.PhoneService.Repo.SaveSessionCode()
	}

}
