package router

import (
	"net/http"

	"github.com/YacineMK/Rill/internal/handlers"
	customMiddleware "github.com/YacineMK/Rill/internal/middleware"
	"github.com/go-chi/chi/v5"
)

var streamHandler = &handlers.StreamHandler{}

func Route(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/create.html", http.StatusMovedPermanently)
	})

	r.Route("/api/v1", func(mux chi.Router) {
		mux.Use(customMiddleware.Json)
		mux.Get("/stream-key", streamHandler.StreamKeyHandler)
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Handle("/tmp/*", http.StripPrefix("/tmp/", http.FileServer(http.Dir("tmp"))))
}
