package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(hnd *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Route("/cashiers", func(rc chi.Router) {
		rc.Post("/", hnd.cashier.Create)
		rc.Get("/", hnd.cashier.List)
		rc.Get("/{id}", hnd.cashier.Detail)
		rc.Put("/{id}", hnd.cashier.Update)
		rc.Delete("/{id}", hnd.cashier.Delete)
	})

	return r
}
