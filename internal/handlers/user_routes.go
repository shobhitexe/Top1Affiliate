package handlers

import (
	"encoding/json"
	"io"
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
