package verify

import (
	"Email-API/packages/api"
	"Email-API/packages/responce"
	"net/http"
)

type VerifyHandlerDeps struct {
	EmailService *EmailService
	PhoneService *PhoneService
}

type VerifyHandler struct {
	EmailService *EmailService
	PhoneService *PhoneService
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	service := &VerifyHandler{
		EmailService: deps.EmailService,
		PhoneService: deps.PhoneService,
	}
	//	router.HandleFunc("POST /send", mail.ReciveEmail())
	router.HandleFunc("GET /verify/{method}/{hash}", service.Verify())
}

func (e *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.PathValue("method")
		hash := r.PathValue("hash")

		if method == "email" {
			if _, ok := e.EmailService.Repo.AdrressAndHash[method][hash]; ok {
				responce.ResponceJSON(w, "You are welcome!", 200)
				e.EmailService.AprooveVerif(e.EmailService.Repo.AdrressAndHash[method][hash].AdrressOrSession)
			} else {
				responce.ResponceJSON(w, "Wrong register hash.", http.StatusUnauthorized)
			}
			e.EmailService.Repo.Delete(method, hash)
			e.EmailService.Repo.Save()
			return
		} else if method == "phone" {
			body, err := api.DecodeJSON[SessionVerifRequest](r.Body)
			if err != nil {
				http.Error(w, "Invaild data", 402)
				return
			}
			if e.PhoneService.Repo.AdrressAndHash[method][body.Code].AdrressOrSession == hash {
				responce.ResponceJSON(w, "You are welcome!", 200)
				token, err := e.PhoneService.AprooveVerif(e.PhoneService.Repo.AdrressAndHash[method][hash].AdrressOrSession)
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

			} else {
				responce.ResponceJSON(w, "Wrong code.", http.StatusUnauthorized)
			}
		}

	}
}
