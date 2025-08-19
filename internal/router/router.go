package router

import (
	"github.com/YacineMK/Rill/internal/handlers"
	"github.com/go-chi/chi/v5"
)

var streamHandler = &handlers.StreamHandler{}

func Route(r *chi.Mux) {
	r.Route("/api/v1", func(mux chi.Router) {
		mux.Get("/stream-key", streamHandler.StreamKeyHandler)
	})
}