package routes

import (
	"net/http"

	"github.com/anirudhk06/go-web-server/controllers"
)

func AuthRoutes() http.Handler {
	authMux := http.NewServeMux()
	authMux.HandleFunc("POST /register", controllers.Register)
	authMux.HandleFunc("POST /login", controllers.Login)
	return authMux
}
