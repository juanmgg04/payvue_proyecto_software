package debt

import (
	"github.com/go-chi/chi/v5"
)

func (h *handler) RouteURLs(router *chi.Mux) {
	router.Route("/api/v1/debts", func(r chi.Router) {
		r.Get("/", h.GetAllDebts)
		r.Get("/{id}", h.GetDebtByID)
	})
}
