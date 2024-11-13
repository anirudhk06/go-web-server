package api

import (
	"fmt"
	"net/http"
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

	fmt.Printf("Server is running on port: %s\n", s.addr)
	return http.ListenAndServe(":"+s.addr, mux)
}
