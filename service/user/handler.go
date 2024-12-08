package user

import (
	"net/http"
	"time"

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

	errors := payload.Validate()

	if len(errors) > 0 {
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	user, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, map[string]string{"error": "Invalid credentials."})
		return
	}

	isValid := auth.ValidatePassword(user.Password, []byte(payload.Password))

	if !isValid {
		utils.WriteError(w, http.StatusBadRequest, map[string]string{"error": "Invalid credentials."})
		return

	}

	token, err := auth.CreateJWT(user.ID)

	if err != nil {
		utils.InternalServerError(w)
		return
	}

	tokenCookie := http.Cookie{
		Name:     "access",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	http.SetCookie(w, &tokenCookie)

	utils.WriteJSON(w, http.StatusOK, map[string]any{"access": token, "refresh": ""})

}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	w.Write([]byte("success"))
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := h.store.FindUsers()

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"users": users,
	})

}
