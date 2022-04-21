package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func NewRouter(logger *logrus.Logger, hnd *Handler, jwtKey string) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	//r.Use(middleware.Logger)
	//r.Use(requestlogmsv2.CustomLoggerV2(logger, nil))

	r.Route("/cashiers", func(rc chi.Router) {
		rc.Post("/", hnd.cashier.Create)
		rc.Put("/{id}", hnd.cashier.Update)
		rc.Delete("/{id}", hnd.cashier.Delete)
		rc.Get("/", hnd.cashier.List)
		rc.Get("/{id}", hnd.cashier.Detail)
		rc.Get("/{id}/passcode", hnd.cashier.PassCode)
		rc.Post("/{id}/login", hnd.cashier.Login)
		rc.Post("/{id}/logout", hnd.cashier.Logout)
	})

	r.Route("/categories", func(rc chi.Router) {
		rc.Post("/", hnd.category.Create)
		rc.Put("/{id}", hnd.category.Update)
		rc.Delete("/{id}", hnd.category.Delete)
		rc.Get("/", JWTValidator(jwtKey, hnd.category.List))
		rc.Get("/{id}", JWTValidator(jwtKey, hnd.category.Detail))
	})

	r.Route("/payments", func(rc chi.Router) {
		rc.Post("/", hnd.payment.Create)
		rc.Put("/{id}", hnd.payment.Update)
		rc.Delete("/{id}", hnd.payment.Delete)
		rc.Get("/", JWTValidator(jwtKey, hnd.payment.List))
		rc.Get("/{id}", JWTValidator(jwtKey, hnd.payment.Detail))
	})

	r.Route("/products", func(rc chi.Router) {
		rc.Post("/", hnd.product.Create)
		rc.Put("/{id}", hnd.product.Update)
		rc.Delete("/{id}", hnd.product.Delete)
		rc.Get("/", JWTValidator(jwtKey, hnd.product.List))
		rc.Get("/{id}", JWTValidator(jwtKey, hnd.product.Detail))
	})

	r.Route("/orders", func(rc chi.Router) {
		rc.Post("/subtotal", JWTValidator(jwtKey, hnd.order.SubTotal))
		rc.Post("/", JWTValidator(jwtKey, hnd.order.Create))
		rc.Get("/", hnd.order.List)
		rc.Get("/{id}", JWTValidator(jwtKey, hnd.order.Detail))
		rc.Get("/{id}/download", JWTValidator(jwtKey, hnd.order.Download))
		rc.Get("/{id}/check-download", JWTValidator(jwtKey, hnd.order.DownloadStatus))
	})

	r.Get("/revenues", JWTValidator(jwtKey, hnd.order.Revenues))
	r.Get("/solds", JWTValidator(jwtKey, hnd.order.Solds))

	return r
}
