package cron

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
	"top1affiliate/internal/models"
)

func (c *Cron) FetchAndSaveLeads(ctx context.Context, cookie string) error {
	client := &http.Client{}
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 2, 2, 0, 0, 0, 0, time.UTC)

	var wg sync.WaitGroup
	concurrentRequests := make(chan struct{}, 5) // Limit concurrency to 5 requests at a time

	// Process data in 6-month chunks
	for date := startDate; date.Before(endDate); date = date.AddDate(0, 6, 0) {
		minDate := date // Capture loop variable to prevent Goroutine issues
		maxDate := minDate.AddDate(0, 6, 0)

		wg.Add(1)
		concurrentRequests <- struct{}{}

		go func(minDate, maxDate time.Time) {
			defer wg.Done()
			defer func() { <-concurrentRequests }() // Release slot

			log.Printf("Starting fetch for %s to %s", minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"))

			page := 1
			for {
				apiURL := fmt.Sprintf(
					"https://publicapi.fxlvls.com/management/leads?limit=100&page=%d&minUpdated=%s&maxUpdated=%s",
					page,
					minDate.Format("2006-01-02 15:04"),
					maxDate.Format("2006-01-02 15:04"),
				)

				req, err := http.NewRequest("GET", apiURL, nil)
				if err != nil {
					log.Println("Error creating request:", err)
					return
				}

				req.Header.Add("Cookie", cookie)
				resp, err := client.Do(req)
				if err != nil {
					log.Println("Error making request:", err)
					return
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Println("Error reading response body:", err)
					return
				}

				var responseData []models.Leads
				if err := json.Unmarshal(body, &responseData); err != nil {
					log.Println("Error unmarshalling JSON:", err)
					return
				}

				// Stop pagination if no more data
				if len(responseData) == 0 {
					log.Printf("No more leads for %s to %s at page %d", minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"), page)
					break
				}

				// Save data concurrently
				for _, data := range responseData {
					if err := c.store.SaveLeadsData(ctx, data); err != nil {
						log.Println("Error saving data:", err)
						return
					}
				}

				log.Printf("Fetched and saved leads from %s to %s, page %d", minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"), page)
				page++
			}
		}(minDate, maxDate) // Pass dates as arguments to avoid Goroutine loop issues
	}

	wg.Wait() // Wait for all Goroutines to finish
	return nil
}
