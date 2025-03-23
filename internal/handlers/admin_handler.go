package handlers

import (
	"encoding/json"
	"io"
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

	aff, err := h.service.GetAffiliates(r.Context())

	if err != nil {
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
