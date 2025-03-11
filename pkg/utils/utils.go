package utils

import (
	"encoding/json"
	"top1affiliate/internal/models"

	"net/http"
)

type Utils interface {
	WriteJSON(w http.ResponseWriter, statusCode int, v models.Response)
}

type utils struct {
}

func NewUtils() Utils {
	return &utils{}
}

func (u *utils) WriteJSON(w http.ResponseWriter, statusCode int, v models.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}
