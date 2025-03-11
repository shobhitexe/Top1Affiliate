package cron

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func (c Cron) LoginToAPI() (string, error) {
	client := &http.Client{}

	data := map[string]string{
		"login":    os.Getenv("API_USERNAME"),
		"password": os.Getenv("API_PASSWORD"),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return "", err
	}

	req, err := http.NewRequest("POST", "https://publicapi.fxlvls.com/login", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if len(resp.Cookies()) == 0 {
		log.Println("No cookies received from the response")
		return "", nil
	}

	cookie := resp.Cookies()[0]

	authCookie := cookie.Name + "=" + cookie.Value

	return authCookie, nil
}
