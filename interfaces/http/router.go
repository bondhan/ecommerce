package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func NewRouter(hndl *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Mount("/api/v1/", http.NotFoundHandler())
	return r
}
