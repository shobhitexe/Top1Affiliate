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

type UserHandler struct {
	service service.UserService
	utils   utils.Utils
}

func NewUserHandler(service service.UserService, utils utils.Utils) *UserHandler {
	return &UserHandler{service: service, utils: utils}
}

func (h *UserHandler) UserLogin(w http.ResponseWriter, r *http.Request) {

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

	res, err := h.service.UserLogin(r.Context(), payload)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed", Data: err.Error()})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: res})
}

func (h *UserHandler) RequestPayout(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	defer r.Body.Close()

	var payload models.RequestPayout

	if err := json.Unmarshal(body, &payload); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	if payload.ID == "" || payload.Type == "" || payload.Method == "" || payload.Amount == 0 {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if err := h.service.RequestPayout(r.Context(), payload); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed", Data: err.Error()})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: true})
}

func (h *UserHandler) GetPayouts(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if id == "" || from == "" || to == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: []any{}})
		return
	}

	s, err := h.service.GetPayouts(r.Context(), id, from, to)

	if err != nil {
		log.Println(err)
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: []any{}})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *UserHandler) GetWalletDetails(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error Reading id", Data: false})
		return
	}

	s, err := h.service.GetWalletDetails(r.Context(), id)

	if err != nil {
		log.Println(err)
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Error", Data: false})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: s})

}

func (h *UserHandler) UpdateWallet(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	defer r.Body.Close()

	var payload models.WalletDetails

	if err := json.Unmarshal(body, &payload); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: err.Error()})
		return
	}

	if payload.ID == "" {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed to read body", Data: false})
		return
	}

	if err := h.service.UpdateWalletDetails(r.Context(), payload); err != nil {
		h.utils.WriteJSON(w, http.StatusInternalServerError, models.Response{Message: "Failed", Data: err.Error()})
		return
	}

	h.utils.WriteJSON(w, http.StatusOK, models.Response{Message: "Fetched", Data: true})
}
