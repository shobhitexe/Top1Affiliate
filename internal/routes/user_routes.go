package routes

import (
	"top1affiliate/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, handler *handlers.UserHandler) {

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.UserLogin)
	})

	r.Route("/wallet", func(r chi.Router) {
		r.Get("/payouts", handler.GetPayouts)
		r.Post("/payout", handler.RequestPayout)

		r.Get("/details", handler.GetWalletDetails)
		r.Post("/details", handler.UpdateWallet)
	})

}
