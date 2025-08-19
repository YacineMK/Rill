package server

import (
	"log"
	"net/http"

	customMiddleware "github.com/YacineMK/Rill/internal/middleware"
	"github.com/YacineMK/Rill/internal/router"
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
	s.Router.Use(customMiddleware.Json)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	router.Route(s.Router)
}

func (s *HttpServer) Start() error {
	s.Setup()
	log.Printf("📡 HTTP server listening on %v\n", s.Port)
	return http.ListenAndServe(":"+s.Port, s.Router)
}
