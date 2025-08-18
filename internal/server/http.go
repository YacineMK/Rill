package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpServer struct {
	Port   string
	Router *chi.Mux
}

func NewHttpServer(port string) *HttpServer {
	return &HttpServer{
		Port:   port,
		Router: chi.NewRouter(),
	}
}

func (s *HttpServer) Setup() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	s.Router.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to API"))
	})
}

func (s *HttpServer) Start() error {
	s.Setup()
	log.Printf("📡 HTTP server listening on %v\n", s.Port)
	return http.ListenAndServe(":"+s.Port, s.Router)
}
