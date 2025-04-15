package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"top1affiliate/internal/models"
	"top1affiliate/internal/service"
	"top1affiliate/pkg/utils"
)

type AdminHandler struct {
	service service.AdminService
	utils   utils.Utils
}

func NewAdminHandler(service service.AdminService, utils utils.Utils) *AdminHandler {
	return &AdminHandler{service: service, utils: utils}
}

func (h *AdminHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	defer r.Body.Close()

	var payload models.LoginRequest

	if err := json.Unmarshal(body, &payload); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	res, err := h.service.AdminLogin(r.Context(), payload)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed", Data: err.Error()})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: res})
}

func (h *AdminHandler) GetAffiliates(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	aff, err := h.service.GetAffiliates(r.Context(), id)

	if err != nil {
		log.Println(err)
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: aff})
}

func (h *AdminHandler) GetAffiliate(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.GetAffiliate(r.Context(), id)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *AdminHandler) AddAffiliate(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	defer r.Body.Close()

	var payload models.AddAffiliate

	if err := json.Unmarshal(body, &payload); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.AffiliateID == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.Country == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.Name == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.Password == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.Commission == 0 {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.ClientLink == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.SubLink == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if err := h.service.AddAffiliate(r.Context(), payload); err != nil {
		log.Println(err)
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: false})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: true})

}

func (h *AdminHandler) BlockAffiliate(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	defer r.Body.Close()

	var payload struct {
		ID string `json:"id"`
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.ID == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if err := h.service.BlockAffiliate(r.Context(), payload.ID); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: false})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: true})

}

func (h *AdminHandler) EditAffiliate(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	defer r.Body.Close()

	var payload models.EditAffiliate

	if err := json.Unmarshal(body, &payload); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.ID == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.Country == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.Name == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.ClientLink == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.SubLink == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.Commission == 0 {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if err := h.service.EditAffiliate(r.Context(), payload); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: err.Error(), Data: false})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: true})

}

func (h *AdminHandler) GetPayouts(w http.ResponseWriter, r *http.Request) {

	typevar := r.URL.Query().Get("type")

	if typevar == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.GetPayouts(r.Context(), typevar)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *AdminHandler) DeclinePayout(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	defer r.Body.Close()

	var payload struct {
		ID string `json:"id"`
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		log.Println(payload)
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.ID == "" {
		log.Println(payload)
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if err := h.service.DeclinePayout(r.Context(), payload.ID); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: err.Error(), Data: false})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: true})

}

func (h *AdminHandler) ApprovePayout(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	defer r.Body.Close()

	var payload struct {
		ID     string  `json:"id"`
		Amount float64 `json:"amount"`
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		log.Println(payload)
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if payload.ID == "" || payload.Amount == 0 {
		log.Println(payload)
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if err := h.service.ApprovePayout(r.Context(), payload.ID, payload.Amount); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: err.Error(), Data: false})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: true})

}
