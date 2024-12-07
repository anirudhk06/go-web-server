package user

import (
	"net/http"

	"github.com/anirudhk06/go-web-server/types"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) UserRoutes(routes *http.ServeMux) {
	routes.HandleFunc("POST /auth/register", h.HandleRegister)
	routes.HandleFunc("POST /auth/login", h.HandleLogin)
	routes.HandleFunc("GET /users", h.GetUsers)
}
