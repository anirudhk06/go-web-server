package controllers

import (
	"net/http"

	"github.com/anirudhk06/go-web-server/models"
	"github.com/anirudhk06/go-web-server/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	payload := models.RegisterPayload{}

	err := utils.ParseJSON(r, &payload)

	if err != nil {
		utils.WriteJSON(w, map[string]string{"detail": "Invalid json data"}, http.StatusBadRequest)
		return
	}

	result := payload.Validate()

	if len(result) != 0 {
		utils.WriteJSON(w, result, http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, map[string]string{"status": "success"}, http.StatusCreated)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {

}
