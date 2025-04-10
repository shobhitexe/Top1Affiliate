package routes

import (
	"top1affiliate/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterDataRoutes(r chi.Router, handler *handlers.DataHandler) {

	r.Route("/data", func(r chi.Router) {
		r.Get("/dashboard", handler.GetDashboardStats)

		r.Get("/balance", handler.GetBalance)

		r.Get("/netstats", handler.GetNetStats)

		r.Get("/statistics", handler.Getstatistics)

		r.Get("/weekly-stats", handler.GetWeeklyStats)

		r.Get("/transactions", handler.GetTransactions)

		r.Get("/leaderboard", handler.GetLeaderboard)

		r.Get("/sub", handler.GetSubAffiliates)

		r.Get("/path", handler.GetAffiliatePath)
		r.Get("/tree", handler.GetAffiliateTree)
		r.Get("/list", handler.GetAffiliateList)
	})

}
