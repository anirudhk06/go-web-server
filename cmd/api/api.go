package api

import (
	"fmt"
	"net/http"

	"github.com/anirudhk06/go-web-server/routes"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	v1.Handle("/auth/", http.StripPrefix("/auth", routes.AuthRoutes()))

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	fmt.Printf("Server is running on port: %s\n", s.addr)
	return http.ListenAndServe(":"+s.addr, mux)
}
