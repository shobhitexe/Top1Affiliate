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
		r.Post("/block", handler.BlockAffiliate)
		r.Post("/edit", handler.EditAffiliate)

		r.Get("/affiliates", handler.GetAffiliates)

		r.Route("/wallet", func(r chi.Router) {

			r.Get("/payouts", handler.GetPayouts)

			r.Route("/payout", func(r chi.Router) {
				r.Post("/decline", handler.DeclinePayout)
				r.Post("/accept", handler.ApprovePayout)
			})

		})
	})

}
