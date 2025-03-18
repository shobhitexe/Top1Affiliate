package cron

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"top1affiliate/internal/models"
)

func (c *Cron) FetchAndSaveTransactionsHistory(ctx context.Context, cookie string) error {
	emails, err := c.store.GetAllEmails(ctx)
	if err != nil {
		return err
	}

	client := &http.Client{}
	limit := 100

	for _, e := range emails {
		offset := 0

		for {
			url := fmt.Sprintf("https://publicapi.fxlvls.com/management/lead-transactions?limit=%d&email=%s&transactionType=Deposit&offset=%d", limit, e.Email, offset)

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return err
			}
			req.Header.Add("Cookie", cookie)

			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Error fetching data for %s at offset %d: %v\n", e.Email, offset, err)
				break
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				log.Printf("Error reading response for %s at offset %d: %v\n", e.Email, offset, err)
				break
			}

			var data []models.Transaction
			if err := json.Unmarshal(body, &data); err != nil {
				log.Printf("Error parsing JSON for %s at offset %d: %v\n", e.Email, offset, err)
				break
			}

			if len(data) == 0 {
				log.Printf("No more transactions for %s, stopping at offset %d\n", e.Email, offset)
				break
			}

			if err := c.store.SaveTransactions(ctx, data, e.Email, e.AffiliateID); err != nil {
				log.Println(err)
				break
			}

			log.Printf("Fetched %d transactions for %s at offset %d\n", len(data), e.Email, offset)

			offset += limit
		}
	}

	return nil
}
