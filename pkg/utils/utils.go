package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"top1affiliate/internal/models"

	"net/http"
)

type Utils interface {
	WriteJSON(w http.ResponseWriter, statusCode int, v models.Response)
	SendNotificationToSlack(ctx context.Context, url, message string) error
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

func (u *utils) SendNotificationToSlack(ctx context.Context, url, message string) error {

	payload := map[string]interface{}{
		"text": message,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", models.Newregistrations, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
