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
		rc.Put("/{id}", hnd.cashier.Update)
		rc.Delete("/{id}", hnd.cashier.Delete)
		rc.Get("/", hnd.cashier.List)
		rc.Get("/{id}", hnd.cashier.Detail)
	})

	r.Route("/categories", func(rc chi.Router) {
		rc.Post("/", hnd.category.Create)
		rc.Put("/{id}", hnd.category.Update)
		rc.Delete("/{id}", hnd.category.Delete)
		rc.Get("/", hnd.category.List)
		rc.Get("/{id}", hnd.category.Detail)
	})

	r.Route("/payments", func(rc chi.Router) {
		rc.Post("/", hnd.payment.Create)
		rc.Put("/{id}", hnd.payment.Update)
		rc.Delete("/{id}", hnd.payment.Delete)
		rc.Get("/", hnd.payment.List)
		rc.Get("/{id}", hnd.payment.Detail)
	})

	r.Route("/products", func(rc chi.Router) {
		rc.Post("/", hnd.product.Create)
		rc.Put("/{id}", hnd.product.Update)
		rc.Delete("/{id}", hnd.product.Delete)
		rc.Get("/", hnd.product.List)
		rc.Get("/{id}", hnd.product.Detail)
	})

	return r
}
