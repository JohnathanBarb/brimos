package main


import (
	"net/http"
)

type Server struct {}

func NewServer() *Server {
	return &Server{}
}


func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/api/persons", s.createPersonHandler)

	return mux
}
