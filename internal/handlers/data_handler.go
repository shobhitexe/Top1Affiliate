package handlers

import (
	"log"
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

func (h *DataHandler) GetBalance(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("affiliateId")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.GetBalance(r.Context(), id)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *DataHandler) GetNetStats(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("affiliateId")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.GetNetStats(r.Context(), id)

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

func (h *DataHandler) GetSubAffiliates(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	aff, err := h.service.GetSubAffiliates(r.Context(), id)

	if err != nil {
		log.Println(err)
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: aff})
}

func (h *DataHandler) GetAffiliatePath(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.GetSubAffiliatePath(r.Context(), id)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *DataHandler) GetAffiliateTree(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.GetSubAffiliateTree(r.Context(), id)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: err.Error(), Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *DataHandler) GetAffiliateList(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.GetSubAffiliateList(r.Context(), id)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: err.Error(), Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}
