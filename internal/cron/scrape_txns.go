package cron

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Cron) FetchAndSaveTransactionsOld(ctx context.Context, cookie string) error {

	emails, err := c.store.GetAllEmails(ctx)

	if err != nil {
		return err
	}

	for _, e := range emails {

		client := &http.Client{}
		limit := 100

		url := fmt.Sprintf("https://publicapi.fxlvls.com/management/lead-transactions?limit=%d&email=%s", limit, e)

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

		log.Println(string(body))

	}

	return nil
}
