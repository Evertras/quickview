package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func New(address, filename string) *Server {
	index := handlerIndex(indexData{
		filename:     filename,
		websocketURL: fmt.Sprintf("ws://%s/watch", address),
	})
	watcher := handlerWatcher(filename)

	mux := func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		log.Println(p)

		switch p {
		case "/":
			index(w, r)

		case "/item":
			log.Println("Serving", filename)
			http.ServeFile(w, r, filename)

		case "/favicon.ico":
			// Figure out favicon later

		case "/watch":
			watcher.ServeHTTP(w, r)

		default:
			log.Println("ERROR: Unknown route")
		}
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
