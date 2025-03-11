package cron

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"top1affiliate/internal/models"
)

const apiURL = "https://publicapi.fxlvls.com/management/leads"

func (c *Cron) FetchAndSaveLeads(ctx context.Context, cookie string) error {
	client := &http.Client{}
	limit := 100
	minRegistrationDate := "2020-01-01 00:00" // Start from an old date

	for {
		// Build the API request URL with pagination
		url := fmt.Sprintf("%s?limit=%d&minRegistrationDate=%s", apiURL, limit, minRegistrationDate)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		req.Header.Add("Cookie", cookie)

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var responseData []models.Leads
		if err := json.Unmarshal(body, &responseData); err != nil {
			log.Println("Error decoding response:", err)
			return err
		}

		// Stop if no more data is returned
		if len(responseData) == 0 {
			log.Println("No more leads to fetch.")
			break
		}

		// Save the fetched leads
		for _, lead := range responseData {
			if err := c.store.SaveLeadsData(ctx, lead); err != nil {
				log.Println("Error saving lead:", err)
				return err
			}
		}

		// Update minRegistrationDate for next batch
		lastLead := responseData[len(responseData)-1]
		minRegistrationDate = lastLead.RegistrationDate // Assuming RegistrationDate field exists
		log.Printf("Fetched %d leads, next minRegistrationDate: %s", len(responseData), minRegistrationDate)

		// Optional: Add a small delay to avoid hitting rate limits
		time.Sleep(1 * time.Second)
	}

	return nil
}
