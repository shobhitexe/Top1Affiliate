package handlers

import (
	"net/http"
	"top1affiliate/internal/models"
	"top1affiliate/internal/service"
	"top1affiliate/pkg/utils"
)

type DataHandler struct {
	service service.DataService
	utils   utils.Utils
}

func NewDataHandler(service service.DataService, utils utils.Utils) *DataHandler {
	return &DataHandler{service: service, utils: utils}
}

func (h *DataHandler) Getstatistics(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("affiliateId")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.Getstatistics(r.Context(), id)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *DataHandler) GetWeeklyStats(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("affiliateId")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.GetweeklyStatsWithMonthly(r.Context(), id)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *DataHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("affiliateId")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	from := r.URL.Query().Get("from")

	if from == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading from", Data: []any{}})
		return
	}

	to := r.URL.Query().Get("to")

	if to == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading to", Data: []any{}})
		return
	}

	s, err := h.service.GetTransactions(r.Context(), id, from, to)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *DataHandler) GetDashboardStats(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("affiliateId")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.GetDashboardStats(r.Context(), id)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *DataHandler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {

	leaderboard, err := h.service.GetLeaderboard(r.Context())

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: leaderboard})
}
