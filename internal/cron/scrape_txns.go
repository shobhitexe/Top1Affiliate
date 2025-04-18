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

func (c *Cron) FetchAndSaveTransactionsDeposit(ctx context.Context, cookie string, date string) error {
	emails, err := c.store.GetEmailsOfLeads(ctx)
	if err != nil {
		return err
	}

	client := &http.Client{}
	limit := 100

	for _, e := range emails {

		if e.Email == "N/A" {
			continue
		}

		offset := 0

		for {
			url := fmt.Sprintf("https://publicapi.fxlvls.com/management/lead-transactions?limit=%d&offset=%d&email=%s&transactionType=Deposit&dateFrom=%s", limit, offset, e.Email, date)

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return err
			}
			req.Header.Add("Cookie", cookie)

			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Error fetching data for %s: %v\n", e.Email, err)
				break
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				log.Printf("Error reading response for %s: %v\n", e.Email, err)
				break
			}

			var data []models.Transaction
			if err := json.Unmarshal(body, &data); err != nil {
				log.Printf("Error parsing JSON for %s: %v\n", e.Email, err)
				break
			}

			if len(data) == 0 {
				log.Printf("No more transactions for %s\n", e.Email)
				break
			}

			go func() {

				gCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				for _, d := range data {

					switch d.Status {

					case "Failed":
						message := fmt.Sprintf(
							`Deposit - Failed
						
						Crm ID:            %d
						Name:              %s
						Country:           %s
						Email:             %s
						Amount:             %.2f
						Aff id:            %s`,
							d.LeadID,
							"",
							"",
							e.Email,
							d.Amount,
							e.AffiliateID,
						)

						if err := c.utils.SendNotificationToSlack(gCtx, models.FailedDeposit, message); err != nil {
							log.Println(err)
							continue
						}

					case "Complete":
						message := fmt.Sprintf(
							`Deposit - successful
						
						Crm ID:            %d
						Name:              %s
						Country:           %s
						Email:             %s
						Amount:             %.2f
						Aff id:            %s`,
							d.LeadID,
							"",
							"",
							e.Email,
							d.Amount,
							e.AffiliateID,
						)

						if err := c.utils.SendNotificationToSlack(gCtx, models.SuccessfullDeposit, message); err != nil {
							log.Println(err)
							continue
						}

					default:
						continue
					}

				}

			}()

			if err := c.store.SaveTransactionsAndUpdateBalanceDeposit(ctx, data, e.Email, e.AffiliateID); err != nil {
				log.Println(err)
				break
			}

			log.Printf("Fetched %d transactions for %s", len(data), e.Email)

			offset += len(data)
		}
	}

	return nil
}

func (c *Cron) FetchAndSaveTransactionsWithdrawals(ctx context.Context, cookie string, date string) error {

	emails, err := c.store.GetEmailsOfLeads(ctx)
	if err != nil {
		return err
	}

	client := &http.Client{}
	limit := 100

	for _, e := range emails {

		if e.Email == "N/A" {
			continue
		}

		offset := 0

		for {
			url := fmt.Sprintf("https://publicapi.fxlvls.com/management/lead-transactions?limit=%d&offset=%d&email=%s&transactionType=Withdrawal&dateFrom=%s", limit, offset, e.Email, date)

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return err
			}
			req.Header.Add("Cookie", cookie)

			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Error fetching data for %s: %v\n", e.Email, err)
				break
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				log.Printf("Error reading response for %s: %v\n", e.Email, err)
				break
			}

			var data []models.Transaction
			if err := json.Unmarshal(body, &data); err != nil {
				log.Printf("Error parsing JSON for %s: %v\n", e.Email, err)
				break
			}

			if len(data) == 0 {
				log.Printf("No more transactions for %s\n", e.Email)
				break
			}

			if err := c.store.SaveTransactionsAndUpdateBalanceWithdraw(ctx, data, e.Email, e.AffiliateID); err != nil {
				log.Println(err)
				break
			}

			log.Printf("Fetched %d transactions for %s", len(data), e.Email)

			offset += len(data)
		}
	}

	return nil
}

// func (c *Cron) FetchAndSaveTransactionsHistory(ctx context.Context, cookie string) error {
// 	emails, err := c.store.GetAllEmails(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	client := &http.Client{}
// 	limit := 100

// 	for _, e := range emails {
// 		offset := 0

// 		for {
// 			url := fmt.Sprintf("https://publicapi.fxlvls.com/management/lead-transactions?limit=%d&email=%s&transactionType=Withdrawal&offset=%d", limit, e.Email, offset)

// 			req, err := http.NewRequest("GET", url, nil)
// 			if err != nil {
// 				return err
// 			}
// 			req.Header.Add("Cookie", cookie)

// 			resp, err := client.Do(req)
// 			if err != nil {
// 				log.Printf("Error fetching data for %s at offset %d: %v\n", e.Email, offset, err)
// 				break
// 			}

// 			body, err := io.ReadAll(resp.Body)
// 			resp.Body.Close()
// 			if err != nil {
// 				log.Printf("Error reading response for %s at offset %d: %v\n", e.Email, offset, err)
// 				break
// 			}

// 			var data []models.Transaction
// 			if err := json.Unmarshal(body, &data); err != nil {
// 				log.Printf("Error parsing JSON for %s at offset %d: %v\n", e.Email, offset, err)
// 				break
// 			}

// 			if len(data) == 0 {
// 				log.Printf("No more transactions for %s, stopping at offset %d\n", e.Email, offset)
// 				break
// 			}

// 			if err := c.store.SaveTransactions(ctx, data, e.Email, e.AffiliateID); err != nil {
// 				log.Println(err)
// 				break
// 			}

// 			log.Printf("Fetched %d transactions for %s at offset %d\n", len(data), e.Email, offset)

// 			offset += limit
// 		}
// 	}

// 	return nil
// }
