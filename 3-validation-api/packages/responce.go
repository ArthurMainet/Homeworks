package packages

import (
	"encoding/json"
	"net/http"
)

func ResponceJSON(w http.ResponseWriter, data any, status int) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	return err
}
