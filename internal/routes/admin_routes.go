package routes

import (
	"top1affiliate/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterAdminRoutes(r chi.Router, handler *handlers.AdminHandler) {

	r.Route("/admin", func(r chi.Router) {
		r.Post("/login", handler.AdminLogin)

		r.Get("/affiliate", handler.GetAffiliate)

		r.Post("/affiliate", handler.AddAffiliate)

		r.Get("/affiliates", handler.GetAffiliates)
	})

}
