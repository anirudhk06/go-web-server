package user

import (
	"net/http"

	"github.com/anirudhk06/go-web-server/service/auth"
	"github.com/anirudhk06/go-web-server/types"
	"github.com/anirudhk06/go-web-server/utils"
)

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterPayload
	utils.ParseJSON(r, &payload)

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, map[string]any{"email": "Email already exists."})
		return
	}

	password, err := auth.HashPassword(payload.Password)

	if err != nil {
		utils.InternalServerError(w)
		return
	}

	err = h.store.CreateUser(types.User{
		Email:    payload.Email,
		Password: password,
	})

	if err != nil {
		utils.InternalServerError(w)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "Registration successfully completed.")

}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginPayload

	utils.ParseJSON(r, &payload)

	err := payload.Validate()

	if len(err) > 0 {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.store.GetUserByEmail("anirudh")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, map[string]string{"error": "Invalidm credentials."})
		return
	}

	err = auth.ValidatePassword("", "")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, map[string]string{"error": "Invalidm credentials."})
		return

	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{"access": "", "refresh": ""})

}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := h.store.FindUsers()

	utils.WriteJSON(w, http.StatusOK, users)

}
