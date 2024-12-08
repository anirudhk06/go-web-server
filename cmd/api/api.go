package api

import (
	"fmt"
	"net/http"

	"github.com/anirudhk06/go-web-server/middleware"
	"github.com/anirudhk06/go-web-server/service/user"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	userStore := user.NewStore(s.db)
	userRoutes := user.NewHandler(userStore)
	userRoutes.UserRoutes(v1)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))
	fmt.Printf("Server is running on port: %s\n", s.addr)

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":" + s.addr,
		Handler: middlewareStack(mux),
	}
	return server.ListenAndServe()
}
