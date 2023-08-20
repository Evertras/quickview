package server

import (
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func New(address, filename string) *Server {
	mux := func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
	}

	s := &http.Server{
		Addr:    address,
		Handler: http.HandlerFunc(mux),
	}

	return &Server{
		server: s,
	}
}

func (s *Server) ListenAndServe() error {
	log.Printf("Listening at http://%s", s.server.Addr)

	return s.server.ListenAndServe()
}
