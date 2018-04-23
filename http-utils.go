package main

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil && data != "" {
		return json.NewEncoder(w).Encode(data)
	}

	return nil
}
