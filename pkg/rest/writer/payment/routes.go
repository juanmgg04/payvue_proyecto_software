package payment

import (
	"github.com/go-chi/chi/v5"
)

func (h *handler) RouteURLs(router *chi.Mux) {
	router.Route("/api/v1/payments", func(r chi.Router) {
		r.Post("/", h.CreatePayment)
		r.Delete("/{id}", h.DeletePayment)
	})
}
