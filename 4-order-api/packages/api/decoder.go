package api

import (
	"encoding/json"
	"io"
	"log"
)

func DecodeJSON[T any](body io.Reader) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		log.Println(err)
		return payload, err
	}
	return payload, nil
}
