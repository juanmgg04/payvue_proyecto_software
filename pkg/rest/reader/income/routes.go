package income

import (
	"github.com/go-chi/chi/v5"
)

func (h *handler) RouteURLs(router *chi.Mux) {
	router.Route("/api/v1/incomes", func(r chi.Router) {
		r.Get("/", h.GetAllIncomes)
		r.Get("/{id}", h.GetIncomeByID)
	})
}
