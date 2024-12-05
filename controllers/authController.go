package controllers

import (
	"net/http"
	"time"

	"github.com/anirudhk06/go-web-server/configs"
	"github.com/anirudhk06/go-web-server/models"
	"github.com/anirudhk06/go-web-server/utils"
	"github.com/golang-jwt/jwt/v5"
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

	var payload struct {
		Email    string `json:"email,required"`
		Password string `json:"password"`
	}

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteJSON(w, map[string]string{"detail": "Invalid json data"}, http.StatusBadRequest)
		return
	}

	user := models.User{}
	err := configs.DB.Where("email = ?", payload.Email).First(&user).Error

	if err != nil {
		utils.WriteJSON(w, map[string]string{"detail": "Invalid credentials"}, http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		utils.WriteJSON(w, map[string]string{"detail": "Invalid credentials"}, http.StatusBadRequest)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Second * 10).Unix(),
		"iat": time.Now().Unix(),
	})

	utils.WriteJSON(w, map[string]string{
		"access": "",
		"refrsh": "",
	}, http.StatusOK)
}
