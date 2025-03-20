package routes

import (
	"top1affiliate/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterDataRoutes(r chi.Router, handler *handlers.DataHandler) {

	r.Route("/data", func(r chi.Router) {
		r.Get("/dashboard", handler.GetDashboardStats)

		r.Get("/statistics", handler.Getstatistics)

		r.Get("/weekly-stats", handler.GetWeeklyStats)

		r.Get("/transactions", handler.GetTransactions)

		r.Get("/leaderboard", handler.GetLeaderboard)
	})

}
