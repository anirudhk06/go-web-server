package controllers

import (
	"net/http"

	"github.com/anirudhk06/go-web-server/configs"
	"github.com/anirudhk06/go-web-server/models"
	"github.com/anirudhk06/go-web-server/utils"
	"golang.org/x/crypto/bcrypt"
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

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)

	if err != nil {
		utils.WriteJSON(w, map[string]string{"detail": "Something went wrong"}, http.StatusInternalServerError)
		return
	}

	user := models.User{
		Email:    payload.Email,
		Password: string(hash),
	}

	configs.DB.Create(&user)

	utils.WriteJSON(w, user, http.StatusCreated)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {

}
