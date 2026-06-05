package api

import (
	"Email-API/packages/responce"
	"net/http"
)

func HandleReq[T any](w http.ResponseWriter, req *http.Request) (*T, error) {
	body, err := DecodeJSON[T](req.Body)
	if err != nil {
		responce.ResponceJSON(w, err.Error(), 402)
		return nil, err
	}
	err = Validate(body)
	if err != nil {
		responce.ResponceJSON(w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}
