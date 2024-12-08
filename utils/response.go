package utils

import (
	"encoding/json"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)

}

func WriteError(w http.ResponseWriter, statusCode int, err any) error {
	return WriteJSON(w, statusCode, err)
}

func InternalServerError(w http.ResponseWriter) error {
	return WriteJSON(w, http.StatusInternalServerError, map[string]any{"detail": "Internal server error"})
}

func Unauthorized(w http.ResponseWriter) error {
	return WriteJSON(w, http.StatusUnauthorized, map[string]any{
		"detail": "Unauthorized",
	})
}
