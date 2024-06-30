package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeJSON(w http.ResponseWriter, status int, content any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(content)
}

func DecodeJSON(r *http.Request, content any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(content)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	EncodeJSON(w, status, map[string]string{"error": err.Error()})
}
