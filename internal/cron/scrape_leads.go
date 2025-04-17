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

func (c *Cron) FetchAndSaveLeads(ctx context.Context, cookie, minDate, maxDate string) error {
	client := &http.Client{}
	limit := 100

	url := fmt.Sprintf("https://publicapi.fxlvls.com/management/leads?limit=%d&minRegistrationDate=%s&maxRegistrationDate=%s", limit, minDate, maxDate)

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

	if len(responseData) == 0 {
		log.Println("No new leads found.")
		return nil
	}

	for _, lead := range responseData {
		if err := c.store.SaveLeadsData(ctx, lead); err != nil {
			log.Println("Error saving lead:", err)
			return err
		}

		go func() {

			gCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			message := fmt.Sprintf(`New Registration

			Crm ID: %s
			Name: %s
			Country: %s
			Email: %s
			Registration date: %s
			Aff id: %s`,
				lead.LeadGuid,
				lead.FirstName,
				lead.Country,
				lead.Email,
				lead.RegistrationDate,
				lead.AffiliateID,
			)

			if err := c.utils.SendNotificationToSlack(gCtx, models.Newregistrations, message); err != nil {
				log.Println(err)
				return
			}
		}()

	}

	log.Printf("Fetched %d leads between %s and %s", len(responseData), minDate, maxDate)
	return nil
}

const apiURL = "https://publicapi.fxlvls.com/management/leads"

func (c *Cron) FetchAndSaveLeadsHistory(ctx context.Context, cookie string) error {
	client := &http.Client{}
	limit := 100
	minRegistrationDate := "2020-01-01 00:00"

	for {
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

		if len(responseData) == 0 {
			log.Println("No more leads to fetch.")
			break
		}

		for _, lead := range responseData {
			if err := c.store.SaveLeadsData(ctx, lead); err != nil {
				log.Println("Error saving lead:", err)
				return err
			}
		}

		lastLead := responseData[len(responseData)-1]
		minRegistrationDate = lastLead.RegistrationDate
		log.Printf("Fetched %d leads, next minRegistrationDate: %s", len(responseData), minRegistrationDate)

		time.Sleep(1 * time.Second)
	}

	return nil
}
